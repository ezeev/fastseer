package shopify

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/ezeev/fastseer/storage"

	"github.com/ezeev/solrg"

	"github.com/ezeev/fastseer/logger"
	"github.com/ezeev/fastseer/search"
)

const tagLimit = 10

func httpClient(token string, shop string, method string, endPoint string, payLoad io.Reader) (*http.Client, *http.Request) {
	var netClient = &http.Client{
		Timeout: time.Second * 30,
	}
	req, _ := http.NewRequest(method, "https://"+shop+endPoint, payLoad)
	req.Header.Add("X-Shopify-Access-Token", token)
	req.Header.Set("Content-Type", "application/json")
	return netClient, req
}

func CreateClientCollections(searchEngine search.SearchEngine, searchAddr string, shop string) error {
	collectionOpts := map[string]string{"numShards": "1", "replicationFactor": "1"}

	err := searchEngine.CreateIndex(shop, searchAddr, collectionOpts)
	if err != nil {
		return err
	}

	time.Sleep(time.Millisecond * 500)

	err = searchEngine.CreateIndex(shop+"_analytics", searchAddr, collectionOpts)
	if err != nil {
		return err
	}

	time.Sleep(time.Millisecond * 500)

	err = searchEngine.CreateIndex(shop+"_typeahead", searchAddr, collectionOpts)
	if err != nil {
		return err
	}

	time.Sleep(time.Millisecond * 500)

	err = searchEngine.CreateIndex(shop+"_rules", searchAddr, collectionOpts)
	if err != nil {
		return err
	}

	return nil

}

func AuthenticateShopifyRequest(params url.Values, secretKey string) bool {

	hmac := params.Get("hmac")
	locale := params.Get("locale")
	shop := params.Get("shop")
	timestamp := params.Get("timestamp")

	messageString := "locale=%s&protocol=https://&shop=%s&timestamp=%s"
	message := fmt.Sprintf(messageString, locale, shop, timestamp)
	return VerifyRequest(hmac, message, secretKey)
}

func VerifyRequest(expectedHMAC, message, sharedSecret string) bool {
	h := hmac.New(sha256.New, []byte(sharedSecret))
	h.Write([]byte(message))
	return hmac.Equal([]byte(expectedHMAC), []byte(hex.EncodeToString(h.Sum(nil))))
}

func IndexProducts(productBatch *ShopifyApiProductsResponse, searchEngine search.SearchEngine, config *ShopifyClientConfig) error {

	// index fields

	docsPrimary := solrg.NewSolrDocumentCollection()
	docsTypeAhead := solrg.NewSolrDocumentCollection()

	for _, product := range productBatch.Products {

		productTitle := product.Title
		productID := product.ID
		idStr := strconv.Itoa(productID)
		productType := product.ProductType
		productTags := strings.Split(product.Tags, ", ")
		if len(productTags) > tagLimit {
			productTags = productTags[0:tagLimit]
		}
		productImage := product.Image.Src

		//build typeahead doc
		log.Println(productType)
		typeAheadDoc := solrg.NewSolrDocument("product-" + strconv.Itoa(productID))
		typeAheadDoc.SetField("productType_s", []string{productType})
		typeAheadDoc.SetField("type_s", []string{"product"})
		typeAheadDoc.SetField("id", []string{idStr})
		typeAheadDoc.SetField("title_t", []string{productTitle})
		typeAheadDoc.SetField("img_s", []string{productImage})

		lowestPrice := 99999.99

		// variants
		for _, variant := range product.Variants {
			doc := solrg.NewSolrDocument("")
			id := variant.ID
			variantTitle := variant.Title
			variantPrice := variant.Price
			variantSku := variant.Sku
			variantKeywords := productTitle + " " + variant.Title + " " + strings.Join(productTags, " ")

			p, err := strconv.ParseFloat(variantPrice, 32)
			if err == nil && p < lowestPrice {
				lowestPrice = p
			}

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
			docsPrimary.AddDoc(doc)
		}
		typeAheadDoc.SetField("from_price_f", []string{fmt.Sprintf("%f", lowestPrice)})
		docsTypeAhead.AddDoc(typeAheadDoc)

	}
	//typeahead
	err := searchEngine.IndexDocuments(config.Shop+"_typeahead", config.IndexAddress, &docsTypeAhead)
	if err != nil {
		return err
	}

	return searchEngine.IndexDocuments(config.Shop, config.IndexAddress, &docsPrimary)
}

// GetPermanentAccessToken returns the shop name and permanent access token respectively
func GetPermanentAccessToken(shop string, apiKey, apiSecret string, code string) ShopifyAuthResponse {

	tokenUrl := fmt.Sprintf("https://%s/admin/oauth/access_token", shop)
	v := url.Values{}
	v.Set("client_id", apiKey)
	v.Set("client_secret", apiSecret)
	v.Set("code", code)
	s := v.Encode()
	req, err := http.NewRequest("POST", tokenUrl, strings.NewReader(s))
	if err != nil {
		logger.Error(shop, "Unable to build request: "+err.Error())
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	c := &http.Client{}

	logger.Info(shop, "requesting token at "+tokenUrl)

	resp, err := c.Do(req)
	if err != nil {
		logger.Error(shop, "executing request: "+err.Error())
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(shop, err.Error())
	}

	var tokenResp ShopifyAuthResponse
	err = json.Unmarshal(data, &tokenResp)
	if err != nil {
		logger.Error(shop, err.Error())
	}
	return tokenResp
}

func ShopClientConfig(shop string, s storage.Storage) (*ShopifyClientConfig, error) {
	var shopClient ShopifyClientConfig
	err := s.Get(shop, &shopClient)
	if err != nil {
		return nil, err
	}
	return &shopClient, nil
}

func NumIndexedProducts(shop *ShopifyClientConfig, engine search.SearchEngine) (int, error) {

	//fq={!collapse%20field=productId_s}
	q := &solrg.SolrParams{
		Q:    "*:*",
		Rows: "0",
		Fq:   []string{"{!collapse field=productId_s}"},
	}

	resp, err := engine.Query(shop.Shop, shop.IndexAddress, q)
	if err != nil {
		return 0, err
	}

	solrResp, ok := resp.(*solrg.SolrSearchResponse)
	if ok {
		return solrResp.Response.NumFound, nil
	} else {
		return 0, fmt.Errorf("Type assertion to *solrg.SolrSearchResponse fails")
	}
	return 0, err
}

func NumIndexedVariants(shop *ShopifyClientConfig, engine search.SearchEngine) (int, error) {
	q := &solrg.SolrParams{
		Q:    "*:*",
		Rows: "0",
	}

	resp, err := engine.Query(shop.Shop, shop.IndexAddress, q)
	if err != nil {
		return 0, err
	}

	solrResp, ok := resp.(*solrg.SolrSearchResponse)
	if ok {
		return solrResp.Response.NumFound, nil
	} else {
		return 0, fmt.Errorf("Type assertion to *solrg.SolrSearchResponse fails")
	}
	return 0, err
}
