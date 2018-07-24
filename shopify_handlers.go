package fastseer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

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

		data := ShopifyPageData{
			Config:       s.Config,
			FlashMessage: s.FlashMessage(r),
			Shop:         params.Get("shop"),
			HMac:         params.Get("hmac"),
			Timestamp:    params.Get("timestamp"),
			Locale:       params.Get("locale"),
			Protocol:     params.Get("protocol"),
			ShopConfig:   shopClient,
		}
		/*cookie, _ := r.Cookie(flashMessageCookieName)
		if cookie != nil {
			data.FlashMessage = cookie.Value
		}*/

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
