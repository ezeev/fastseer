package shopify

import (
	"encoding/json"
	"testing"

	"github.com/ezeev/fastseer/search"
)

const appDomain = "https://b76cf0eb.ngrok.io"

func testShopifyConfig() *ShopifyClientConfig {
	return &ShopifyClientConfig{
		AuthResponse: ShopifyAuthResponse{AccessToken: "637c936f5f6a04bcde91f74c082a993c"},
		Shop:         "fastseer-staging.myshopify.com",
		IndexAddress: "172.104.9.135:8983/solr",
	}
}

func testSearchEngine() search.SearchEngine {
	engine, _ := search.NewSearchEngine("solr")
	return engine
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

func TestGetScriptTags(t *testing.T) {
	shop := testShopifyConfig()

	resp, err := GetScriptTagsBySrc(shop, appDomain)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)

}

func TestPostScriptTag(t *testing.T) {

	shop := testShopifyConfig()

	resp, err := PostScriptTag(shop, appDomain)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.ScriptTag)

}

func TestInstallScriptTag(t *testing.T) {

	shop := testShopifyConfig()
	resp, err := InstallShopScriptTag(shop, appDomain)
	if err != nil {
		t.Error(err)
	}

	t.Log(resp)

}
func TestDeleteScriptTags(t *testing.T) {
	// get first
	shop := testShopifyConfig()

	resp, err := GetScriptTagsBySrc(shop, appDomain)
	if err != nil {
		t.Error(err)
	}

	t.Logf("There are %d script tags. Deleting...", len(resp.ScriptTags))

	for _, v := range resp.ScriptTags {
		t.Logf("Deleting script tag ID: %d", v.ID)
		err = DeleteScriptTag(shop, v.ID)
		if err != nil {
			t.Error(err)
		}
	}

	resp, err = GetScriptTagsBySrc(shop, appDomain)
	if err != nil {
		t.Error(err)
	}
	t.Logf("There are now %d script tags", len(resp.ScriptTags))

}

func TestShopifySearches(t *testing.T) {

	shop := testShopifyConfig()

	num, err := NumIndexedProducts(shop, testSearchEngine())
	if err != nil {
		t.Error(err)
	}

	t.Logf("There are %d indexed prodcuts", num)

	//variantes
	num, err = NumIndexedVariants(shop, testSearchEngine())
	if err != nil {
		t.Error(err)
	}

	t.Logf("There are %d indexed variantes", num)

}

func TestGetThemes(t *testing.T) {
	resp, err := GetThemes(testShopifyConfig())
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestPutSearchFormAsset(t *testing.T) {
	// get the themes
	resp, err := GetThemes(testShopifyConfig())
	if err != nil {
		t.Error(err)
	}
	id := resp.Themes[0].ID

	err = PutSearchFormThemeAsset(testShopifyConfig(), id)
	if err != nil {
		t.Error(err)
	}

}

func TestInstallSearchForm(t *testing.T) {
	err := InstallSearchFormThemeAsset(testShopifyConfig())
	if err != nil {
		t.Error(err)
	}
}
