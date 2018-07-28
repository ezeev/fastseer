package shopify

import "encoding/json"

func GetThemes(shop *ShopifyClientConfig) (*ShopifyGetThemesResponse, error) {

	cli, req := httpClient(shop.AuthResponse.AccessToken, shop.Shop, "GET", "/admin/themes.json", nil)

	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var apiResp ShopifyGetThemesResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	return &apiResp, err

}
