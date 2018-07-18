package fastseer

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	conf := LoadConfigFromFile("config.yaml")
	if conf.ShopifyApiKey == "" || conf.ShopifyApiSecret == "" {
		t.Error("Shopify Api Key or Secret wasn't parsed from conf file")
	}
	t.Log(conf.DbOptions)
	t.Log(conf.IndexingWorkerServices)
}

func ServerTest(t *testing.T) {

}
