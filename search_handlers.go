package fastseer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ezeev/solrg"
)

// solr query:
// _typeahead/select?defType=edismax&qf=title_txt&q=ro* AND rac*

// expected format by typeahead js:

// TBD

func (s *Server) handleTypeAheadRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// cors
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Content-Type", "application/json")

		params := r.URL.Query()
		shop := params.Get("shop")
		q := params.Get("q")

		if len(q) < 2 {
			return
		}

		// is the shop config cached?
		/*var shopConfig shopify.ShopifyClientConfig

		cachedConfig := shopConfigCache.Get(shop)
		if cachedConfig == nil {
			shopConfigPointer, err := shopify.ShopClientConfig(shop, s.ClientsStore)
			if err != nil {
				logger.Error(shop, err.Error())
			}
			shopConfig = *shopConfigPointer
			//store
			shopConfigCache.Put(shop, shopConfig)
			log.Println("got config from store")
		} else {
			shopConfig = cachedConfig.(shopify.ShopifyClientConfig)
			log.Println("got config from cache")
		}*/

		shopConfig := s.CachedShopConfig(shop)

		parts := strings.Split(q, " ")

		var query string
		for i, v := range parts {
			query += v + "*"
			if i < len(parts)-1 {
				query += " AND "
			}
		}

		log.Println(query)

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

		// cors
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Content-Type", "application/json")

		docs := []solrg.SolrSearchDocument{}

		doc1 := solrg.SolrSearchDocument{}
		doc1["search_txt"] = "ipad"
		doc1["types_ss"] = []string{"Tech Accessories", "Tablets"}

		doc2 := solrg.SolrSearchDocument{}
		doc2["search_txt"] = "iphone"
		doc2["types_ss"] = []string{"Tech Accessories", "Phones"}

		fmt.Println(doc2)

		docs = append(docs, doc1, doc2)
		json.NewEncoder(w).Encode(docs)

	}
}
