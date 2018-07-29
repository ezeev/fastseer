package fastseer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/schema"

	"github.com/ezeev/fastseer/logger"

	"github.com/ezeev/fastseer/shopify"
)

const flashMessageCookieName = "msg"

func (s *Server) authorizeShopifyHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		/*
			These params must be passed for auth to succeed:
				hmac
				locale
				shop
				timestamp
		*/
		if !shopify.AuthenticateShopifyRequest(params, s.Config.ShopifyApiSecret) {
			fmt.Fprintf(w, messageInvalidRequest)
			return
		}
		h(w, r)
	}
}

type ShopifyPageData struct {
	FlashMessage string
	Config       *Config
	Shop         string
	HMac         string
	Timestamp    string
	Locale       string
	Protocol     string
	NumProducts  int
	NumVariants  int
	ShopConfig   *shopify.ShopifyClientConfig
}

const messageInvalidRequest = "Invalid request. Please contact support."

func (s *Server) handleShopifyHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tmpl, err := template.ParseFiles("template/home.html")
		if err != nil {
			panic(err)
		}

		params := r.URL.Query()

		shop := params.Get("shop")

		shopClient, err := shopify.ShopClientConfig(shop, s.ClientsStore)
		if err != nil {
			logger.Error(shop, "Unable to load config: %s"+err.Error())
			fmt.Fprintf(w, "Unable to load config for shop.")
			return
		}

		numProd, err := shopify.NumIndexedProducts(shopClient, s.Search)
		if err != nil {
			logger.Error(shop, err.Error())
		}
		numVar, err := shopify.NumIndexedVariants(shopClient, s.Search)
		if err != nil {
			logger.Error(shop, err.Error())
		}

		data := ShopifyPageData{
			Config:       s.Config,
			FlashMessage: s.FlashMessage(r),
			Shop:         params.Get("shop"),
			HMac:         params.Get("hmac"),
			Timestamp:    params.Get("timestamp"),
			Locale:       params.Get("locale"),
			Protocol:     params.Get("protocol"),
			ShopConfig:   shopClient,
			NumProducts:  numProd,
			NumVariants:  numVar,
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			panic(err)
		}
	}
}

func (s *Server) handleBuildIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		shop := params.Get("shop")
		shopClient, _ := shopify.ShopClientConfig(shop, s.ClientsStore)

		// send job to indexing server
		indexWorker := s.Config.IndexingWorkerServices[0]
		cli := &http.Client{
			Timeout: time.Second * 10,
		}

		b, err := json.Marshal(shopClient)
		if err != nil {
			logger.Error(shop, err.Error())
		}
		req, err := http.NewRequest("POST", indexWorker+s.Config.IndexingWorkerEndpoint, bytes.NewBuffer(b))
		if err != nil {
			logger.Error(shop, err.Error())
		}
		resp, err := cli.Do(req)
		if err != nil {
			logger.Error(shop, err.Error())
		}
		defer resp.Body.Close()

		s.SetFlashMessage(w, "Builing Index...")

		redir := r.Referer()
		http.Redirect(w, r, redir, 307)
	}
}

func (s *Server) handleClearIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		shop := params.Get("shop")
		shopClient, _ := shopify.ShopClientConfig(shop, s.ClientsStore)

		// send job to indexing server
		indexWorker := s.Config.IndexingWorkerServices[0]
		cli := &http.Client{
			Timeout: time.Second * 10,
		}

		b, err := json.Marshal(shopClient)
		if err != nil {
			logger.Error(shop, err.Error())
		}
		req, err := http.NewRequest("DELETE", indexWorker+s.Config.IndexingWorkerEndpoint, bytes.NewBuffer(b))
		if err != nil {
			logger.Error(shop, err.Error())
		}
		resp, err := cli.Do(req)
		if err != nil {
			logger.Error(shop, err.Error())
		}
		defer resp.Body.Close()

		s.SetFlashMessage(w, "Clearing Index...")

		redir := r.Referer()
		http.Redirect(w, r, redir, 307)

	}
}

func (s *Server) handleReInstallSearchForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		shop := params.Get("shop")
		shopClient, _ := shopify.ShopClientConfig(shop, s.ClientsStore)

		err := shopify.InstallSearchFormThemeAsset(shopClient)
		if err != nil {
			logger.Error(shop, err.Error())
			s.SetFlashMessage(w, err.Error())
		} else {
			s.SetFlashMessage(w, "Re-installed Search Form on active theme.")
		}

		redir := r.Referer()
		http.Redirect(w, r, redir, 307)
	}
}

func (s *Server) handleUpdateSearchConfig() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		shop := params.Get("shop")

		//remove auth params
		/*params.Del("locale")
		params.Del("hmac")
		params.Del("timestamp")
		params.Del("shop")*/

		conf := new(shopify.ShopifySearchConfig)
		decoder := schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)

		err := decoder.Decode(conf, params)
		if err != nil {
			logger.Error(shop, err.Error())
		}

		// fields not currently set by the form:
		conf.Name = "Default Search Config"
		conf.IsActive = true

		//update shop conf
		var shopConf shopify.ShopifyClientConfig
		err = s.ClientsStore.Get(shop, &shopConf)
		if err != nil {
			logger.Error(shop, err.Error())
		}
		//only one config for now
		shopConf.SearchConfigs = make([]*shopify.ShopifySearchConfig, 1)
		shopConf.SearchConfigs[0] = conf

		// save
		err = s.ClientsStore.Put(shop, shopConf)
		if err != nil {
			logger.Error(shop, err.Error())
		}

		// clear cache
		shopConfigCache.Delete(shop)

		s.SetFlashMessage(w, "Updated Search Config")

		redir := r.Referer()
		http.Redirect(w, r, redir, 307)

	}
}
