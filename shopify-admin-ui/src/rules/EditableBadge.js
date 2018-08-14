import React from 'react';
import {
  Badge,
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
            tempValue: null,
            deleted: false,
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

    handleDelete = () => {
        this.setState({
            value: '',
        });
        //this.props.valueChange(this.props.field, this.props.index, this.state.value)*/
        this.props.onDelete(this.props.field, this.props.index)
    }

    handleOnChange = (item) => {
        this.setState({
            value: item,
        })
    }

    handleKeyPress = (e) => {
        if (e.key === 'Enter') {
            // exit the text box
            // same as blur
            this.handleBlur();
        }
     }

    render() {
        let content;
        if (this.state.deleted) {
            return (<span></span>);
        }

        if (this.state.editMode) {
            content = <div onKeyDown={this.handleKeyPress}>
                        <TextField 
                            tabIndex="0"
                            autoFocus
                            onChange={this.handleOnChange}
                            value={this.state.value}
                            onBlur={this.handleBlur}
                            ></TextField></div>
        } else {
            if (this.state.value.length <= 0) {
                content = <Link onClick={this.handleBadgeClick}>set</Link>
            } else {
                content = <Badge><Link onClick={this.handleBadgeClick}>{this.state.value}</Link><Link onClick={this.handleDelete}><Icon source="cancel" color="skyDark" style={{height:"1rem", maxHeight:"1rem"}} /></Link></Badge>
            }
        }
        return (
            <span>
                {content}
            </span>
        );
    }
}
export default EditableBadge
