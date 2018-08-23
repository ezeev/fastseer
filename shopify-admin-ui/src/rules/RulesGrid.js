import React from 'react';
import {
  Layout,
  Card,
  Modal,
  SkeletonBodyText,
  Banner,
} from '@shopify/polaris';
import RulesItem from './RulesItem.js'


const ruleTemplate = {
        "id":"",
        "matchQueryTriggers_ss":[],
        "containsAnyQueryTriggers_txt":[],
        "containsFqs_ss":[],
        "actReplaceQuery_s":"",
        "name_s":"",
        "actAddFqs_ss":[],
        "actAddBqs_ss":[],
        "tags_ss":[],
        "order_i":1,
    }

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
            showNewRuleModal: false,
            newRule: {},
            successMsg: null,
            errorMsg: null,
        }
    }

    componentDidMount() {
        this.getRules();
    }

  
    componentWillUnmount() {
    }


    newRuleTemplate = () => {
        let newRule = Object.assign({}, ruleTemplate);
        return newRule;
    }

    handleNewRule = () => {
        // 1. open a modal
        // 2. create an empty rule
        // 3. render a new <RulesItem /> with the new rule
        this.setState({
            showNewRuleModal: true,
            newRule: this.newRuleTemplate(),
        });

    }

    handleCloseModal = () => {
        this.setState({
            showNewRuleModal: false,
        });
        this.getRules();
    }
    
    getRules() {
        this.setState({loading: true, rules: []});
        fetch(this.props.appDomain + "/api/v1/shop/rules?" + window.authQueryString(this.props) + "&q=*")
        .then(res => res.json())
        .then(
            (result) => {
                this.setState({
                    rules: result,
                    loading: false,
                });
            },
        )
        .catch(error => {
            this.setState({
                errorMsg: "Unable to get index stats. Please contact support.",
                loading: false,
            });
        });
    }

    handleDeleteItem = (id) => {
        console.log("deleting "+ id);
       fetch(this.props.appDomain + "/api/v1/shop/rules?id=" + id + "&" + window.authQueryString(this.props), {
            method: 'DELETE', // or 'PUT'
            headers:{
              'Content-Type': 'application/json'
            }
        })
            .then(res => res.json())
            .then((result) => {
                if (result.message) {
                    this.setState({successMsg: "Deleted rule id " + id});
                    // refresh
                    this.getRules();
                } else if (result.error) {
                    this.setState({errorMsg: result.error});
                }
                this.setState({changed: false})
            })
            .catch(error => this.setState({errorMsg: "Failed to delete: " + error, loading: false}))     
    }


    render() {

        let loading;
        if (this.state.loading) {
            loading = <SkeletonBodyText lines={100}></SkeletonBodyText>;
            
        } 

        let ruleItems = [];
        for (var i=0;i<this.state.rules.length;i++) {
            ruleItems.push(<RulesItem key={i} rule={this.state.rules[i]} onDelete={this.handleDeleteItem}
                appDomain={this.props.appDomain} shop={this.props.shop} locale={this.props.locale} timestamp={this.props.timestamp} hmac={this.props.hmac}
                ></RulesItem>)
        }

        let successBanner;
        if (this.state.successMsg) {
            successBanner = <Banner status="success" onDismiss={() => {this.setState({successMsg: null})}}>
                <p>{this.state.successMsg}</p>
            </Banner>
        }

        return (
            <Layout.AnnotatedSection
                title="Search Rules"
                description="Manage your search rules here!"
                >
                {loading}
                <Card sectioned title="Rules" 
                    actions={[{
                        content: 'New Rule',
                        onAction: this.handleNewRule,
                    }]}>                                               
                    {successBanner}
                    {ruleItems}
                    </Card>                                                        
                    <Modal
                        open={this.state.showNewRuleModal}
                        onClose={this.handleCloseModal}
                        title="Create a New Search Rule"
                        height={200}
                        /*primaryAction={{
                        content: 'Import customers',
                        onAction: this.handleChange,
                        }}
                        secondaryActions={[
                        {
                            content: 'Cancel',
                            onAction: () => {this.handleCloseModal},
                        },
                        ]}*/
                    >
                        <Modal.Section>
                            <RulesItem rule={this.state.newRule}
                                appDomain={this.props.appDomain} shop={this.props.shop} locale={this.props.locale} timestamp={this.props.timestamp} hmac={this.props.hmac}
                                ></RulesItem>
                        </Modal.Section>
                    </Modal>


            </Layout.AnnotatedSection>
        );
    }
}

export default RulesGrid