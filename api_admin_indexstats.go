package fastseer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ezeev/fastseer/logger"

	"github.com/ezeev/fastseer/shopify"
)

func (s *Server) apiHandleGetNumProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		shop := params.Get("shop")
		shopConfig := s.CachedShopConfig(shop)
		log.Println(shop)
		products, err := shopify.NumIndexedProducts(shopConfig, s.Search)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			logger.Error(shop, err.Error())
		}

		variants, err := shopify.NumIndexedVariants(shopConfig, s.Search)
		if err != nil {
			r.Response.StatusCode = 500
			fmt.Fprintf(w, err.Error())
			logger.Error(shop, err.Error())
		}

		type resp struct {
			Products int `json:"products"`
			Variants int `json:"variants"`
		}

		json.NewEncoder(w).Encode(resp{
			Products: products,
			Variants: variants,
		})

	}
}
