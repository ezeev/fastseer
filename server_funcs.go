package fastseer

import (
	"net/http"
	"time"
)

const flashMessageCookieID = "msg"

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
