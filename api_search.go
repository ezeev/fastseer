package fastseer

import (
	"encoding/json"
	"net/http"

	"github.com/ezeev/fastseer/logger"

	"github.com/ezeev/solrg"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	counterApiSearchRequest = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "fs_search_request_count",
			Help: "The total number of times the search api has been called",
		}, []string{"shop"})
)

func init() {
	prometheus.MustRegister(counterApiSearchRequest)
}

func (s *Server) handleApiSearch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		shop := params.Get("shop")
		q := params.Get("q")
		rows := params.Get("rows")
		start := params.Get("start")

		counterApiSearchRequest.WithLabelValues(shop).Inc()

		shopConfig := s.CachedShopConfig(shop)

		solrParams := &solrg.SolrParams{
			DefType: "edismax",
			Qf:      "variantKeywords_txt_en",
			Q:       q,
			Fl:      "title:productTitle_txt_en,image:productImage_s,id,type:productType_s,productId:productId_s,price:variantPrice_f,score",
			Fq:      []string{"{!collapse field=productId_s}"},
			Rows:    rows,
			Start:   start,
		}

		resp, err := s.Search.Query(shop, shopConfig.IndexAddress, solrParams)
		if err != nil {
			logger.Error(shop, err.Error())
		}
		solrResp := resp.(*solrg.SolrSearchResponse)
		json.NewEncoder(w).Encode(solrResp)

	}
}
