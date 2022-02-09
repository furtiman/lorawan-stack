// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
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

package migrations

import (
	"context"

	"github.com/jinzhu/gorm"
	store "go.thethings.network/lorawan-stack/v3/pkg/identityserver/gormstore"
)

// EUIBlocks removes the unique index constraint from `type` field.
type EUIBlocks struct{}

// Name implements Migration.
func (EUIBlocks) Name() string {
	return "eui_blocks"
}

// Apply implements Migration.
func (m EUIBlocks) Apply(ctx context.Context, db *gorm.DB) error {
	db.Model(store.EUIBlock{}).RemoveIndex("eui_block_index")
	return nil
}

// Rollback implements migration. For the EUIBlocks migration this is a no-op.
func (m EUIBlocks) Rollback(ctx context.Context, db *gorm.DB) error {
	return nil
}

func init() {
	All = append(All, EUIBlocks{})
}
