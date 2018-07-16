package shopify

import (
	"encoding/json"
	"testing"
)

func testShopifyConfig() *ShopifyClientConfig {
	return &ShopifyClientConfig{
		AuthResponse: ShopifyAuthResponse{AccessToken: "637c936f5f6a04bcde91f74c082a993c"},
		Shop:         "fastseer-staging.myshopify.com",
		IndexAddress: "172.104.9.135:8983/solr",
	}
}

func TestProductCountApi(t *testing.T) {

	shop := testShopifyConfig()

	res, err := GetNumProducts(shop.AuthResponse.AccessToken, shop.Shop)
	if err != nil {
		t.Error(err)
	}
	if res.Count == 0 {
		t.Error("Count is 0!")
	}
	t.Logf("There are %d products", res.Count)

}

func TestCrawlProducts(t *testing.T) {
	shop := testShopifyConfig()
	b, _ := json.Marshal(shop)
	t.Log(string(b))
	CrawlProducts(shop, 3, "", nil)
}
