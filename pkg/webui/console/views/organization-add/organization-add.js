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

import React, { useState, useCallback } from 'react'
import { Container, Col, Row } from 'react-grid-system'
import { defineMessages } from 'react-intl'

import PageTitle from '@ttn-lw/components/page-title'
import SubmitBar from '@ttn-lw/components/submit-bar'
import SubmitButton from '@ttn-lw/components/submit-button'
import Form from '@ttn-lw/components/form'

import OrganizationForm from '@console/components/organization-form'

import withFeatureRequirement from '@console/lib/components/with-feature-requirement'

import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'
import { getOrganizationId } from '@ttn-lw/lib/selectors/id'

import { mayCreateOrganizations } from '@console/lib/feature-checks'

const initialValues = {
  ids: {
    organization_id: '',
  },
  name: '',
  description: '',
}

const m = defineMessages({
  createOrganization: 'Create organization',
})

const OrganizationAdd = props => {
  const { createOrganization, createOrganizationSuccess } = props

  const [error, setError] = useState('')

  const handleSubmitSuccess = useCallback(
    organization => {
      const orgId = getOrganizationId(organization)

      createOrganizationSuccess(orgId)
    },
    [createOrganizationSuccess],
  )

  const handleSubmitFailure = useCallback(error => {
    setError(error)
  }, [])

  return (
    <Container>
      <PageTitle tall title={sharedMessages.addOrganization} />
      <Row>
        <Col md={10} lg={9}>
          <OrganizationForm
            error={error}
            onSubmit={createOrganization}
            onSubmitSuccess={handleSubmitSuccess}
            onSubmitFailure={handleSubmitFailure}
            initialValues={initialValues}
          >
            <SubmitBar>
              <Form.Submit message={m.createOrganization} component={SubmitButton} />
            </SubmitBar>
          </OrganizationForm>
        </Col>
      </Row>
    </Container>
  )
}

OrganizationAdd.propTypes = {
  createOrganization: PropTypes.func.isRequired,
  createOrganizationSuccess: PropTypes.func.isRequired,
}

export default withFeatureRequirement(mayCreateOrganizations, { redirect: '/organizations' })(
  OrganizationAdd,
)
