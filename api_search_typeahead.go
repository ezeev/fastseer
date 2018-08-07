package fastseer

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/ezeev/solrg"
	"github.com/prometheus/client_golang/prometheus"
)

// solr query:
// _typeahead/select?defType=edismax&qf=title_txt&q=ro* AND rac*

// expected format by typeahead js:

// TBD

var (
	counterApiTypeAheadRequest = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "fs_typeahead_request_count",
			Help: "The total number of times typeahead has been called",
		}, []string{"shop"})
	counterApiProductSuggestTypeAheadRequest = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "fs_typeahead_prodsuggest_request_count",
			Help: "The total number of times product suggest typeahead has been called",
		}, []string{"shop"})
)

func init() {
	prometheus.MustRegister(counterApiTypeAheadRequest)
	prometheus.MustRegister(counterApiProductSuggestTypeAheadRequest)
}

func (s *Server) handleTypeAheadRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		shop := params.Get("shop")

		counterApiProductSuggestTypeAheadRequest.WithLabelValues(shop).Inc()
		q := params.Get("q")

		if len(q) < 2 {
			return
		}

		shopConfig := s.CachedShopConfig(shop)

		parts := strings.Split(q, " ")

		var query string
		for i, v := range parts {
			query += v + "*"
			if i < len(parts)-1 {
				query += " AND "
			}
		}

		solrParams := &solrg.SolrParams{
			DefType: "edismax",
			Qf:      "title_txt",
			Q:       query,
			Fl:      "title_txt,img_s",
			Rows:    "3",
		}

		//resp, err := s.Search.Query(shop+"_typeahead", "172.104.9.135:8983/solr", solrParams)
		resp, err := s.Search.Query(shop+"_typeahead", shopConfig.IndexAddress, solrParams)
		if err != nil {
			log.Println("Error: " + err.Error())
		}

		solrResp := resp.(*solrg.SolrSearchResponse)

		docs := solrResp.Response.Docs
		//add doc for raw search
		doc := solrg.SolrSearchDocument{}
		doc["title_txt"] = []string{"Search for \"" + q + "\""}
		docs = append(docs, doc)

		json.NewEncoder(w).Encode(docs)
	}
}

func (s *Server) handleTypeAheadTopSearches() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		shop := params.Get("shop")

		counterApiTypeAheadRequest.WithLabelValues(shop).Inc()

		docs := []solrg.SolrSearchDocument{}

		doc1 := solrg.SolrSearchDocument{}
		doc1["search_txt"] = "ipad"
		//doc1["type_s"] = "*"

		doc2 := solrg.SolrSearchDocument{}
		doc2["search_txt"] = "ipad"
		doc2["type_s"] = "Tech Accessories"

		doc3 := solrg.SolrSearchDocument{}
		doc3["search_txt"] = "iphone"
		//doc3["type_s"] = "*"

		doc4 := solrg.SolrSearchDocument{}
		doc4["search_txt"] = "iphone"
		doc4["type_s"] = "Tech Accessories"

		docs = append(docs, doc1, doc2, doc3, doc4)
		json.NewEncoder(w).Encode(docs)

	}
}
