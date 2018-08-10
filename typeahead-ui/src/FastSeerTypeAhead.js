import React, { Component } from 'react';
import Autosuggest from 'react-autosuggest';

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
  constructor() {
    super();
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
            this.setState({
              suggestions: result,
            });
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

  render() {
    const { value, suggestions } = this.state;

    // Autosuggest will pass through all these props to the input.
    const inputProps = {
      placeholder: window.placeholder,
      value,
      onChange: this.onChange
    };


    // Finally, render it!
    return (
      <Autosuggest
        suggestions={suggestions}
        onSuggestionsFetchRequested={this.onSuggestionsFetchRequested}
        onSuggestionsClearRequested={this.onSuggestionsClearRequested}
        getSuggestionValue={getSuggestionValue}
        renderSuggestion={renderSuggestion}
        inputProps={inputProps}

      />
    );
  }
}

export default FastSeerTypeAhead