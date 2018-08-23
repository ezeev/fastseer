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

/**
 * @api {put} /shop/rules Put one to many search rules
 * @apiName PutRules
 * @apiGroup Rules
 * @apiParam {Object} rules The request body, accepts a json <code>[]rules.SearchRule</code>. Leave id blank if you are creating a new rule. Otherwise, and existing rule will be overwritten if the ids match.
 * @apiParam {String} shopify The shopify store i.e. fastseer.myshopify.com
 * @apiParam {String} hmac Secure hmac param from Shopify admin.
 * @apiParam {String} locale Locale from the Shopify admin.
 * @apiParam {Number} timestamp Timestamp from the Shopify admin.
 * @apiParamExample {json} Request Body Example:
		[
      {
        "id":"b1f3cd69-9d44-404e-a97b-993568391988",
        "name_s":"ipad sale",
        "actAddBqs_ss":["id:1234^10"],
        "order_i":1},
      {
        "id":"dd75c626-c62b-4595-bae5-9692f9449408",
        "name_s":"iphone promotion",
        "order_i":1}]
 * @apiSuccess {Object} MessageResponse See <code>MessageResponse</code> type.
 * @apiError {Object} ErrorResponse See <code>ErrorResponse</code> type.
*/
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

		//does it have an ID? if not, create one
		for i, v := range rules {
			if v.ID == "" {
				rules[i].ID = uuid.Must(uuid.NewRandom()).String()
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

/**
 * @api {delete} /shop/rules Delete one to many search rules by id
 * @apiName DeleteRules
 * @apiGroup Rules
 * @apiParam {String} id The id of the rule you want to delete. You can pass as many id params as desired.
 * @apiParam {String} shopify The shopify store i.e. fastseer.myshopify.com
 * @apiParam {String} hmac Secure hmac param from Shopify admin.
 * @apiParam {String} locale Locale from the Shopify admin.
 * @apiParam {Number} timestamp Timestamp from the Shopify admin.
 * @apiSuccess {Object} MessageResponse See <code>MessageResponse</code> type.
 * @apiError {Object} ErrorResponse See <code>ErrorResponse</code> type.
 */
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

/**
 * @api {get} /shop/rules Search and retrieve rules
 * @apiName GetRules
 * @apiGroup Rules
 * @apiParam {String} q The query to search for rules
 * @apiParam {String} shopify The shopify store i.e. fastseer.myshopify.com
 * @apiParam {String} hmac Secure hmac param from Shopify admin.
 * @apiParam {String} locale Locale from the Shopify admin.
 * @apiParam {Number} timestamp Timestamp from the Shopify admin.
 * @apiSuccess {Object} SearchRules Returns a json array of <code>SearchRule</code>.
 * @apiError {Object} ErrorResponse See <code>ErrorResponse</code> type.
 */
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
