package fastseer

import (
	"html/template"
	"log"
	"net/http"

	"github.com/ezeev/fastseer/shopify"
)

func (s *Server) handleShopifyJs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		shop := params.Get("shop")

		conf := s.CachedShopConfig(shop)

		type JsData struct {
			Shop      string
			AppDomain string
			Conf      shopify.ShopifySearchConfig
		}

		data := JsData{shop, s.Config.AppDomain, *conf.SearchConfigs[0]}

		log.Println(data.Conf.IncludeProductSuggesitons)

		//tmpl, err := template.ParseFiles("template/shopify.js")
		pageFile := "template/shopify.js"
		tmpl, err := template.New("shopify.js").Delims("[[", "]]").ParseFiles(pageFile)
		if err != nil {
			log.Panic(err)
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			log.Panic(err)
		}

		//fmt.Fprintf(w, js)
	}
}
