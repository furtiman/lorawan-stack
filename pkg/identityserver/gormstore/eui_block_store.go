// Copyright © 2021 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package store

import (
	"context"
	"fmt"
	"runtime/trace"

	"github.com/jinzhu/gorm"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/identityserver/store"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
)

// GetEUIStore returns an EUIStore on the given db (or transaction).
func GetEUIStore(db *gorm.DB) store.EUIStore {
	return &euiStore{baseStore: newStore(db)}
}

type euiStore struct {
	*baseStore
}

func getMaxCounter(addressBlock types.EUI64Prefix) int64 {
	return int64(^(uint64(0)) >> (addressBlock.Length))
}

var (
	errMaxGlobalEUILimitReached = errors.DefineInvalidArgument("global_eui_limit_reached", "global eui limit from address block reached")
	errAppDevEUILimitReached    = errors.DefineInvalidArgument("application_dev_eui_limit_reached", "application issued DevEUI limit ({dev_eui_limit}) reached")
)

func (s *euiStore) incrementApplicationDevEUICounter(ctx context.Context, ids *ttnpb.ApplicationIdentifiers, applicationLimit int) error {
	var appModel Application
	// Check if application exists.
	query := s.query(ctx, Application{}, withApplicationID(ids.GetApplicationId()))
	if err := query.First(&appModel).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errNotFoundForID(ids)
		}
		return err
	}
	// If app DevEUI limit not configured, skip this step.
	if applicationLimit == 0 {
		return nil
	}
	// Atomically check and increment DevEUI counter if its below threshold.
	result := s.query(ctx, Application{}).
		Where(Application{ApplicationID: appModel.ApplicationID}).
		Where(`"applications"."dev_eui_counter" < ?`, applicationLimit).
		Update("dev_eui_counter", gorm.Expr("dev_eui_counter + ?", 1))
	if err := result.Error; err != nil {
		return err
	}
	// If application DevEUI counter not updated, the limit was reached.
	if result.RowsAffected == 0 {
		return errAppDevEUILimitReached.WithAttributes("dev_eui_limit", fmt.Sprint(applicationLimit))
	}
	return nil
}

func (s *euiStore) issueDevEUIAddressFromBlock(ctx context.Context) (*types.EUI64, error) {
	var devEUIResult types.EUI64
	var euiBlock EUIBlock
	for {
		// Get current value from DevEUI db.
		query := s.query(ctx, EUIBlock{}).
			Where(`"eui_blocks"."type"='dev_eui' AND "eui_blocks"."current_counter" <= "eui_blocks"."end_counter"`)
		if err := query.First(&euiBlock).Error; err != nil {
			return nil, errMaxGlobalEUILimitReached.New()
		}
		euiBlock.CurrentCounter++
		query.Save(euiBlock)
		devEUI := euiBlock.CurrentCounter - 1
		devEUIResult.UnmarshalNumber(euiBlock.StartEUI.toPB().MarshalNumber() | uint64(devEUI))
		deviceQuery := s.query(ctx, EndDevice{}).Where(EndDevice{
			DevEUI: eui(&devEUIResult),
		})
		// Return the address assigned if no existing device uses the DevEUI
		if err := deviceQuery.First(&EndDevice{}).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				return &devEUIResult, nil
			}
			return nil, err
		}
	}
}

func (s *euiStore) createEUIBlock(ctx context.Context, euiType string, block types.EUI64Prefix, initCounterValue int64) (err error) {
	euiBlock := &EUIBlock{
		Type:           euiType,
		StartEUI:       *eui(&block.EUI64),
		MaxCounter:     getMaxCounter(block),
		CurrentCounter: initCounterValue,
	}
	return s.query(ctx, EUIBlock{}).Create(&euiBlock).Error
}

// InitializeDevEUIBlock initializes the block in IS db based on the configured block prefix
func (s *euiStore) InitializeDevEUIBlock(ctx context.Context, configPrefix types.EUI64Prefix, initCounter int64) error {
	defer trace.StartRegion(ctx, "initialize eui block").End()
	var block EUIBlock
	if err := s.query(ctx, EUIBlock{}).
		Where(EUIBlock{Type: "dev_eui", StartEUI: *eui(&configPrefix.EUI64)}).
		First(&block).Error; err != nil {
		// If block is not found, create it
		if gorm.IsRecordNotFoundError(err) {
			return s.createEUIBlock(ctx, "dev_eui", configPrefix, initCounter)
		}
		return err
	}
	return nil
}

func (s *euiStore) IssueDevEUIForApplication(ctx context.Context, ids *ttnpb.ApplicationIdentifiers, applicationLimit int) (*types.EUI64, error) {
	defer trace.StartRegion(ctx, "assign DevEUI address to application").End()
	// Check if max DevEUI per application reached.
	err := s.incrementApplicationDevEUICounter(ctx, ids, applicationLimit)
	if err != nil {
		return nil, err
	}
	// Issue an unused DevEUI address.
	devEUIAddress, err := s.issueDevEUIAddressFromBlock(ctx)
	if err != nil {
		return nil, err
	}
	return devEUIAddress, nil
}
