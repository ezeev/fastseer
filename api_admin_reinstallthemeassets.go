package fastseer

import (
	"encoding/json"
	"net/http"

	"github.com/ezeev/fastseer/logger"
	"github.com/ezeev/fastseer/shopify"
)

func (s *Server) apiHandleReinstallThemeAssets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		shop := params.Get("shop")
		shopClient, _ := shopify.ShopClientConfig(shop, s.ClientsStore)

		err := shopify.InstallSearchFormThemeAsset(shopClient, s.Config.ScriptTagDomain)
		if err != nil {
			logger.Error(shop, err.Error())
			SendErrorResponse(w, r, "Unable to reinstall assets!", err)
		}

		// re-install script tag
		/*_, err = shopify.InstallShopScriptTag(shopClient, s.Config.ScriptTagDomain)
		//_, err = shopify.InstallShopScriptTag(shopClient, s.Config.ScriptTagDomain)
		if err != nil {
			logger.Error(shop, err.Error())
			SendErrorResponse(w, r, "Unable to reinstall assets!", err)
		}*/
		json.NewEncoder(w).Encode(MessageResponse{Message: "Reinstalled Theme Assets"})
	}
}
