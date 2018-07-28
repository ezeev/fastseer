package fastseer

import (
	"fmt"
	"net/http"
)

func (s *Server) handleShopifyJs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		shop := params.Get("shop")

		js := `
	   $(document).ready(function(){


		// Defining the local dataset
		var cars = ['Audi', 'BMW', 'Bugatti', 'Ferrari', 'Ford', 'Lamborghini', 'Mercedes Benz', 'Porsche', 'Rolls-Royce', 'Volkswagen'];
		
		// Constructing the suggestion engine
		var cars = new Bloodhound({
			datumTokenizer: Bloodhound.tokenizers.whitespace,
			queryTokenizer: Bloodhound.tokenizers.whitespace,
			local: cars
		});
		
		// Initializing the typeahead
		$('.typeahead').typeahead({
			hint: true,
			highlight: true, /* Enable substring highlighting */
			minLength: 1 /* Specify minimum characters required for showing result */
		},
		{
			name: 'cars',
			source: cars
		});
	});  
		`

		js = `
		$(document).ready(function(){


			/* MULTI SET TEST */
			// Defining the local dataset
			var carsArr = ['Audi', 'BMW', 'Bugatti', 'Ferrari', 'Ford', 'Lamborghini', 'Mercedes Benz', 'Porsche', 'Rolls-Royce', 'Volkswagen'];
			
			// Constructing the suggestion engine
			var carsSource = new Bloodhound({
				datumTokenizer: Bloodhound.tokenizers.whitespace,
				queryTokenizer: Bloodhound.tokenizers.whitespace,
				local: carsArr
			});
			/* END */



			var shopTopSearches = new Bloodhound({
				datumTokenizer: Bloodhound.tokenizers.obj.whitespace('search_txt'),
				queryTokenizer: Bloodhound.tokenizers.whitespace,
				remote: {
				url: '` + s.Config.AppDomain + `/search/typeahead/topsearches?shop=` + shop + `&q=%%QUERY',
				wildcard: '%%QUERY'
				}
			});

			var shopTA = new Bloodhound({
				datumTokenizer: Bloodhound.tokenizers.obj.whitespace('title_txt'),
				queryTokenizer: Bloodhound.tokenizers.whitespace,
				//prefetch: '../data/films/post_1960.json', // TODO! populate w/ top searches
				remote: {
				//url: '../data/films/queries/%QUERY.json',
				url: '` + s.Config.AppDomain + `/search/typeahead?shop=` + shop + `&q=%%QUERY',
				wildcard: '%%QUERY'
				}
			});
			
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
				limit: 3,
				highlight: true,
				name: 'shop-ta',
				display: 'title_txt',
				source: shopTA,
				templates: {
					header: '<div class="ta-section-header"><h3>Products</h3></div>',
					empty: [
						'<div class="ta-item">',
							'Unable to find any results',
						'</div>'
					].join('\n'),
					suggestion: Handlebars.compile(` + "`" + `
						<div class="ta-item">
							<div class="ta-img">
								<img src="{{img_s}}"/>
							</div>
							<div class="ta-txt">
								{{title_txt}} 
							</div>
						</div>
					` + "`" + `)
					}
				},
				{
				minLength: 2,
				highlight: true,
				name: 'shop-ta2',
				display: 'title_txt',
				source: shopTopSearches,
				templates: {
					header: '<div class="ta-section-header"><h3>Top Searches</h3></div>',
					empty: [
						'<div class="ta-item">',
							'Unable to find any results',
						'</div>'
					].join('\n'),
					suggestion: Handlebars.compile(` + "`" + `
						<div class="ta-item">
							<div class="ta-txt">
								<strong>{{search_txt}}</strong>
								{{#each types_ss}}
										<br>
										&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; <i>in {{this}}</i>

								{{/each}}			
							</div>
						</div>
					` + "`" + `)
					}
				}
			);

			$('.typeahead').bind('typeahead:select', function(ev, suggestion) {
				window.location.href = "https://` + shop + `/search?q=" + suggestion.title_txt + "&fq=productType_s:" + suggestion.productType_s;
			});


		  });  
		
		`
		fmt.Fprintf(w, js)
	}
}
