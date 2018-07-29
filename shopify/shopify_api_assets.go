package shopify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func InstallSearchFormThemeAsset(shop *ShopifyClientConfig) error {

	//what is the active theme?
	themes, err := GetThemes(shop)
	if err != nil {
		return err
	}
	for _, v := range themes.Themes {
		if v.Role == "main" {
			err = PutSearchFormThemeAsset(shop, v.ID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func PutSearchFormThemeAsset(shop *ShopifyClientConfig, themeID int) error {

	html := `

	<!-- FastSeer Typeahead Search -->

	<script src="https://twitter.github.io/typeahead.js/releases/latest/typeahead.bundle.js" defer="defer"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/handlebars.js/4.0.11/handlebars.min.js" defer="defer"></script>	
	<style>
	
	.typeahead {
		height:50px;
	}
		
	
	.ta-section-header{
		background-color: #2471A3;
		color:white;
		width: 500px;
		padding-left:5px;
	}
		
	.ta-product {
		width: 400px;
		padding-top: 0px;
		display: flex;
		vertical-align: middle;
		/*border: 1px solid #D4E6F1;*/
		cursor: pointer;
		
	}
	
	/*div.tt-cursor.ta-item { 
			background-color: yellow;
	}*/
	
		
	.tt-suggestion {
		background-color: white;
	}
	.tt-suggestion.tt-cursor {
		background-color:#EBF5FB;
	}
		
	/*div.tt-cursor { 
			background-color: yellow;
	}*/
	
	.ta-item {
		width: 400px;
		padding-left: 10px;
		cursor: pointer;
	}
	
	.tt-dataset-shop-ta-products {
		background-color: white;
	}
		
	.ta-sub {
		width: 400px;
		padding-left: 10px;
		cursor: pointer;
	}
		
	.tt-dataset-shop-ta {
		padding-top:5px;
		padding-bottom:5px;  
		background-color: white;
	}
	 
	.ta-img {
		float: left;
		width: 60px;
		max-height: 60px;
		vertical-align:middle;
		padding:0px;
		cursor: pointer;
	}
		
	.ta-img img {
		width: 60px;
		padding:0px;
		max-height: 60px;
		max-width: 100%;
	}
		
	.ta-txt {
		width: 100%;
		vertical-align: middle;
		padding:5px;
	}
		
	.ta-txt i {
		font-weight:50;
	}
		
	.clear:after {
			clear: both;
			display: table;
			content: "";
	}
		
	
	 
	.fs-modal {
			display: none; /* Hidden by default */
			position: fixed; /* Stay in place */
			z-index: 1; /* Sit on top */
			left: 0;
			top: 0;
			width: 100%; /* Full width */
			height: 100%; /* Full height */
			overflow: auto; /* Enable scroll if needed */
			background-color: rgb(0,0,0); /* Fallback color */
			background-color: rgba(0,0,0,0.4); /* Black w/ opacity */
	}
	
	.fs-centered {
		position: fixed;
		z-index: 2; /* Sit on top */
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		width: 400px;
		height: 100%;
		padding-top: 50px;
		max-height: 100%;
		overflow-y: auto;
	}
		
	.fs-closeForm {
		font-size: large;
		color:white;
	}
		
	</style>
	
	<div id="fs-searchForm" class="fs-modal">
		<div class="fs-centered" >
			<input class="typeahead" style="width:400px;background-color:#EBF5FB;font-size:18px;color: #17202A;" type="search" name="q" value="{{ search.terms | escape }}" placeholder="{{ 'layout.search_bar.placeholder' | t }}">
			<br/>
			<a href="#" class="fs-closeForm">Close</a>
		</div>
	</div>
	
	<!-- /FastSeer -->
	
	`
	payload := ShopifyAssetPutPayload{}
	payload.Asset.Key = "snippets/fs-search-form.liquid"
	payload.Asset.Value = html

	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("/admin/themes/%d/assets.json", themeID)
	cli, req := httpClient(shop.AuthResponse.AccessToken, shop.Shop, "PUT", path, bytes.NewBuffer(b))
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("Did not receive 200 status code, response body: %s", string(body))
	}
	return nil

}
