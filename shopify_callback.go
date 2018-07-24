package fastseer

import (
	"fmt"
	"net/http"

	"github.com/ezeev/fastseer/logger"
	"github.com/ezeev/fastseer/shopify"
)

func (s *Server) handleShopifyCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()

		code := params.Get("code")
		hmac := params.Get("hmac")
		//timestamp := params.Get("timestamp")
		//state := params.Get("state")
		shop := params.Get("shop")
		params.Del("hmac")
		params.Del("signature")
		message := params.Encode()

		apiKey := s.Config.ShopifyApiKey
		apiSecret := s.Config.ShopifyApiSecret

		if shopify.VerifyRequest(hmac, message, apiSecret) {
			logger.Info(shop, "recieved valid HMAC response")
		} else {
			logger.Error(shop, "received invalid HMAC response")
		}

		// now get client api key
		tokenResp := shopify.GetPermanentAccessToken(shop, apiKey, apiSecret, code)

		logger.Info(shop, "Received access token: "+tokenResp.AccessToken)

		client := shopify.ShopifyClientConfig{
			Shop:         shop,
			IndexAddress: s.Config.DefaultIndexAddress,
			AuthResponse: tokenResp,
		}
		err := s.ClientsStore.Put(shop, client)
		if err != nil {
			logger.Error(shop, err.Error())
		}

		// create the search collections for this customer
		// In the future, we may allocate different customers to different Solr clusters
		// once a cluster reaches full capacity
		go shopify.CreateClientCollections(s.Search, client.IndexAddress, shop)

		// redirect back to shopify admin
		redir := fmt.Sprintf("https://%s/admin/apps/%s", shop, apiKey)
		http.Redirect(w, r, redir, 301)

	}
}
