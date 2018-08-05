package fastseer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ezeev/fastseer/logger"
	"github.com/ezeev/fastseer/shopify"
)

func (s *Server) apiHandleCloneSearchConfig() http.HandlerFunc {
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
			SendErrorResponse(w, r, "Unable to clone config", err)
		}

		if len(shopConf.SearchConfigs) >= 3 {
			SendErrorResponse(w, r, "You can only have 3 search configurations!", fmt.Errorf("You can only have 3 search configurations!"))
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
				SendErrorResponse(w, r, "Unable to clone config", err)
			}

			// clear cache
			shopConfigCache.Delete(shop)
			json.NewEncoder(w).Encode(MessageResponse{Message: "Cloned \"" + shopConf.SearchConfigs[idx].Name + "\""})

		}

	}
}
