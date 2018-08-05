/*import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import registerServiceWorker from './registerServiceWorker';

ReactDOM.render(<App />, document.getElementById('root'));
registerServiceWorker();*/

import React from 'react';
import ReactDOM from 'react-dom';
import '@shopify/polaris/styles.css';
import {AppProvider} from '@shopify/polaris';

import App from './App';

ReactDOM.render(
  <AppProvider>
    <App />
  </AppProvider>,
  document.getElementById('root'),
);