import React from 'react';
import {
  Layout,
  Card,
  ResourceList,
} from '@shopify/polaris';
import RulesItem from './RulesItem.js'

/*const rules =    [
    {"id":"rule01",
    "matchQueryTriggers_ss":["ipad", "apple ipad", "ipad pro"],
    "containsAnyQueryTriggers_txt":["ipad", "apple", "pro"],
    "containsFqs_ss":["cat:001",
      "cat:002"],
    "actReplaceQuery_s":"i was replaced!",
    "actAddFqs_ss":["cat:001",
      "cat:002"],
    "actAddBqs_ss":["sku:001^0.4",
      "sku:002^0.2"],
    "tags_ss":["test"],
    "actAddFacetFields_ss": ["cat","platform"],
    "order_i":1,
    "tags_ss": ["tag1","tag2"]},
    {"id":"rule02",
    "matchQueryTriggers_ss":["ipad", "apple ipad", "ipad pro"],
    "containsAnyQueryTriggers_txt":["ipad", "apple", "pro"],
    "containsFqs_ss":["cat:001",
      "cat:002"],
    "actReplaceQuery_s":"i was replaced!",
    "actAddFqs_ss":["cat:001",
      "cat:002"],
    "actAddBqs_ss":["sku:001^0.4",
      "sku:002^0.2"],
    "tags_ss":["test"],
    "actAddFacetFields_ss": ["cat","platform"],
    "order_i":1,
    "tags_ss": ["tag1","tag2"]}
]*/

class RulesGrid extends React.Component {

    // Constructor is the only place you can assign state
    constructor(props) {
        super(props);
        this.state = {
            rules : [],
            loading: false,
        }
    }

    componentDidMount() {
        this.getRules();
    }

  
    componentWillUnmount() {
    }

    handleNewRule() {
        console.log("I was clicked")
        //TODO: Finish this!
        // thinking that it should open a model with a new scaffolded RulesItem
    }
    
    getRules() {
        fetch(this.props.appDomain + "/api/v1/shop/rules?" + window.authQueryString(this.props) + "&q=*")
        .then(res => res.json())
        .then(
            (result) => {
                console.log(result);
                this.setState({
                    rules: result,
                });
            },
        )
        .catch(error => {
            this.setState({
                errorMsg: "Unable to get index stats. Please contact support."
            });
        });
    }


    render() {

        let ruleItems = [];
        for (var i=0;i<this.state.rules.length;i++) {
            ruleItems.push(<RulesItem key={i} rule={this.state.rules[i]}
                appDomain={this.props.appDomain} shop={this.props.shop} locale={this.props.locale} timestamp={this.props.timestamp} hmac={this.props.hmac}
                ></RulesItem>)
        }


        return (
            <Layout.AnnotatedSection
                title="Search Rules"
                description="Manage your search rules here!"
                >
                <Card sectioned primaryFooterAction={{
                        content: 'New Rule',
                        onAction: this.handleNewRule,
                    }}>not sure</Card>
                    {ruleItems}
            </Layout.AnnotatedSection>
            
        );
    }
}

export default RulesGrid