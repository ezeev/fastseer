import React from 'react';
import {
  Layout,
  Banner,
  DataTable,
  Card,
  Button,
  FormLayout,
  Stack,
  TextField,
  Modal,
  Select,
  SkeletonBodyText,
  Checkbox,
} from '@shopify/polaris';

class SearchConfigGrid extends React.Component {

    // Constructor is the only place you can assign state
    constructor(props) {
        super(props);   
        this.state = {
            shopConfig: NaN,
            rows: [],
            message: NaN,
            errorMsg: NaN,
            editSearchConfIndex: NaN,
            editSearchConf: NaN,
            loading: false
        }
    }

    componentDidMount() {
        this.fetchConfig();
    }
  
    componentWillUnmount() {
        
    }

    handleCloneButton = (index) => {
        fetch(this.props.appDomain + "/api/v1/shop/search/config/clone?index=" + index + "&" + window.authQueryString(this.props), {
            method: "POST"
        })
        .then(res => res.json())
        .then(
            (result) => {
                console.log(result)
                if (result.error) {
                    this.setState({errorMsg: result.message});
                } else {
                    this.setState({message: result.message});
                    this.fetchConfig();
                }
            })
    }

    handleAllocTextChange = (index, newVal) => {
        //console.log(newVal)
        var rows = this.state.rows;
        //console.log(rows[index][1].props.value)
        //replace text field val
        rows[index][1] = <TextField type="number" key={index} value={newVal} suffix="%"
        onChange={this.handleAllocTextChange.bind(this, index)}
        ></TextField>;
        this.setState({
            rows: rows
        });

    }

    fetchConfig = () => {
        this.setState({loading: true})
        fetch(this.props.appDomain + "/api/v1/shop/config?" + window.authQueryString(this.props))
            .then(res => res.json())
            .then(
                (result) => {
                    var rows = []
                    result.shopifySearchConfigs.forEach((conf, index) => {
                        const allocInput = <TextField type="number" key={index} value={result.searchConfigAlloc[index]} suffix="%"
                            onChange={this.handleAllocTextChange.bind(this, index)}
                            ></TextField>;

                        const editButton = <Button size="slim" key={index} onClick={this.handleEditButton.bind(this, index)}>Edit</Button>; 
                        const cloneButton = <Button size="slim" key={index} onClick={this.handleCloneButton.bind(this, index)}>Clone</Button>;
                        const row = [conf.name, allocInput, editButton, cloneButton]
                        rows.push(row);
                    });
                    this.setState({
                        shopConfig: result,
                        rows: rows,
                        loading: false
                    });                     
                },
            )
            .catch(error => {
                this.setState({
                    errorMsg: error,
                    loading: false
                });
            });
    }

    postConfig = () => {
        this.setState({loading: true})
        fetch(this.props.appDomain + "/api/v1/shop/config?" + window.authQueryString(), {
            method: 'POST', // or 'PUT'
            body: JSON.stringify(this.state.shopConfig), // data can be `string` or {object}!
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
                this.setState({loading: false})
            })
            .catch(error => this.setState({errorMsg: "Failed to update: " + error, loading: false}))
    }

    handleUpdateAlloc = () => {
        var searchConfigAlloc = []
        var total = 0.0;
        this.state.rows.forEach((data, index) => {
            total += parseFloat(data[1].props.value)
            searchConfigAlloc.push(parseFloat(data[1].props.value));
          });
        if (total !== 100) {
            this.setState({
                errorMsg : "Allocations must total 100%!"
            });
        } else {
            var updatedConf = this.state.shopConfig;
            updatedConf.searchConfigAlloc = searchConfigAlloc
            this.setState({
                errorMsg : NaN,
                shopConfig : updatedConf
            });
            this.postConfig();
        }
    }

    handleEditButton = (index) => {
        this.setState({
            editSearchConfIndex: index,
            editSearchConf: this.state.shopConfig.shopifySearchConfigs[index]
        })
    }


    handleCloseSearchConfModal = () => {
        this.setState({
            editSearchConfIndex: NaN,
            editSearchConf: NaN
        })
        this.fetchConfig();
    }

    handleSaveSearchConf = () => {
    
        this.setState({
            editSearchConfIndex: NaN,
            editSearchConf: NaN,
        })
        this.postConfig()
        //this.fetchConfig()
    }

    handleConfNameChange = (value) => {
        var conf = this.state.editSearchConf
        conf.name = value

        var rows = this.state.rows
        var row = this.state.rows[this.state.editSearchConfIndex]
        row[0] = value
        this.setState({
            editSearchConf: conf,
            rows: rows
        })
    }

    handleConfLocaleChange = (value) => {
        var conf = this.state.editSearchConf
        conf.searchLocale = value
        this.setState({
            editSearchConf: conf
        })
    }    

    handleConfChangePSTA = (value) => {
        var conf = this.state.editSearchConf
        conf.includeProductSuggesitons = value
        this.setState({
            editSearchConf: conf
        })
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

        let loading;
        if (this.state.loading) {
            loading = <SkeletonBodyText lines={10} />
        }

        let table;
        if (!this.state.loading) {
            table = <DataTable columnContentTypes={[
                'text',
                'text',
                'text',
                'text'
                ]}
                headings={[
                'Name',
                'Traffic %',
                '',
                ''
                ]}
                rows={this.state.rows}
            />
        }

        let modal;
        if (this.state.editSearchConfIndex > -1) {

            const localeOptions  = [
                {label: 'English', value: 'en'},
                {label: 'Spanish', value: 'es'},
            ];

            modal = <Modal
            open={true}
            onClose={this.handleCloseSearchConfModal}
            title="Editing Search Config"
            primaryAction={{
                content: 'Save',
                onAction: this.handleSaveSearchConf,
            }}
            secondaryActions={{
              content: 'Cancel',
              onAction: this.handleCloseSearchConfModal,
            }}
          >
                <Modal.Section>
                    <Stack>
                        <Stack.Item>
                            <FormLayout>
                                <TextField label="Name" value={this.state.editSearchConf.name} onChange={this.handleConfNameChange} />
                                <Checkbox label="Include Product Suggestions in Typeahead" onChange={this.handleConfChangePSTA} checked={this.state.editSearchConf.includeProductSuggesitons} />
                                <Select
                                        label="Locale"
                                        onChange={this.handleConfLocaleChange}
                                        value={this.state.editSearchConf.searchLocale}
                                        options={localeOptions}
                                    />
                            </FormLayout>
                        </Stack.Item>
                    </Stack>
                </Modal.Section>
            </Modal>
        }

        return (
            <Layout.AnnotatedSection
            title="Search Configurations"
            description="Manage global settings and weights for search."
            >
                {message}
                {error}
                {modal}
                <Card sectioned primaryFooterAction={{
                    content: "Update Traffic Allocation",
                    onAction: this.handleUpdateAlloc
                    }}>
                    {loading}
                    {table}
                </Card>
            </Layout.AnnotatedSection>
        );
    }
}

export default SearchConfigGrid