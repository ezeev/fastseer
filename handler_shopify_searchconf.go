package fastseer

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/ezeev/fastseer/logger"
	"github.com/ezeev/fastseer/shopify"
	"github.com/gorilla/schema"
)

func (s *Server) handleSearchConfForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		shop := params.Get("shop")
		confIdx := params.Get("confIdx")

		tmpl, err := template.ParseFiles("template/form-search-conf.html", "template/shopify-admin-head-includes.html")
		if err != nil {
			logger.Error(shop, err.Error())
		}

		var shopConf shopify.ShopifyClientConfig
		err = s.ClientsStore.Get(shop, &shopConf)
		if err != nil {
			logger.Error(shop, err.Error())
		}
		idx, err := strconv.Atoi(confIdx)
		if err != nil {
			logger.Error(shop, err.Error())
			fmt.Fprintf(w, err.Error())
		}
		searchConf := shopConf.SearchConfigs[idx]

		type data struct {
			SearchConfig shopify.ShopifySearchConfig
			FlashMessage string
			Config       *Config
			HMacAuth     *HmacAuthParams
			Shop         string
			Index        string
		}

		dataInst := data{
			SearchConfig: *searchConf,
			Config:       s.Config,
			FlashMessage: s.FlashMessage(w, r),
			Shop:         shop,
			HMacAuth:     NewHmacAuthFromParams(params),
			Index:        confIdx,
		}

		err = tmpl.ExecuteTemplate(w, "head-includes", dataInst)
		if err != nil {
			logger.Error(shop, err.Error())
		}
		err = tmpl.ExecuteTemplate(w, "layout", dataInst)
		if err != nil {
			logger.Error(shop, err.Error())
		}

	}
}

func (s *Server) handleCloneSearchConfig() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		shop := params.Get("shop")
		idx, err := strconv.Atoi(params.Get("index"))
		if err != nil {
			logger.Error(shop, err.Error())
			fmt.Fprintf(w, err.Error())
			return
		}
		//load config
		var shopConf shopify.ShopifyClientConfig
		err = s.ClientsStore.Get(shop, &shopConf)
		if err != nil {
			logger.Error(shop, err.Error())
		}

		if len(shopConf.SearchConfigs) >= 3 {
			s.SetFlashError(w, r, "Sorry! You can only have 3 search configurations.")
			//redir := s.Config.AppDomain + routeShopifyHome + "?error=Sorry! You can only have 3 search configurations.&" + NewHmacAuthFromParams(params).QueryString()
			s.RedirectToHome(w, r)
			//http.Redirect(w, r, redir, 307)
		} else {
			confToClone := *shopConf.SearchConfigs[idx]
			confToClone.Name = confToClone.Name + " Clone"

			// Add it back to the conf
			shopConf.SearchConfigs = append(shopConf.SearchConfigs, &confToClone)

			// Add allocation but set to 0
			shopConf.SearchConfigAlloc = append(shopConf.SearchConfigAlloc, 0.0)

			// Save
			err = s.ClientsStore.Put(shop, shopConf)
			if err != nil {
				logger.Error(shop, err.Error())
			}

			// clear cache
			shopConfigCache.Delete(shop)
			s.RedirectToHome(w, r)

		}

	}
}

func (s *Server) handleUpdateSearchConfig() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		shop := params.Get("shop")
		idx, err := strconv.Atoi(params.Get("index"))
		if err != nil {
			logger.Error(shop, err.Error())
			fmt.Fprintf(w, err.Error())
			return
		}

		conf := new(shopify.ShopifySearchConfig)
		decoder := schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)

		err = decoder.Decode(conf, params)
		if err != nil {
			logger.Error(shop, err.Error())
		}

		// fields not currently set by the form:

		//update shop conf
		var shopConf shopify.ShopifyClientConfig
		err = s.ClientsStore.Get(shop, &shopConf)
		if err != nil {
			logger.Error(shop, err.Error())
		}
		if len(shopConf.SearchConfigs) == 0 {
			shopConf.SearchConfigs = make([]*shopify.ShopifySearchConfig, 1)
		} else {
			shopConf.SearchConfigs[idx] = conf
		}

		// save
		err = s.ClientsStore.Put(shop, shopConf)
		if err != nil {
			logger.Error(shop, err.Error())
		}

		// clear cache
		shopConfigCache.Delete(shop)

		s.SetFlashMessage(w, r, "Saved searched configuration "+shopConf.SearchConfigs[idx].Name+".")

		redir := r.Referer()
		http.Redirect(w, r, redir, 307)
		//s.RedirectToHomeWithSuccess(w, r, "Saved searched configuration "+shopConf.SearchConfigs[idx].Name+"."

	}
}

func (s *Server) handleUpdateSearchConfAllocation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		shop := params.Get("shop")
		// decode the form
		tempClientConfig := new(shopify.ShopifyClientConfig)
		decoder := schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		err := decoder.Decode(tempClientConfig, params)
		if err != nil {
			logger.Error(shop, err.Error())
			fmt.Fprintf(w, err.Error())
			return
		}

		// load client conf from db
		var clientConfig shopify.ShopifyClientConfig
		err = s.ClientsStore.Get(shop, &clientConfig)
		if err != nil {
			logger.Error(params.Get("shop"), err.Error())
			fmt.Fprintf(w, err.Error())
			return
		}

		// update allocation
		clientConfig.SearchConfigAlloc = tempClientConfig.SearchConfigAlloc

		// save
		err = s.ClientsStore.Put(shop, clientConfig)
		shopConfigCache.Delete(shop)

		s.SetFlashMessage(w, r, "Updated search allocations.")
		time.Sleep(time.Second * 1)

		s.RedirectToHome(w, r)

	}
}
