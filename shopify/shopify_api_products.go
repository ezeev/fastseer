package shopify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/ezeev/fastseer/logger"
	"github.com/ezeev/fastseer/search"
)

func GetNumProducts(token string, shop string) (*ShopifyApiProductCount, error) {

	cli, req := httpClient(token, shop, "GET", "/admin/products/count.json", nil)
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
//func CrawlProducts(shop *ShopifyClientConfig, pageSize int, sinceID string, engine search.SearchEngine) error {

func CrawlProducts(shop *ShopifyClientConfig, pageSize int, pageNum int, engine search.SearchEngine) error {

	// make recursive calls until done
	cli, req := httpClient(shop.AuthResponse.AccessToken, shop.Shop, "GET", "/admin/products.json", nil)
	params := req.URL.Query()
	params.Add("limit", strconv.Itoa(pageSize))
	if pageNum != 0 {
		params.Add("page", strconv.Itoa(pageNum))
	}
	req.URL.RawQuery = params.Encode()

	logger.Info(shop.Shop, fmt.Sprintf("Executing query with page: %d\n, page size: %d", pageNum, pageSize))

	resp, err := cli.Do(req)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(shop.Shop, err.Error())
	}

	var products ShopifyApiProductsResponse
	err = json.Unmarshal(data, &products)
	if err != nil {
		logger.Error(shop.Shop, err.Error())
	}
	logger.Info(shop.Shop, fmt.Sprintf("Received %d products in response", len(products.Products)))

	var numVariants int
	for _, v := range products.Products {
		numVariants += len(v.Variants)
	}

	if len(products.Products) > 0 {
		lastID := strconv.Itoa(products.Products[len(products.Products)-1].ID)
		logger.Info(shop.Shop, "last ID: "+lastID)
		time.Sleep(time.Second * 1)

		// send to search
		err = IndexProducts(&products, engine, shop)
		if err != nil {
			return err
		}
		CrawlProducts(shop, pageSize, pageNum+1, engine)
	} else {
		logger.Info(shop.Shop, "No more products to crawl")
	}
	return nil
}
