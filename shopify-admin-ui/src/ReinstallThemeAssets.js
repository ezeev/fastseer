import React from 'react';
import {
  Layout,
  Banner,
  SettingToggle,
} from '@shopify/polaris';

class ReinstallThemeAssets extends React.Component {

    // Constructor is the only place you can assign state
    constructor(props) {
        super(props);
        this.state = {
            errorMsg: NaN,
            message: NaN
        }
    }

    componentDidMount() {
    
    }
  
    componentWillUnmount() {
     
    }

    handleReinstall = () => {
        fetch(this.props.appDomain + "/api/v1/shop/theme/install?" + window.authQueryString(this.props), {
            method: 'POST', // or 'PUT'
        })
            .then(res => res.json())
            .then((result) => {
                if (result.message) {
                    this.setState({message: result.message});
                } else if (result.error) {
                    this.setState({errorMsg: result.error});
                }
            })
            .catch(error => this.setState({errorMsg: "Failed to update: " + error}))
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
            title="Theme"
                >
                {message}
                {error}         
                <SettingToggle
                    action={{
                        content: 'Reinstall',
                        onAction: this.handleReinstall,
                    }}
                >
              Reinstall theme assets if you've installed a new theme or wish to reset FastSeer's theme assets to their original states.
            </SettingToggle>   
            </Layout.AnnotatedSection>
            
        );
    }
}

export default ReinstallThemeAssets