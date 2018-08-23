package shopify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func InstallSearchFormThemeAsset(shop *ShopifyClientConfig, appDomain string) error {

	//what is the active theme?
	themes, err := GetThemes(shop)
	if err != nil {
		return err
	}
	for _, v := range themes.Themes {
		if v.Role == "main" {
			err = PutSearchFormThemeAsset(shop, v.ID, appDomain)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func PutSearchFormThemeAsset(shop *ShopifyClientConfig, themeID int, appDomain string) error {

	html := `
		<!-- FAST SEER SEARCH FORM -->
		<style>
		.fa-modal {
			position: fixed; /* Stay in place */
			z-index: 1; /* Sit on top */
			left: 0;
			top: 0 !important;
			width: 100%%; /* Full width */
			height: 100%%; /* Full height */
			overflow: auto; /* Enable scroll if needed */
			background-color: rgb(0,0,0); /* Fallback color */
			background-color: rgba(0,0,0,0.4); /* Black w/ opacity */
		}
		
		.fa-search {
			position: fixed; /* Stay in place */
			z-index: 2; /* Sit on top */
			left: 50%%;
			margin-left: -140px;
			top:5%% !important;
		}
		
		.fa-close {
			position: fixed; /* Stay in place */
			z-index: 3; /* Sit on top */
			left: 50%%;
			margin-left: -180px;
			top:6%% !important;
			cursor: pointer;
		}
		
		
		.react-autosuggest__input {
		  all: initial;
		  width: 240px !important;
		  height: 30px !important;
		  padding: 10px 20px !important;
		  background-color: white !important;
		  font-family: 'Open Sans', sans-serif !important;
		  font-weight: 300 !important;
		  font-size: 16px !important;
		  border: 1px solid #aaa !important;
		  border-radius: 4px !important;
		  -webkit-appearance: none !important;
		  }
	  
		  .react-autosuggest__input input {
			  background-color: white;
		}
	  
		  
		.react-autosuggest__input--focused {
			outline: none;
		}
		
		.react-autosuggest__input::-ms-clear {
			display: none;
		}
		
		.react-autosuggest__input--open {
			border-bottom-left-radius: 0;
			border-bottom-right-radius: 0;
		}
		
		.react-autosuggest__suggestions-container {
			display: none;
		}
		
		.react-autosuggest__suggestions-container--open {
			display: block;
			position: fixed;
			left: 50%%;
			margin-left: -140px;
			/*left: 37%%;*/
			/*z-index: 3;*/
			width: 280px;
			border: 1px solid #aaa;
			background-color: #fff;
			font-family: 'Open Sans', sans-serif;
			font-weight: 300;
			font-size: 16px;
			border-bottom-left-radius: 4px;
			border-bottom-right-radius: 4px;
		}
		
		.react-autosuggest__suggestions-list {
			margin: 0;
			padding: 0;
			list-style-type: none;
		}
		
		.react-autosuggest__suggestion {
			cursor: pointer;
			padding: 10px 10px;
		}
		
		.react-autosuggest__suggestion--highlighted {
			background-color: #ddd;
		}
		
		.suggestItemWrapper {
			margin: auto;
			clear: both;
		}
		
		.suggestImage {
			float: left;
			width: 18%%;
		}
		
		.suggestImage img {
			max-width: 40px;
		}
		
		.suggestName {
			margin-left: 22%%;
			font-size: 0.8pc;
			color: #000000;
		}
							
			
  
		</style>

		<script type="text/javascript">
			var appDomain="%s";
			var shop="%s";
			var placeholder="What are you looking for?";

			function fsOpen(h){
				var url = location.href;               
				location.href = "#"+h;
				console.log("going to: " + location.href);
			  }
		</script>
		<div id="fs-type-ahead"></div>
		<script type="text/javascript" src="https://static.fastseer.com/static/js/main.1eab4c0d.js"></script>
		<!-- / END FS SEARCH -->
	`
	payload := ShopifyAssetPutPayload{}
	payload.Asset.Key = "snippets/fs-search-form.liquid"
	payload.Asset.Value = fmt.Sprintf(html, appDomain, shop.Shop)

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
