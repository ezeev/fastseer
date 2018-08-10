import React, { Component } from 'react';
import FastSeerTypeAhead from './FastSeerTypeAhead'
//import './App.css';

class FastSeerTypeAheadApp extends Component {

  constructor(props) {
    super(props);
    /*var urlParams = new URLSearchParams(window.location.search);
    this.state = {
        shop: urlParams.get("shop"),
        appDomain: urlParams.get("appDomain"),
    }*/
  }


  render() {
    return (
      <div className="App">
        <FastSeerTypeAhead appDomain={window.appDomain} shop={window.shop}></FastSeerTypeAhead>
      </div>

    );
  }
}

export default FastSeerTypeAheadApp;
