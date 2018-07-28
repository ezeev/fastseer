package fastseer

import (
	"net/http"
	"time"

	"github.com/ezeev/fastseer/logger"
	"github.com/ezeev/fastseer/memkv"
	"github.com/ezeev/fastseer/shopify"
)

const flashMessageCookieID = "msg"

var shopConfigCache *memkv.MemKV

func init() {

	shopConfigCache = memkv.New(100, 100)

}

func (s *Server) SetFlashMessage(w http.ResponseWriter, msg string) {
	exp := time.Now().Add(time.Second * 3)
	cookie := http.Cookie{Name: flashMessageCookieID, Value: msg, Expires: exp}
	http.SetCookie(w, &cookie)
}

func (s *Server) FlashMessage(r *http.Request) string {
	cookie, _ := r.Cookie(flashMessageCookieID)
	if cookie != nil {
		return cookie.Value
	}
	return ""
}

func (s *Server) CachedShopConfig(shop string) *shopify.ShopifyClientConfig {
	var shopConfig shopify.ShopifyClientConfig
	cachedConfig := shopConfigCache.Get(shop)
	if cachedConfig == nil {
		shopConfigPointer, err := shopify.ShopClientConfig(shop, s.ClientsStore)
		if err != nil {
			logger.Error(shop, err.Error())
		}
		shopConfig = *shopConfigPointer
		//store
		shopConfigCache.Put(shop, shopConfig)
	} else {
		shopConfig = cachedConfig.(shopify.ShopifyClientConfig)
	}
	return &shopConfig
}
