package fastseer

import (
	"net/http"

	"github.com/ezeev/fastseer/logger"
	"github.com/ezeev/fastseer/memkv"
	"github.com/ezeev/fastseer/shopify"
)

const sessionName = "fastseer-admin-session"
const flashMessageCookieID = "msg"
const flashErrorCookieID = "errorMsg"

var shopConfigCache *memkv.MemKV

func init() {
	shopConfigCache = memkv.New(100, 60)
}

func (s *Server) AppDomain(r *http.Request) string {
	return "https://" + r.Host
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

//var store = sessions.NewCookieStore([]byte("something-very-secret"))

/*func (s *Server) SetFlashMessage(w http.ResponseWriter, r *http.Request, msg string) {
/*exp := time.Now().Add(time.Second * 5)
cookie := http.Cookie{Name: flashMessageCookieID, Value: msg, Expires: exp}
http.SetCookie(w, &cookie)*/

/*
	session, _ := store.Get(r, sessionName)
	session.Values[flashMessageCookieID] = msg
	session.Save(r, w)
}*/

/*func (s *Server) FlashMessage(w http.ResponseWriter, r *http.Request) string {
	session, _ := store.Get(r, sessionName)
	msg := ""
	if session.Values[flashMessageCookieID] != nil {
		msg = session.Values[flashMessageCookieID].(string)
		// delete it
		session.Values[flashMessageCookieID] = nil
		session.Save(r, w)
	}
	return msg
}*/

/*func (s *Server) SetFlashError(w http.ResponseWriter, r *http.Request, msg string) {
	session, _ := store.Get(r, sessionName)
	session.Values[flashErrorCookieID] = msg
	session.Save(r, w)
}*/

/*func (s *Server) FlashError(w http.ResponseWriter, r *http.Request) string {
	session, _ := store.Get(r, sessionName)
	msg := ""
	if session.Values[flashErrorCookieID] != nil {
		msg = session.Values[flashErrorCookieID].(string)
		// delete it
		session.Values[flashErrorCookieID] = nil
		session.Save(r, w)
	}
	return msg
}*/

/*func (s *Server) RedirectToHome(w http.ResponseWriter, r *http.Request) {
	redir := s.Config.AppDomain + routeShopifyHome + "?" + NewHmacAuthFromParams(r.URL.Query()).QueryString()
	http.Redirect(w, r, redir, 307)
}*/
