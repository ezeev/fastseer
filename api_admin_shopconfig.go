package fastseer

import (
	"encoding/json"
	"net/http"

	"github.com/ezeev/fastseer/logger"
	"github.com/ezeev/fastseer/shopify"
)

func (s *Server) apiHandleGetShopConfig() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		shop := params.Get("shop")
		conf, err := shopify.ShopClientConfig(shop, s.ClientsStore)
		if err != nil {
			logger.Error(shop, err.Error())
		}
		json.NewEncoder(w).Encode(conf)

	}
}

func (s *Server) apiHandlePostShopConfig() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)
		var shop shopify.ShopifyClientConfig
		err := decoder.Decode(&shop)
		if err != nil {
			logger.Error(shop.Shop, err.Error())
			SendErrorResponse(w, r, "Error posting config", err)
		}
		//save
		err = s.ClientsStore.Put(shop.Shop, &shop)
		if err != nil {
			logger.Error(shop.Shop, err.Error())
			SendErrorResponse(w, r, "Error posting config", err)
		}

		//clear cache
		shopConfigCache.Delete(shop.Shop)

		json.NewEncoder(w).Encode(MessageResponse{Message: "Updated Shop Configuration."})
	}
}
