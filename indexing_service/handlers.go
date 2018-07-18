package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/ezeev/fastseer/shopify"
)

var runningJobs sync.Map

func (s *Server) restEndpoint(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		h(w, r)
	}
}

func (s *Server) handleIndexShopifyCatalog() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// get the user config
		decoder := json.NewDecoder(r.Body)
		var shop shopify.ShopifyClientConfig
		err := decoder.Decode(&shop)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error reading shopify client config: %s", err)
			fmt.Fprintf(w, "Error reading shopify client config")
		}

		pageSize := 3
		count, err := shopify.GetNumProducts(shop.AuthResponse.AccessToken, shop.Shop)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error getting product count: %s", err)
		}
		//required params
		go func() {
			// start the job
			runningJobs.Store(shop.Shop, time.Now().Unix())
			err := shopify.CrawlProducts(&shop, pageSize, "", s.Search)
			if err != nil {
				log.Println(err)
			}
			time.Sleep(time.Second)
			// remove job
			runningJobs.Delete(shop.Shop)
		}()

		resp := IndexingStatusResponse{
			Shop:            shop.Shop,
			Message:         "Indexing job started.",
			ProductsToCrawl: count.Count,
			PageSize:        pageSize,
		}

		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Printf("Error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

	}
}

func (s *Server) handleClearShopifyCatalog() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// get the user config
		decoder := json.NewDecoder(r.Body)
		var shop shopify.ShopifyClientConfig
		err := decoder.Decode(&shop)
		if err != nil {
			log.Printf("Error reading shopify client config: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error reading shopify client config")
		}

		err = s.Search.DeleteDocuments(shop.Shop, shop.IndexAddress, "*:*")
		if err != nil {
			log.Printf("Error clearing index: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error clearing index: "+err.Error())
		}

		fmt.Fprintf(w, "Clearing index")

	}
}

func (s *Server) handleAdminIndexStats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jobs := make(map[string]interface{})
		runningJobs.Range(func(k, v interface{}) bool {
			jobs[k.(string)] = v
			return true
		})
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(jobs)
	}
}
