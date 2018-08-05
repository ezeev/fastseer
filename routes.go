package fastseer

import (
	"net/http"

	"github.com/gorilla/mux"
)

const routePing = "/ping"
const routeShopifyCallBack = "/shopify/callback"
const routeShopifyHome = "/shopify"
const routeShopifyBuildIndex = "/shopify/buildIndex"
const routeShopifyClearIndex = "/shopify/clearIndex"
const routeShopifyShopJs = "/shopify/shop.js"
const routeShopifyReInstallSearchForm = "/shopify/reinstallSearchForm"
const routeShopifyUpdateSearchConfig = "/shopify/updateSearchConfig"

const routeShopifySearchConfForm = "/shopify/forms/searchConf"
const routeShopifySearchConfAllocation = "/shopify/forms/updateSearchAllocation"

//handleCloneSearchConfig()
const routeShopifyCloneSearchConf = "/shopify/forms/cloneSearchConf"

const routeSearchTypeAhead = "/search/typeahead"
const routeSearchTypeAheadTopSearches = "/search/typeahead/topsearches"

func (s *Server) Routes() {

	// SECURITY RULES!!!
	// If a handler renders a page with the Shopify JS in the header (i.e. the admin home, /shopify),
	// then the shopify JS will handle verification for us automatically
	// IF the handler IS A REDIRECT, i.e. it is an endpoint that forms are posting do, then you
	// need to pass it through server.authortizeShopifyHandler()!!!! - this will perform HMAC auth using
	// the params passed to the handler. Use the NewHmacAuthFromParams struct to simplify. See /shopify for
	// examples

	s.Router = mux.NewRouter()
	s.Router.HandleFunc(routePing, s.handlePing())
	s.Router.HandleFunc(routeShopifyCallBack, s.handleShopifyCallback())
	s.Router.HandleFunc(routeShopifyHome, s.handleShopifyHome())
	s.Router.HandleFunc(routeShopifyBuildIndex, s.authorizeShopifyHandler(s.handleBuildIndex()))
	s.Router.HandleFunc(routeShopifyClearIndex, s.authorizeShopifyHandler(s.handleClearIndex()))
	s.Router.HandleFunc(routeShopifyReInstallSearchForm, s.authorizeShopifyHandler(s.handleReInstallSearchForm()))
	s.Router.HandleFunc(routeShopifyUpdateSearchConfig, s.authorizeShopifyHandler(s.handleUpdateSearchConfig()))
	s.Router.HandleFunc(routeShopifySearchConfForm, s.handleSearchConfForm())
	s.Router.HandleFunc(routeShopifySearchConfAllocation, s.authorizeShopifyHandler(s.handleUpdateSearchConfAllocation()))
	s.Router.HandleFunc(routeShopifyCloneSearchConf, s.handleCloneSearchConfig())

	// front-end handlers
	s.Router.HandleFunc(routeShopifyShopJs, s.handleShopifyJs())
	s.Router.HandleFunc(routeSearchTypeAhead, s.handleTypeAheadRequest())
	s.Router.HandleFunc(routeSearchTypeAheadTopSearches, s.handleTypeAheadTopSearches())

	s.ApiRoutes()

	// static web handler
	s.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("shopify-admin-ui/build/")))

}
