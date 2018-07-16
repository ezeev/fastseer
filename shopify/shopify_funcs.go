package shopify

import (
	"strconv"
	"strings"
	"time"

	"github.com/ezeev/solrg"

	"github.com/ezeev/fastseer/search"
)

func CreateClientCollections(searchEngine search.SearchEngine, searchAddr string, shop string) {
	collectionOpts := map[string]string{"numShards": "1", "replicationFactor": "1"}
	searchEngine.CreateIndex(shop, searchAddr, collectionOpts)

	time.Sleep(time.Second * 3)

	searchEngine.CreateIndex(shop+"_analytics", searchAddr, collectionOpts)

	time.Sleep(time.Second * 3)
}

func IndexProducts(productBatch *ShopifyApiProductsResponse, searchEngine search.SearchEngine, config *ShopifyClientConfig) error {

	// index fields

	docs := solrg.NewSolrDocumentCollection()

	for _, product := range productBatch.Products {

		productTitle := product.Title
		productID := product.ID
		productType := product.ProductType
		productTags := strings.Split(product.Tags, ", ")
		productImage := product.Image.Src

		// variants
		for _, variant := range product.Variants {
			doc := solrg.NewSolrDocument("")
			id := variant.ID
			variantTitle := variant.Title
			variantPrice := variant.Price
			variantSku := variant.Sku
			variantKeywords := productTitle + " " + variant.Title + " " + product.Tags

			doc.SetField("id", []string{strconv.Itoa(id)})
			doc.SetField("productTitle_txt_en", []string{productTitle})
			doc.SetField("productId_s", []string{strconv.Itoa(productID)})
			doc.SetField("productType_s", []string{productType})
			doc.SetField("productTags_ss", productTags)
			doc.SetField("productImage_s", []string{productImage})
			doc.SetField("variantTitle_txt_en", []string{variantTitle})
			doc.SetField("variantPrice_f", []string{variantPrice})
			doc.SetField("variantSku_s", []string{variantSku})
			doc.SetField("variantKeywords_txt_en", []string{variantKeywords})
			docs.AddDoc(doc)
		}

	}
	return searchEngine.IndexDocuments(config.Shop, config.IndexAddress, &docs)

}
