package fastseer

import (
	"fmt"
	"html/template"
	"net/url"
)

type HmacAuthParams struct {
	Shop      string
	HMac      string
	Timestamp string
	Locale    string
	Protocol  string
}

func NewHmacAuthFromParams(params url.Values) *HmacAuthParams {
	auth := &HmacAuthParams{
		Shop:      params.Get("shop"),
		HMac:      params.Get("hmac"),
		Timestamp: params.Get("timestamp"),
		Locale:    params.Get("locale"),
		Protocol:  params.Get("protocol"),
	}
	return auth
}

func (h *HmacAuthParams) QueryString() string {
	//shop={{.Shop}}&hmac={{.HMac}}&timestamp={{.Timestamp}}&locale={{.Locale}}
	return fmt.Sprintf("shop=%s&hmac=%s&timestamp=%s&locale=%s", h.Shop, h.HMac, h.Timestamp, h.Locale)
}

func (h *HmacAuthParams) HiddenFormFields() template.HTML {

	html := fmt.Sprintf(`
		<input type="hidden" name="shop" value="%s" />
		<input type="hidden" name="hmac" value="%s" />
		<input type="hidden" name="timestamp" value="%s" />
		<input type="hidden" name="locale" value="%s" />
	`, h.Shop, h.HMac, h.Timestamp, h.Locale)

	return template.HTML(html)

}
