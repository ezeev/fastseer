import React, { Component } from 'react';
import FastSeerTypeAhead from './FastSeerTypeAhead'
import {
  BrowserRouter as Router,
  Route,
  Link
} from 'react-router-dom'
import Switch from '../node_modules/react-router-dom/Switch';

class FastSeerTypeAheadApp extends Component {

  constructor(props) {
    super(props);
    this.state = {
        openModal: false,
    }
  }

  componentDidMount() {

  }

  componentWillMount() {
  }

  render() {
    return (
      <div className="App">
        <Switch>
          <Route path="/search" render={()=><FastSeerTypeAhead appDomain={window.appDomain} shop={window.shop}/>}/>
          <Route path="/" render={()=><Link to="/search">Open Search</Link>}></Route>
        </Switch>
      </div>

    );
  }
}

export default FastSeerTypeAheadApp;
