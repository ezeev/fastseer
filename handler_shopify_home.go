package fastseer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/ezeev/fastseer/logger"
	"github.com/ezeev/fastseer/shopify"
)

func (s *Server) handleShopifyHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cookie, _ := r.Cookie("_secure_admin_session_id")
		if cookie != nil {
			log.Println(cookie.Value)
		}

		tmpl, err := template.ParseFiles("template/shopify-admin-home.html", "template/shopify-admin-head-includes.html")
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

		type homeData struct {
			FlashMessage string
			Config       *Config
			Shop         string
			Error        string
			Success      string
			HmacAuth     *HmacAuthParams
			NumProducts  int
			NumVariants  int
			ShopConfig   *shopify.ShopifyClientConfig
		}
		//shopClient.SearchConfigs
		data := homeData{
			Config:       s.Config,
			FlashMessage: s.FlashMessage(w, r),
			Shop:         params.Get("shop"),
			Error:        s.FlashError(w, r),
			Success:      params.Get("success"),
			HmacAuth:     NewHmacAuthFromParams(params),
			ShopConfig:   shopClient,
			NumProducts:  numProd,
			NumVariants:  numVar,
		}
		err = tmpl.ExecuteTemplate(w, "head-includes", data)
		if err != nil {
			logger.Error(shop, err.Error())
		}
		err = tmpl.ExecuteTemplate(w, "layout", data)
		if err != nil {
			logger.Error(shop, err.Error())
		}
	}
}
