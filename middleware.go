package fastseer

import (
	"fmt"
	"net/http"

	"github.com/ezeev/fastseer/shopify"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	counterTotalApiRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "fs_api_request_count",
			Help: "The total number of times an API call was made",
		}, []string{"shop"})
)

func init() {
	prometheus.MustRegister(counterTotalApiRequests)
}

func (s *Server) handleCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token, content-type, Accept")
		w.Header().Set("Content-Type", "application/json")

		shop := r.URL.Query().Get("shop")
		counterTotalApiRequests.WithLabelValues(shop).Inc()

		if r.Method == "OPTIONS" {
		} else {
			h(w, r)
		}
	}
}

func (s *Server) authorizeShopifyHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()

		// who is the referred?
		/*
			These params must be passed for auth to succeed:
				hmac
				locale
				shop
				timestamp
		*/
		if !shopify.AuthenticateShopifyRequest(params, s.Config.ShopifyApiSecret) {
			w.WriteHeader(500)
			fmt.Fprintf(w, messageInvalidRequest)
			return
		}
		h(w, r)
	}
}

const messageInvalidRequest = "Invalid request. Please contact support."
