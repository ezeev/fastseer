package shopify

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ezeev/fastseer/search"
)

// example: https://fastseer-staging.myshopify.com/admin/products.json
// X-Shopify-Access-Token: a9788d44d43105577f0c11ce41e5d97a
// see https://help.shopify.com/en/api/reference/products

type ShopifyApiProductCount struct {
	Count int `json:"count"`
}

func httpClient(token string, shop string, method string, endPoint string) (*http.Client, *http.Request) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	req, _ := http.NewRequest(method, "https://"+shop+endPoint, nil)
	req.Header.Add("X-Shopify-Access-Token", token)
	return netClient, req
}

func GetNumProducts(token string, shop string) (*ShopifyApiProductCount, error) {

	cli, req := httpClient(token, shop, "GET", "/admin/products/count.json")
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var apiResp ShopifyApiProductCount
	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	return &apiResp, err
}

//func CrawlProducts(token string, shop string, pageSize int, sinceID string, engine search.SearchEngine, indexAddress string) error {
func CrawlProducts(shop *ShopifyClientConfig, pageSize int, sinceID string, engine search.SearchEngine) error {

	// make recursive calls until done
	cli, req := httpClient(shop.AuthResponse.AccessToken, shop.Shop, "GET", "/admin/products.json")
	params := req.URL.Query()
	params.Add("limit", strconv.Itoa(pageSize))
	if sinceID != "" {
		params.Add("since_id", sinceID)
	}
	req.URL.RawQuery = params.Encode()

	log.Printf("Shop: %s, Executing query with last ID: %s\n, page size: %d", shop.Shop, sinceID, pageSize)

	resp, err := cli.Do(req)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Shop: %s, ioutil.ReadAll() error: %v\n", shop.Shop, err)
	}

	var products ShopifyApiProductsResponse
	err = json.Unmarshal(data, &products)
	if err != nil {
		log.Printf("Shop: %s, Error serializing json: %s", shop.Shop, err)
	}
	log.Printf("Shop: %s, Received %d products in response", shop.Shop, len(products.Products))
	var numVariants int
	for _, v := range products.Products {
		numVariants += len(v.Variants)
		log.Printf("Product %d has %d variants", v.ID, numVariants)
	}

	if len(products.Products) > 0 {
		lastID := strconv.Itoa(products.Products[len(products.Products)-1].ID)
		log.Printf("Shop: %s, last ID: %s", shop.Shop, lastID)
		time.Sleep(time.Second * 1)

		// send to search
		err = IndexProducts(&products, engine, shop)
		if err != nil {
			return err
		}

		CrawlProducts(shop, pageSize, string(lastID), engine)
	} else {
		log.Printf("Shop: %s, No more products to crawl.", shop.Shop)
	}

	return nil
}
