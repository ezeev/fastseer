// @SubApi User [/api/v1/shop/rules]
// @SubApi Allows you access to different features of the users
package fastseer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ezeev/solrg"
	"github.com/google/uuid"

	"github.com/ezeev/fastseer/logger"
	"github.com/ezeev/fastseer/rules"
	"github.com/ezeev/fastseer/shopify"
)

func (s *Server) apiHandlePutRules() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		shop := params.Get("shop")
		shopClient, _ := shopify.ShopClientConfig(shop, s.ClientsStore)

		decoder := json.NewDecoder(r.Body)
		var rules []rules.SearchRule
		err := decoder.Decode(&rules)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error(shop, err.Error())
			SendErrorResponse(w, r, err.Error(), err)
			return
		}

		//does it have an ID?
		for _, v := range rules {
			if v.ID == "" {
				v.ID = uuid.Must(uuid.NewRandom()).String()
			}
		}

		err = s.Search.IndexStruct(shop+"_rules", shopClient.IndexAddress, rules)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error(shop, err.Error())
			SendErrorResponse(w, r, err.Error(), err)
			return
		}

		json.NewEncoder(w).Encode(MessageResponse{Message: "Added/Updated rules"})

	}
}

func (s *Server) apiHandleDeleteRule() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		shop := params.Get("shop")
		ids := params["id"]

		config := s.CachedShopConfig(shop)

		var q string
		for _, k := range ids {
			q += "id:" + k + " "
		}
		log.Println(q)

		err := s.Search.DeleteDocuments(shop+"_rules", config.IndexAddress, q)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error(shop, err.Error())
			SendErrorResponse(w, r, err.Error(), err)
		}

		fmt.Println(ids)

		json.NewEncoder(w).Encode(MessageResponse{Message: "Deleted rules"})

	}
}

// @Title Get Rules Information
// @Description Get Rules
// @Accept json
// @Param q
// @Success 200 {object} string &quot;Success&quot;
// @Failure 401 {object} string &quot;Access denied&quot;
// @Failure 404 {object} string &quot;Not Found&quot;
// @Resource /rules
// @Router /api/v1/shop/rules [get]
func (s *Server) apiHandleGetRules() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		shop := params.Get("shop")
		query := params.Get("q")

		config := s.CachedShopConfig(shop)

		solrParams := &solrg.SolrParams{
			Q: query,
		}

		resp, err := s.Search.Query(shop+"_rules", config.IndexAddress, solrParams)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error(shop, err.Error())
			SendErrorResponse(w, r, err.Error(), err)
		}

		solrResp := resp.(*solrg.SolrSearchResponse)
		data, err := json.Marshal(solrResp.Response.Docs)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error(shop, err.Error())
			SendErrorResponse(w, r, err.Error(), err)
		}

		var rules []rules.SearchRule
		err = json.Unmarshal(data, &rules)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error(shop, err.Error())
			SendErrorResponse(w, r, err.Error(), err)
		}

		json.NewEncoder(w).Encode(rules)

	}
}
