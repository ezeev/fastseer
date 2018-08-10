import React from 'react';
import {
  Layout,
  Banner,
  SettingToggle,

} from '@shopify/polaris';
import RulesItem from './RulesItem.js'

class RulesGrid extends React.Component {

    // Constructor is the only place you can assign state
    constructor(props) {
        super(props);
        this.state = {
        }
    }

    componentDidMount() {
    
    }
  
    componentWillUnmount() {
     
    }
    
    render() {
        return (
            <Layout.AnnotatedSection
                title="Search Rules"
                description="Manage your search rules here!">
            <RulesItem></RulesItem>
            </Layout.AnnotatedSection>
            
        );
    }
}

export default RulesGrid