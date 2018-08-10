package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/ezeev/fastseer/search"
	"github.com/ezeev/fastseer/shopify"
)

const testToken = "637c936f5f6a04bcde91f74c082a993c"
const testShop = "fastseer-staging.myshopify.com"
const searchUrl = "172.104.9.135:8983/solr"

func testShopifyConfig() *shopify.ShopifyClientConfig {
	return &shopify.ShopifyClientConfig{
		AuthResponse: shopify.ShopifyAuthResponse{AccessToken: "637c936f5f6a04bcde91f74c082a993c"},
		Shop:         "fastseer-staging.myshopify.com",
		IndexAddress: "172.104.9.135:8983/solr",
	}
}

func testHttpClient() *http.Client {

	return &http.Client{
		Timeout: time.Second * 10,
	}

}

func TestClearAPI(t *testing.T) {
	b, _ := json.Marshal(testShopifyConfig())

	server, err := NewServer(8083, "solr")
	if err != nil {
		t.Error(err)
	}
	go server.Start()

	server.ServerReady()

	cli := testHttpClient()
	req, err := http.NewRequest("DELETE", "http://localhost:8083/api/indexShopify", bytes.NewBuffer(b))
	resp, err := cli.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
}

func TestIndexingServer(t *testing.T) {

	solr, _ := search.NewSearchEngine("solr")

	//clear index first
	config := testShopifyConfig()
	err := solr.DeleteDocuments(config.Shop, config.IndexAddress, "*:*")
	if err != nil {
		t.Error(err)
	}

	b, _ := json.Marshal(testShopifyConfig())

	server, err := NewServer(8083, "solr")
	if err != nil {
		t.Error(err)
	}
	go server.Start()

	server.ServerReady()

	cli := testHttpClient()
	req, err := http.NewRequest("POST", "http://localhost:8083/api/indexShopify", bytes.NewBuffer(b))
	resp, err := cli.Do(req)
	if err != nil {
		t.Error(err)
	}
	var iresp IndexingStatusResponse
	err = json.NewDecoder(resp.Body).Decode(&iresp)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(iresp.Message)
	t.Logf("products to crawl: %d", iresp.ProductsToCrawl)

	time.Sleep(time.Second * 30)

	server.Shutdown()

}
