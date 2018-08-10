

import React, {Component} from 'react';
import {
  Layout,
  Page,
  Tabs,
} from '@shopify/polaris';
import IndexingOps from './IndexingOps'
import SearchConfigGrid from './SearchConfigGrid';
import ReinstallThemeAssets from './ReinstallThemeAssets';
import RulesGrid from './rules/RulesGrid';


class App extends Component {
  constructor(props) {
    super(props);

    /*
    hmac := params.Get("hmac")
	  locale := params.Get("locale")
	  shop := params.Get("shop")
	  timestamp := params.Get("timestamp")
    */
    var urlParams = new URLSearchParams(window.location.search);
    //var appDomain = "https://api.fastseer.com"; //default to public api
    var appDomain = "";
    if (urlParams.has("appDomain")) {
        appDomain = urlParams.get("appDomain");
    }

    this.state = {
        shop: urlParams.get("shop"),
        locale: urlParams.get("locale"),
        timestamp: urlParams.get("timestamp"),
        hmac: urlParams.get("hmac"),
        appDomain: appDomain,
        selected: 0
    }
  }


  handleTabChange = (selectedTabIndex) => {
    this.setState({selected: selectedTabIndex});
  };

  render() {
    //const primaryAction = {content: 'New product'};
    //const secondaryActions = [{content: 'Import', icon: 'import'}];

    /*const choiceListItems = [
      {label: 'I accept the Terms of Service', value: 'false'},
      {label: 'I consent to receiving emails', value: 'false2'},
    ];*/
    const {selected} = this.state.selected;
    const tabs = [
      {
        id: 'home',
        content: 'Home',
        accessibilityLabel: 'Home',
        panelID: 'home',
      },
      {
        id: 'search-config',
        content: 'Search Configuration',
        panelID: 'search-config',
      },
      {
        id: 'search-rules',
        content: 'Search Rules',
        panelID: 'search-rules',
      },
    ];

    let content
    if (this.state.selected === 0) {
      content = <Layout>
          <IndexingOps appDomain={this.state.appDomain} shop={this.state.shop} locale={this.state.locale} timestamp={this.state.timestamp} hmac={this.state.hmac}></IndexingOps>
          <ReinstallThemeAssets appDomain={this.state.appDomain} shop={this.state.shop} locale={this.state.locale} timestamp={this.state.timestamp} hmac={this.state.hmac}></ReinstallThemeAssets>
        </Layout>
    } if (this.state.selected === 1) {
      content = <Layout>
          <SearchConfigGrid appDomain={this.state.appDomain} shop={this.state.shop} locale={this.state.locale} timestamp={this.state.timestamp} hmac={this.state.hmac}></SearchConfigGrid>
        </Layout>
    } if (this.state.selected === 2) {
      content = <Layout>
          <RulesGrid appDomain={this.state.appDomain} shop={this.state.shop}></RulesGrid>
        </Layout>
    }


    return (
      <Page
        //title="FastSeer"
        //primaryAction={primaryAction}
        //secondaryActions={secondaryActions}
      >
        <Tabs tabs={tabs} selected={selected} onSelect={this.handleTabChange} />
            {content}
      </Page>
    );
  }
}

export default App;