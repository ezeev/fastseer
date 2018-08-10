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

var jsTemplate *template.Template

func init() {
	prometheus.MustRegister(counterClientJs)
	pageFile := "template/shopify.js"
	var err error
	jsTemplate, err = template.New("shopify.js").Delims("<<", ">>").ParseFiles(pageFile)
	if err != nil {
		log.Fatal("Couldn't load JS template!: Error: ", err.Error())
	}

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

		data := JsData{shop, s.Config.ScriptTagDomain, *conf.SearchConfigs[0]}
		//data := JsData{shop, s.Config.ScriptTagDomain, *conf.SearchConfigs[0]}

		//tmpl, err := template.ParseFiles("template/shopify.js")
		/*pageFile := "template/shopify.js"
		tmpl, err := template.New("shopify.js").Delims("[[", "]]").ParseFiles(pageFile)
		if err != nil {
			log.Panic(err)
		}*/
		err := jsTemplate.Execute(w, data)
		if err != nil {
			log.Panic(err)
		}

		//fmt.Fprintf(w, js)
	}
}
