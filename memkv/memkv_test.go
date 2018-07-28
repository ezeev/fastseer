package memkv

import (
	"testing"
	"time"

	"github.com/ezeev/fastseer/shopify"
)

func testShopifyConfig() *shopify.ShopifyClientConfig {
	return &shopify.ShopifyClientConfig{
		AuthResponse: shopify.ShopifyAuthResponse{AccessToken: "637c936f5f6a04bcde91f74c082a993c"},
		Shop:         "fastseer-staging.myshopify.com",
		IndexAddress: "172.104.9.135:8983/solr",
	}
}

func TestMemKV(t *testing.T) {

	conf := testShopifyConfig()

	kv := New(10, 1)

	kv.Put("test1", *conf)

	v := kv.Get("test1")
	t.Log(v)

	//now sleep
	time.Sleep(time.Second * 3)

	v = kv.Get("test1")
	if v != nil {
		t.Error("test1 is supposed to be expired!")
	}

	t.Log(v)
}
