import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import { HashRouter } from 'react-router-dom'
import FastSeerTypeAheadApp from './App';
import SearchResultsApp from './SearchResultsApp'
//import registerServiceWorker from './registerServiceWorker';

ReactDOM.render(
    <HashRouter>
        <FastSeerTypeAheadApp />
    </HashRouter>
, document.getElementById('fs-type-ahead'));


ReactDOM.render(
    <HashRouter>
        <SearchResultsApp />
    </HashRouter>
, document.getElementById('fs-results'));
//registerServiceWorker();
