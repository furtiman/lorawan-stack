// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
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

import React from 'react'
import DOM from 'react-dom'
import { Provider } from 'react-redux'
import * as Sentry from '@sentry/browser'

import sentryConfig from '@ttn-lw/constants/sentry'
import store, { history } from '@account/store'

import WithLocale from '@ttn-lw/lib/components/with-locale'
import Init from '@ttn-lw/lib/components/init'

import { selectSentryDsnConfig, selectAssetsRootPath } from '@ttn-lw/lib/selectors/env'

import '@ttn-lw/lib/yup'

// Initialize sentry before creating store
if (selectSentryDsnConfig()) {
  Sentry.init(sentryConfig)
}

const assetsRoot = selectAssetsRootPath()

// Set asset path based on passed env configuration.
// eslint-disable-next-line no-undef
__webpack_public_path__ = `${assetsRoot}/`

const render = () => {
  const App = require('./account/views/app').default

  DOM.render(
    <Provider store={store}>
      <WithLocale>
        <div id="modal-container" />
        <Init>
          <App history={history} />
        </Init>
      </WithLocale>
    </Provider>,
    document.getElementById('app'),
  )
}

if (module.hot) {
  module.hot.accept('./account/views/app', () => {
    setTimeout(render)
  })
}

render()
