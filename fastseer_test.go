package fastseer

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	conf := LoadConfigFromFile("config-stage.yaml")
	if conf.ShopifyApiKey == "" || conf.ShopifyApiSecret == "" {
		t.Error("Shopify Api Key or Secret wasn't parsed from conf file")
	}

	if len(conf.DbOptions) == 0 || len(conf.IndexingWorkerServices) == 0 {
		t.Error("Emtpy config attribute. Was it loaded?")
	}
	t.Log(conf.DbOptions)
	t.Log(conf.IndexingWorkerServices)
}

func ServerTest(t *testing.T) {

}
