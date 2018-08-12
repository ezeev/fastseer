import React from 'react';
import {
  TextContainer,
  Layout,
  Card,
  Stack,
  Badge,
  Caption,
  Heading,
  Banner,
  Subheading,
  Icon,
  Link,
} from '@shopify/polaris';
import EditableBadge from './EditableBadge.js';

class RulesItem extends React.Component {

    // Constructor is the only place you can assign state
    constructor(props) {
        super(props);
        this.state = {
            rule: [],
            changed: false,
        }
    }

    componentDidMount() {
        this.setState({
            rule: this.props.rule,
            changed: false,
        });
    }
  
    componentWillUnmount() {
     
    }

    handleAddItem = (field, label) => {
        var newRule = this.state.rule;
        if(newRule[field] == null) {
            newRule[field] = []
        }
        newRule[field].push("new " + label + ' ' + (newRule[field].length+1))
        console.log(newRule);
        this.setState({
            rule: newRule,
            changed: true,
        });
    }

    handleValueChange = (field, index, item) => {
        console.log(field + "," + index + "," + item)
        var newRule = this.state.rule;
        newRule[field][index] = item;
        this.setState({
            rule: newRule,
            changed: true,
        });
        console.log(this.state.rule);
    }

    handleReplaceQChange = (i1, i2, item) => {
        var newRule = this.state.rule;
        newRule.actReplaceQuery_s = item
        this.setState({
            rule: newRule,
            changed: true,
        })
        console.log(this.state.rule);
    }

    handleNameChange = (i1, i2, item) => {
        var newRule = this.state.rule;
        newRule.name_s = item
        this.setState({
            rule: newRule,
            changed: true,
        })
        console.log(this.state.rule);
    }


    handleSaveRule = () => {
        fetch(this.props.appDomain + "/api/v1/shop/rules?" + window.authQueryString(this.props), {
            method: 'PUT', // or 'PUT'
            body: JSON.stringify([this.state.rule]), // data can be `string` or {object}!
            headers:{
              'Content-Type': 'application/json'
            }
        })
            .then(res => res.json())
            .then((result) => {
                if (result.message) {
                    this.setState({message: result.message});
                } else if (result.error) {
                    this.setState({errorMsg: result.error});
                }
                this.setState({changed: false})
            })
            .catch(error => this.setState({errorMsg: "Failed to update: " + error, loading: false}))
    }

    render() {
        
        let matchQueryTriggers;
        if (this.props.rule.matchQueryTriggers_ss) {
            matchQueryTriggers = this.props.rule.matchQueryTriggers_ss.map((item, index) =>
                <EditableBadge key={item+index}  value={item} field="matchQueryTriggers_ss" index={index} valueChange={this.handleValueChange}></EditableBadge>
            );
        }

        let containsAnyQueryTriggers;
        if (this.props.rule.containsAnyQueryTriggers_txt) {
            containsAnyQueryTriggers = this.props.rule.containsAnyQueryTriggers_txt.map((item, index) =>
                <EditableBadge key={item+index}  value={item} field="containsAnyQueryTriggers_txt" index={index} valueChange={this.handleValueChange}></EditableBadge>
            );
        }

        let containsFilterQueryTriggers;
        if (this.props.rule.containsFqs_ss) {
            containsFilterQueryTriggers = this.props.rule.containsFqs_ss.map((item, index) => 
                <EditableBadge key={item+index} value={item} field="containsFqs_ss" index={index} valueChange={this.handleValueChange}></EditableBadge>
            );  
        }

        let actAddFqs;
        if (this.props.rule.actAddFqs_ss) {
            actAddFqs = this.props.rule.actAddFqs_ss.map((item, index) =>
                <EditableBadge key={item+index}  value={item} field="actAddFqs_ss" index={index} valueChange={this.handleValueChange}></EditableBadge>
            );
        }

        let actAddBqs;
        if (this.props.rule.actAddBqs_ss) {
            actAddBqs = this.props.rule.actAddBqs_ss.map((item, index) =>
                <EditableBadge key={item+index}  value={item} field="actAddBqs_ss" index={index} valueChange={this.handleValueChange}></EditableBadge>
            );
        }

        let actAddFacetFields;
        if (this.props.rule.actAddFacetFields_ss) {
            actAddFacetFields = this.props.rule.actAddFacetFields_ss.map((item, index) =>
                <EditableBadge key={item+index}  value={item} field="actAddFacetFields_ss" index={index} valueChange={this.handleValueChange}></EditableBadge>
            );
         }

        let tags;
        if (this.props.rule.tags_ss) {
            tags = this.props.rule.tags_ss.map((item, index) =>
                <EditableBadge key={item+index}  value={item} field="tags_ss" index={index} valueChange={this.handleValueChange}></EditableBadge>
            ); 
        }

        let changeBanner;
        console.log(this.state.changed)
        if (this.state.changed === true) {
            changeBanner = <Banner>
                You have unsaved changes!
            </Banner>
        }
        return (
                <Card sectioned title={this.props.rule.name_s} actions={[{content: 'Save', onClick: this.handleSaveRule}, {content: 'Delete'}]}>
                    {changeBanner}
                    <div style={{position:"relative", float:"left", width:"50%", padding:"15px"}}>
                        <Subheading>Name & Tags:</Subheading>   
                        <EditableBadge value={this.props.rule.name_s} valueChange={this.handleNameChange}>                        
                            </EditableBadge> | tags:                
                            {tags} 
                            <Link onClick={this.handleAddItem.bind(this, "tags_ss", "tag")}>add</Link>
                        <Subheading>When (Triggers):</Subheading>
                        <Caption>Query Matches</Caption> {matchQueryTriggers} <Link onClick={this.handleAddItem.bind(this, "matchQueryTriggers_ss", "query")}>add</Link> <br />
                        <Caption>Query Contains</Caption> {containsAnyQueryTriggers} <Link onClick={this.handleAddItem.bind(this, "containsAnyQueryTriggers_txt", "contains query")}>add</Link> <br/>
                        <Caption>Contains Filter </Caption> {containsFilterQueryTriggers} <Link onClick={this.handleAddItem.bind(this, "containsFqs_ss", "filter")}>add</Link> 
                    </div>
                    <div style={{position:"relative", float:"left", width:"50%", padding:"15px"}}>
                        <Subheading>Then (Actions):</Subheading>
                        <Caption>Replace Query</Caption> 
                            <EditableBadge value={this.props.rule.actReplaceQuery_s} valueChange={this.handleReplaceQChange}>                        
                            </EditableBadge> <br/>                        
                        <Caption>Add Filter</Caption> {actAddFqs} <Link onClick={this.handleAddItem.bind(this, "actAddFqs_ss", "filter")}>add</Link>  <br />
                        <Caption>Add Boost</Caption> {actAddBqs} <Link onClick={this.handleAddItem.bind(this, "actAddBqs_ss", "boost")}>add</Link>  <br />
                        <Caption>Add Facet</Caption> {actAddFacetFields} <Link onClick={this.handleAddItem.bind(this, "actAddFacetFields_ss", "facet")}>add</Link> 
                    </div>
                </Card>
        );
    }
}
export default RulesItem
