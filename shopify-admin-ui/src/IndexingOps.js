import React from 'react';
import {
  Layout,
  Banner,
  Card,
  DisplayText,
  Stack,
  Badge,
} from '@shopify/polaris';

class IndexingOps extends React.Component {

    // Constructor is the only place you can assign state
    constructor(props) {
        super(props);
        this.state = {
            date: new Date(),
            callCount: 0,
            numProducts: 0,
            numVariants: 0,
            errorMsg: NaN,
            message: NaN
        }
    }

    componentDidMount() {
        this.timerID = setInterval(
            () => this.tick(),
            5000
            );
            this.fetchCounts();

        console.log(window.authQueryString(this.props));
    }
  
    componentWillUnmount() {
        clearInterval(this.timerID);
    }

    fetchCounts = () => {
        if (this.state.callCount < 30) {
            fetch(this.props.appDomain + "/api/v1/products/count?" + window.authQueryString(this.props))
                .then(res => res.json())
                .then(
                    (result) => {
                        this.setState({
                            numProducts: result.products,
                            numVariants: result.variants,
                            errorMsg: NaN
                        });
                    },
                )
                .catch(error => {
                    this.setState({
                        errorMsg: "Unable to get index stats. Please contact support."
                    });
                });
        }

    }

    tick() {
        var currentCount = this.state.callCount;
        this.setState({
            date: new Date(),
            callCount: currentCount + 1
        }); 
        this.fetchCounts();
    }

    handleClearIndexClick = () => {
        fetch(this.props.appDomain+ "/api/v1/products/index?" + window.authQueryString(this.props) 
        , {method: "DELETE"})
        .then(res => res.json())
        .then(
            (result) => {
                if (result.error) {
                    this.setState({
                        errorMsg: result.error
                    });       
                } else {
                    this.setState({
                        message: result.message
                    });
                }
            },
        )
        .catch(error => {
            this.setState({
                errorMsg: "Unable to clear index. Please contact support."
            });
        });
    }

    handleRebuildClick = () => {
        fetch(this.props.appDomain + "/api/v1/products/index?" + window.authQueryString(this.props), {method: "POST"})
            .then(res => res.json())
            .then(
                (result) => {
                    if (result.error) {
                        this.setState({
                            errorMsg: result.error
                        });       
                    } else {
                        this.setState({
                            message: result.message
                        });
                    }
                },
            )
            .catch(error => {
                this.setState({
                    errorMsg: "Unable to clear index. Please contact support."
                });
            });
    }

    render() {
        let message;
        if (this.state.message) {        
            message = <Banner
                title={this.state.message}
                status="info"
                onDismiss={() => { this.setState({message: NaN})}}
                >
            </Banner>
        }
        let error;
        if (this.state.errorMsg) {
            error = <Banner
                title={this.state.errorMsg}
                status="critical"
                onDismiss={() => { this.setState({errorMsg: NaN})}}
                >
            </Banner>
        }

        return (
            <Layout.AnnotatedSection
            title="Indexing"
            description="If you have made bulk changes to your catalog, it is recommended to Rebuild your index.">
                {message}
                {error}
                <Card sectioned primaryFooterAction={{
                    content: "Rebuild Index",
                    onAction: this.handleRebuildClick
                    }}
                    secondaryFooterAction={{
                        content: "Clear Index",
                        onAction: this.handleClearIndexClick
                    }}>      

                    <Stack>
                        <Badge><DisplayText size="extraLarge">{this.state.numProducts}</DisplayText> Products</Badge>
                        <Badge><DisplayText size="extraLarge">{this.state.numVariants}</DisplayText> Variants</Badge>
                    </Stack>
                </Card>
            </Layout.AnnotatedSection>
        );
    }
}

export default IndexingOps