import React from 'react';
import {
  TextContainer,
  Layout,
  Card,
  Stack,
  Badge,
  Caption,
  Heading,
  Subheading,
  Icon,
  TextField,
  Link,
} from '@shopify/polaris';

class EditableBadge extends React.Component {

    // Constructor is the only place you can assign state
    constructor(props) {
        super(props);
        this.state = {
            editMode: false,
            value: '',
        }
    }

    componentDidMount() {
        this.setState({
            value: this.props.value,
        });
    }
  
    componentWillUnmount() {
    }

    handleBadgeClick = (item) => {
        this.setState({
            editMode: true,
        });
    }
    
    handleBlur = (item) => {
        this.setState({
            editMode: false,
        })
        this.props.valueChange(this.props.field, this.props.index, this.state.value)
    }

    handleOnChange = (item) => {
        this.setState({
            value: item,
        })
    }

    render() {

        let content;
        if (this.state.editMode) {
            content = <TextField 
                        autoFocus
                        onChange={this.handleOnChange}
                        value={this.state.value}
                        onBlur={this.handleBlur}></TextField>
        } else {
            content = <Link onClick={this.handleBadgeClick}><Badge>{this.state.value}</Badge></Link>
        }
        return (
            <span>
                {content}
            </span>
        );
    }
}
export default EditableBadge
