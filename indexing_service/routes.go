package main

import "github.com/gorilla/mux"

const routePing = "/ping"
const routeIndexShopify = "/api/indexShopify"
const routeIndexStats = "/api/indexStats"

func (s *Server) Routes() {
	s.Router = mux.NewRouter()
	s.Router.HandleFunc(routePing, s.handlePing())
	s.Router.HandleFunc(routeIndexShopify, s.restEndpoint(s.handleIndexShopifyCatalog())).Methods("POST")
	s.Router.HandleFunc(routeIndexShopify, s.restEndpoint(s.handleClearShopifyCatalog())).Methods("DELETE")
	s.Router.HandleFunc(routeIndexStats, s.restEndpoint(s.handleAdminIndexStats())).Methods("GET")
}
