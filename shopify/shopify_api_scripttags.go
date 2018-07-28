package shopify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

func jsSrc(shop string, appDomain string) string {
	return fmt.Sprintf("%s/shopify/shop.js?shop=%s", appDomain, shop)
}

func GetScriptTagsBySrc(shop *ShopifyClientConfig, appDomain string) (*ShopifyScriptTagsResponse, error) {
	cli, req := httpClient(shop.AuthResponse.AccessToken, shop.Shop, "GET", "/admin/script_tags.json", nil)

	params := req.URL.Query()
	//params.Set("src", "https://b76cf0eb.ngrok.io/shopify/shop.js?shop=fastseer-staging.myshopify.com")
	params.Set("src", jsSrc(shop.Shop, appDomain))
	req.URL.RawQuery = params.Encode()
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var apiResp ShopifyScriptTagsResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	return &apiResp, err
}

func InstallShopScriptTag(shop *ShopifyClientConfig, appDomain string) (*ShopifyPostScriptTagResponse, error) {

	// Step 1, see if there are any script tags matching our source
	// GetScriptTagsBySrc will generate the src string for us (if the app needs more than one script
	// in the future, then this could be parameterized)
	tags, err := GetScriptTagsBySrc(shop, appDomain)
	if err != nil {
		return nil, err
	}

	// Step 2, delete any matching script tags (we want to replace)
	for _, v := range tags.ScriptTags {
		err := DeleteScriptTag(shop, v.ID)
		if err != nil {
			return nil, err
		}
	}

	// Step 3, we can now assume any matching script tags are gone, post it
	return PostScriptTag(shop, appDomain)

}

func DeleteScriptTag(shop *ShopifyClientConfig, tagId int) error {
	cli, req := httpClient(shop.AuthResponse.AccessToken, shop.Shop, "DELETE", "/admin/script_tags/"+strconv.Itoa(tagId)+".json", nil)
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("Deleting script tag failed with status code: %d, body: %s", resp.StatusCode, string(body))
	}
	return nil
}

func PostScriptTag(shop *ShopifyClientConfig, appDomain string) (*ShopifyPostScriptTagResponse, error) {
	js := fmt.Sprintf(`
		{
			"script_tag": {
				"event": "onload",
				"src": "%s"
			}
		}`, jsSrc(shop.Shop, appDomain))
	cli, req := httpClient(shop.AuthResponse.AccessToken, shop.Shop, "POST", "/admin/script_tags.json", bytes.NewBuffer([]byte(js)))

	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var apiResp ShopifyPostScriptTagResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	return &apiResp, err

}
