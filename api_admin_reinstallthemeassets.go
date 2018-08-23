package fastseer

import (
	"encoding/json"
	"net/http"

	"github.com/ezeev/fastseer/logger"
	"github.com/ezeev/fastseer/shopify"
)

///api/v1/shop/theme/install

/**
 * @api {post} /shop/theme/install Reinstall theme assets into the current shopify theme
 * @apiName PostThemeAssets
 * @apiGroup Theme
 *
 * @apiParam {String} shopify The shopify store i.e. fastseer.myshopify.com
 * @apiParam {String} hmac Secure hmac param from Shopify admin.
 * @apiParam {String} locale Locale from the Shopify admin.
 * @apiParam {Number} timestamp Timestamp from the Shopify admin.
 *
 * @apiSuccess {Object} MessageResponse See <code>MessageResponse</code> type.
 * @apiError {Object} ErrorResponse See <code>ErrorResponse</code> type.
 */
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
