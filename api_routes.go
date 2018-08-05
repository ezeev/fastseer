package fastseer

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const apiV1HandleGetNumProducts = "/api/v1/products/count"
const apiV1HandleBuildIndex = "/api/v1/products/index"
const apiV1HandleClearIndex = "/api/v1/products/index"
const apiV1HanldGetShopConf = "/api/v1/shop/config"
const apiV1handlePostShopConf = "/api/v1/shop/config"
const apiV1handleCloneSearchConf = "/api/v1/shop/search/config/clone"
const apiV1handleReinstallThemeAssets = "/api/v1/shop/theme/install"

// apiHandleReinstallThemeAssets()

func (s *Server) ApiRoutes() {

	//metrics
	s.Router.Handle("/metrics", promhttp.Handler())

	//api routes
	s.Router.HandleFunc(apiV1HandleGetNumProducts, s.handleCORS(s.authorizeShopifyHandler(s.apiHandleGetNumProducts())))
	s.Router.HandleFunc(apiV1HandleBuildIndex, s.handleCORS(s.authorizeShopifyHandler(s.apiHandleBuildIndex()))).Methods("POST")
	s.Router.HandleFunc(apiV1HandleBuildIndex, s.handleCORS(s.authorizeShopifyHandler(s.apiHandleClearIndex()))).Methods("DELETE", "OPTIONS")
	s.Router.HandleFunc(apiV1HanldGetShopConf, s.handleCORS(s.authorizeShopifyHandler(s.apiHandleGetShopConfig()))).Methods("GET")
	s.Router.HandleFunc(apiV1handlePostShopConf, s.handleCORS(s.authorizeShopifyHandler(s.apiHandlePostShopConfig()))).Methods("POST", "OPTIONS")
	s.Router.HandleFunc(apiV1handleCloneSearchConf, s.handleCORS(s.authorizeShopifyHandler(s.apiHandleCloneSearchConfig()))).Methods("POST")
	s.Router.HandleFunc(apiV1handleReinstallThemeAssets, s.handleCORS(s.authorizeShopifyHandler(s.apiHandleReinstallThemeAssets()))).Methods("POST")
}
