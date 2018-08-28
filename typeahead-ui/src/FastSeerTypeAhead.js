import React, { Component } from 'react';
import Autosuggest from 'react-autosuggest';
import Switch from '../node_modules/react-router-dom/Switch';
import Route from '../node_modules/react-router-dom/Route';
import withRouter from '../node_modules/react-router/Router'
import FastSeerTypeAheadApp from './App';
import { Link } from '../node_modules/react-router-dom';
import close from './close.svg'

// When suggestion is clicked, Autosuggest needs to populate the input
// based on the clicked suggestion. Teach Autosuggest how to calculate the
// input value for every given suggestion.
const getSuggestionValue = suggestion => suggestion.name;

// Use your imagination to render suggestions.
const renderSuggestion = suggestion => (
  <div className="suggestItemWrapper">
    <div className="suggestImage"><img src={suggestion.image}/></div>
    <div className="suggestName">{suggestion.name.substring(0, 50)}...from <strong>${suggestion.price}</strong></div>
  </div>
);

class FastSeerTypeAhead extends React.Component {
  constructor(props) {
    super(props);
    // Autosuggest is a controlled component.
    // This means that you need to provide an input value
    // and an onChange handler that updates this value (see below).
    // Suggestions also need to be provided to the Autosuggest,
    // and they are initially empty because the Autosuggest is closed.
    this.state = {
      value: '',
      suggestions: []
    };
  }

  componentDidMount() {
    this.mounted = true;
  }
  componentWillUnmount() {
    this.mounted = false;
  }

  onChange = (event, { newValue }) => {
    this.setState({
      value: newValue
    });
  };

  // Autosuggest will call this function every time you need to update suggestions.
  // You already implemented this logic above, so just use it.
  onSuggestionsFetchRequested = ({ value }) => {  
    if (value.length < 2) {
      this.setState({
        suggestions: [],
      });
      return;
    }
    fetch(this.props.appDomain + "/search/typeahead?shop=" + this.props.shop + "&q=" + value)
      .then(res => res.json())
      .then(
          (result) => {
            if (this.mounted) {
              this.setState({
                suggestions: result,
              });
            }
          },
      )
      .catch(error => {
        this.setState({
          suggestions: [],
        });
        console.error(error)
      });    
  };

  // Autosuggest will call this function every time you need to clear suggestions.
  onSuggestionsClearRequested = () => {
    this.setState({
      suggestions: []
    });
  };

  onSuggestionSelected = () => {
    console.log(this.timer);
    clearTimeout(this.timer);
    window.fsOpen('/results?q=' + this.state.value);
  }

  onEnterPress = (e) => {
    if (e.key === 'Enter') {
        // exit the text box
        // same as blur
        clearTimeout(this.timer)
        window.fsOpen('/results?q=' + this.state.value);
    }
 }

  onFocusLost = () => {
    window.fsOpen('/')
  }

  render() {
    const { value, suggestions } = this.state;

    // Autosuggest will pass through all these props to the input.
    const inputProps = {
      placeholder: window.searchConfig.placeHolder,
      value,
      type: 'search',
      onChange: this.onChange,
      onBlur: this.onFocusLost,
      id: 'fsInput', 
      onKeyDown: this.onEnterPress,
    };

    return (
      <div className="fa-modal">
        <div className="fa-search">
          <Autosuggest
            suggestions={suggestions}
            onSuggestionsFetchRequested={this.onSuggestionsFetchRequested}
            onSuggestionsClearRequested={this.onSuggestionsClearRequested}
            onSuggestionSelected={this.onSuggestionSelected}
            getSuggestionValue={getSuggestionValue}
            renderSuggestion={renderSuggestion}
            inputProps={inputProps}
            ref={ () => { document.getElementById('fsInput').focus(); } } 
          />
        </div>
        <div className="fa-close">
          <Link to="/">
          <img height="37" width="32" src="https://static.fastseer.com/static/media/close.6f9bd8cd.svg"/>
          </Link>
        </div>
      </div>
    );
  }
}



export default FastSeerTypeAhead;