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

import { defineSmokeTest } from '../utils'

const checkCollapsingFields = defineSmokeTest(
  'succeeds showing contents of collapsing fields in device general settings',
  () => {
    const userId = 'collapsing-fields-test-user'
    const user = {
      ids: { user_id: userId },
      primary_email_address: 'test-user@example.com',
      password: 'ABCDefg123!',
      password_confirm: 'ABCDefg123!',
      email: 'collapsing-fields-test-user@example.com',
    }

    const applicationId = 'collapsing-fields-app-test'
    const application = {
      ids: { application_id: applicationId },
    }
    cy.createUser(user)
    cy.createApplication(application, user.ids.user_id)
    cy.createEndDeviceAllComponents(applicationId)
    cy.loginConsole({ user_id: user.ids.user_id, password: user.password })
    cy.visit(
      `${Cypress.config(
        'consoleRootPath',
      )}/applications/${applicationId}/devices/device-all-components/general-settings`,
    )
    cy.findByText('Network layer', { selector: 'h3' })
      .closest('[data-test-id="collapsible-section"]')
      .within(() => {
        cy.findByRole('button', { name: 'Expand' }).click()
        cy.findByText('Frequency plan')
        cy.findByText('Advanced MAC settings').click()
        cy.findByText('Frame counter width')
      })

    cy.findByText('Application layer', { selector: 'h3' })
      .closest('[data-test-id="collapsible-section"]')
      .within(() => {
        cy.findByRole('button', { name: 'Expand' }).click()
        cy.findByText('Payload crypto override')
      })

    cy.findByText('Join Settings', { selector: 'h3' })
      .closest('[data-test-id="collapsible-section"]')
      .within(() => {
        cy.findByRole('button', { name: 'Expand' }).click()
        cy.findByText('Home NetID')
      })
  },
)

export default [checkCollapsingFields]
