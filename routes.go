package fastseer

import "github.com/gorilla/mux"

const routePing = "/ping"
const routeShopifyCallBack = "/shopify/callback"
const routeShopifyHome = "/shopify"
const routeShopifyBuildIndex = "/shopify/buildIndex"
const routeShopifyClearIndex = "/shopify/clearIndex"
const routeShopifyShopJs = "/shopify/shop.js"
const routeShopifyReInstallSearchForm = "/shopify/reinstallSearchForm"
const routeSearchTypeAhead = "/search/typeahead"
const routeSearchTypeAheadTopSearches = "/search/typeahead/topsearches"

func (s *Server) Routes() {
	s.Router = mux.NewRouter()
	s.Router.HandleFunc(routePing, s.handlePing())
	s.Router.HandleFunc(routeShopifyCallBack, s.handleShopifyCallback())
	s.Router.HandleFunc(routeShopifyHome, s.authorizeShopifyHandler(s.handleShopifyHome()))
	s.Router.HandleFunc(routeShopifyBuildIndex, s.authorizeShopifyHandler(s.handleBuildIndex()))
	s.Router.HandleFunc(routeShopifyClearIndex, s.authorizeShopifyHandler(s.handleClearIndex()))
	s.Router.HandleFunc(routeShopifyReInstallSearchForm, s.authorizeShopifyHandler(s.handleReInstallSearchForm()))

	// front-end handlers
	s.Router.HandleFunc(routeShopifyShopJs, s.handleShopifyJs())
	s.Router.HandleFunc(routeSearchTypeAhead, s.handleTypeAheadRequest())
	s.Router.HandleFunc(routeSearchTypeAheadTopSearches, s.handleTypeAheadTopSearches())

}
