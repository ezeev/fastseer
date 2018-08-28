import React, { Component } from 'react';
import {
  BrowserRouter as Router,
  Route,
  Link,
  browserHistory,
  withRouter,
} from 'react-router-dom'
import Switch from '../node_modules/react-router-dom/Switch';
import SearchResults from './SearchResults';

class SearchResultsApp extends Component {

  constructor(props) {
    super(props);
  }

  componentDidMount() {

  }

  componentWillMount() {
  }


  render() {
    return (
      <div className="SearchResultsApp">
        <Switch >
          <Route path="/results" render={()=><SearchResults appDomain={window.appDomain} shop={window.shop}/>}/>
          <Route path="/" render={()=><span></span>}></Route>
        </Switch>
      </div>
    );
  }
}

export default withRouter(SearchResultsApp);
