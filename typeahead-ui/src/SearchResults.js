import React, { Component } from 'react';
import queryString from 'query-string'
import {withRouter} from 'react-router-dom'

class SearchResults extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            results: null,
            loading: true,
        }
    }

    componentDidMount() {        
        this.mounted = true;
        var params = new URLSearchParams(this.props.location.search);
        var start = '0';
        var q = params.get('q');
        if (params.has('start')) {
            start = params.get('start');
        }

        this.fetchResults(q, start)
    }
    componentWillUnmount(q) {
        this.mounted = false;
    }


    onNextPage = (resp) => {        
        const nextStart = resp.start + window.searchConfig.pageSize;        
        window.fsOpen('/results?q=' + this.state.q + '&start=' + nextStart);
        this.fetchResults(this.state.q, nextStart);
    }

    onPrevPage = (resp) => {
        const prevStart = resp.start - window.searchConfig.pageSize;      
        window.fsOpen('/results?q=' + this.state.q + '&start=' + prevStart);
        this.fetchResults(this.state.q, prevStart);  
    }

    roundFloat(f) {
        return f.toFixed(2);
    }

    shortenText(t, len) {
        const tokens = t.split(' ');
        if (tokens.length <= len) {
            return t;
        }
        var str = '';
        const firstN = tokens.slice(0, len);
        str = firstN.join(' ');
        const last5 = tokens.slice(Math.max(tokens.length - 3, 1));
        str += ' ..... ' + last5.join(' ')

        /*for (var i=0;i<last5.length;i++) {
            str += last5[i] + ' ';
        }*/
        return str;
    }

    fetchResults(q, start) {

        this.setState({loading: true,})
        console.log("getting results for " + q)    
        fetch(this.props.appDomain + "/api/v1/search?shop=" 
            + this.props.shop 
            + "&q=" + q
            + "&rows=" + window.searchConfig.pageSize
            + "&start=" + start
        )
        .then(res => res.json())
        .then(
            (result) => {
                if (this.mounted) {
                    this.setState({
                        q: q,
                        results: result,
                        loading: false,
                    });                    
                    console.log(this.state.results)
                }
            },
        )
        .catch(error => {
            this.setState({
                results: {},
                loading: false,
            });
            console.error(error)
        });    
    }

    render() {
        console.log("render called")
        const { match, location, history } = this.props

        let loading;
        if (this.state.loading) {
            loading = <div className="ui active loader"></div>
        }

        let message;
        if (this.state.results) {
            message = <div className="resultsHeader">Showing {this.state.results.response.start+1} - {(window.searchConfig.pageSize+this.state.results.response.start)} of {this.state.results.response.numFound} results for "<i>{this.state.q}</i>"</div>
        }
    
        let paging;
        if (this.state.results) {

            let prev;
            if (this.state.results.response.start > 0) {
                prev = <span><span className="pagingLink" onClick={() => this.onPrevPage(this.state.results.response)}><i className="angle left icon"></i> Previous</span> |</span>
            }

            let next;
            if (this.state.results.response.start < this.state.results.response.numFound) {
                next = <span className="pagingLink" onClick={() => this.onNextPage(this.state.results.response)}> Next <i className="angle right icon"></i></span>
            }

            paging = <div className="right aligned one column row">
                        <div className="column">
                             {prev}                            
                             {next}
                        </div>
                    </div>
        }

        let results;
        if (this.state.results) {
            results = this.state.results.response.docs.map((item) =>
                    <div className="column" key={item.productId}>
                        <div className="resultThumb">
                            <img src={item.image} />
                        </div>
                        <div>{this.shortenText(item.title, 8)}</div>
                        <div className="resultPrice">${this.roundFloat(item.price)}</div>
                    </div>   
            );
        }
        return(
            <div className="ui container">
                {loading}
                {message}
                <div className="ui stackable three column grid">
                    {paging}
                    {results}
                </div>
            </div>
        )
    }

}

export default withRouter(SearchResults);