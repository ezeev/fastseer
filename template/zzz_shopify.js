
$(document).ready(function(){


    
    var shopTopSearches = new Bloodhound({
        datumTokenizer: Bloodhound.tokenizers.obj.whitespace('search_txt'),
        queryTokenizer: Bloodhound.tokenizers.whitespace,
        remote: {
        url: '[[.AppDomain]]/search/typeahead/topsearches?shop=[[.Shop]]&q=%%QUERY',
        wildcard: '%%QUERY'
        }
    });

[[ if .Conf.IncludeProductSuggesitons ]]
    var shopTA = new Bloodhound({
        datumTokenizer: Bloodhound.tokenizers.obj.whitespace('title_txt'),
        queryTokenizer: Bloodhound.tokenizers.whitespace,
        //prefetch: '../data/films/post_1960.json', // TODO! populate w/ top searches
        remote: {
        //url: '../data/films/queries/%QUERY.json',
        url: '[[.AppDomain]]/search/typeahead?shop=[[.Shop]]&q=%%QUERY',
        wildcard: '%%QUERY'
        }
    });
[[ end ]]

    $('.launch-typeahead').click(function(){
        $('#fs-searchForm').show('fast')
        //focus
        $('.typeahead').focus();

        return false;
    });

    $('.fs-closeForm').click(function(){
        $('#fs-searchForm').hide('fast');
        return false;
    });

    $('.typeahead').typeahead(null, 
        {
            minLength: 2,
            highlight: true,
            limit: Infinity,
            name: 'shop-ta',
            display: 'search_txt',
            source: shopTopSearches,
            templates: {
                //header: '<div class="ta-section-header"><h3>Top Searches</h3></div>',
                empty: [
                    '<div class="ta-item">',
                        'Unable to find any results',
                    '</div>'
                ].join('\n'),
                suggestion: Handlebars.compile(`
                    <div>
                        {{#if type_s}}
                            <div class="ta-sub">
                                &nbsp;&nbsp;&nbsp;&nbsp;<small>in {{type_s}}</small>
                            </div>
                        {{else}}
                            <div class="ta-item">
                                {{search_txt}}
                            </div>
                        {{/if}}
                    </div>
                    `)
                }
        }
[[ if .Conf.IncludeProductSuggesitons ]]        
        ,
        {
        minLength: 2,
        limit: Infinity,
        highlight: true,
        name: 'shop-ta-products',
        display: 'title_txt',
        source: shopTA,
        templates: {
            header: '<div class="ta-item"><i>Product suggestions:</i></div>',
            empty: [
                '<div class="ta-product">',
                    'Unable to find any results',
                '</div>'
            ].join('\n'),
            suggestion: Handlebars.compile(`
                <div class="ta-product">
                    <div class="ta-img">
                        <img src="{{img_s}}"/>
                    </div>
                    <div class="ta-txt">
                        {{title_txt}}
                    </div>
                </div>
                `)
            }
        }
[[ end ]]        
    );

    $('.typeahead').bind('typeahead:select', function(ev, suggestion) {
        if (suggestion.title_txt) { //product suggestion
            window.location.href = "https://[[ .Shop ]]/search?q=" + suggestion.title_txt;
        } else if (suggestion.search_txt) {
            var url =  "https://[[ .Shop ]]/search?q=" + suggestion.search_txt;
            if (suggestion.type_s) {
                url += '&fq=cat:"' + suggestion.type_s + '"'
            }
            window.location.href = url;
        }
    });
  });  