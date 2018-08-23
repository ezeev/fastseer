package fastseer

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const apiV1HandleGetNumProducts = "/api/v1/products/count"
const apiV1HandleBuildIndex = "/api/v1/products/index"
const apiV1HandleClearIndex = "/api/v1/products/index"
const apiV1HanldGetShopConf = "/api/v1/shop/config"
const apiV1HandlePostShopConf = "/api/v1/shop/config"
const apiV1HandleCloneSearchConf = "/api/v1/shop/search/config/clone"
const apiV1HandleReinstallThemeAssets = "/api/v1/shop/theme/install"
const apiV1HandlePutRules = "/api/v1/shop/rules"
const apiV1HandleDeleteRules = "/api/v1/shop/rules"
const apiV1HandleGetRules = "/api/v1/shop/rules"

//public api
const apiSearchTypeAhead = "/search/typeahead"
const apiSearchTypeAheadTopSearches = "/search/typeahead/topsearches"

func (s *Server) ApiRoutes() {

	//metrics
	s.Router.Handle("/metrics", promhttp.Handler())

	// admin api
	s.Router.HandleFunc(apiV1HandleGetNumProducts, s.handleCORS(s.authorizeShopifyHandler(s.apiHandleGetNumProducts())))
	s.Router.HandleFunc(apiV1HandleBuildIndex, s.handleCORS(s.authorizeShopifyHandler(s.apiHandleBuildIndex()))).Methods("POST")
	s.Router.HandleFunc(apiV1HandleBuildIndex, s.handleCORS(s.authorizeShopifyHandler(s.apiHandleClearIndex()))).Methods("DELETE", "OPTIONS")
	s.Router.HandleFunc(apiV1HanldGetShopConf, s.handleCORS(s.authorizeShopifyHandler(s.apiHandleGetShopConfig()))).Methods("GET")
	s.Router.HandleFunc(apiV1HandlePostShopConf, s.handleCORS(s.authorizeShopifyHandler(s.apiHandlePostShopConfig()))).Methods("POST", "OPTIONS")
	s.Router.HandleFunc(apiV1HandleCloneSearchConf, s.handleCORS(s.authorizeShopifyHandler(s.apiHandleCloneSearchConfig()))).Methods("POST")
	s.Router.HandleFunc(apiV1HandleReinstallThemeAssets, s.handleCORS(s.authorizeShopifyHandler(s.apiHandleReinstallThemeAssets()))).Methods("POST")
	s.Router.HandleFunc(apiV1HandlePutRules, s.handleCORS(s.authorizeShopifyHandler(s.apiHandlePutRules()))).Methods("PUT", "OPTIONS")
	s.Router.HandleFunc(apiV1HandleDeleteRules, s.handleCORS(s.authorizeShopifyHandler(s.apiHandleDeleteRule()))).Methods("DELETE")
	s.Router.HandleFunc(apiV1HandleGetRules, s.handleCORS(s.authorizeShopifyHandler(s.apiHandleGetRules()))).Methods("GET")
	//apiHandleGetRules()

	// public api
	s.Router.HandleFunc(apiSearchTypeAhead, s.handleCORS(s.handleTypeAheadRequest()))
	s.Router.HandleFunc(apiSearchTypeAheadTopSearches, s.handleCORS(s.handleTypeAheadTopSearches()))

}
