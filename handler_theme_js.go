package fastseer

import (
	"html/template"
	"log"
	"net/http"

	"github.com/ezeev/fastseer/shopify"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	counterClientJs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "fs_client_js_request_count",
			Help: "The total number of times a client requested shop js",
		}, []string{"shop"})
)

func init() {
	prometheus.MustRegister(counterClientJs)
}

func (s *Server) handleShopifyJs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		shop := params.Get("shop")

		counterClientJs.WithLabelValues(shop).Inc()

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
