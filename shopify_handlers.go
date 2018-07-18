package fastseer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ezeev/fastseer/shopify"
)

// GetPermanentAccessToken returns the shop name and permanent access token respectively
func GetPermanentAccessToken(shop string, apiKey, apiSecret string, code string) shopify.ShopifyAuthResponse {

	tokenUrl := fmt.Sprintf("https://%s/admin/oauth/access_token", shop)
	v := url.Values{}
	v.Set("client_id", apiKey)
	v.Set("client_secret", apiSecret)
	v.Set("code", code)
	s := v.Encode()
	req, err := http.NewRequest("POST", tokenUrl, strings.NewReader(s))
	if err != nil {
		log.Printf("http.NewRequest() error: %v\n", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	c := &http.Client{}

	log.Printf("requesting token at %s", tokenUrl)

	resp, err := c.Do(req)
	if err != nil {
		log.Printf("http.Do() error: %v\n", err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll() error: %v\n", err)
	}

	var tokenResp shopify.ShopifyAuthResponse
	err = json.Unmarshal(data, &tokenResp)
	if err != nil {
		log.Fatal("Error in json.Unmarshal() for Shopify Token Response")
		log.Fatal(err)
	}

	return tokenResp
}

type ShopifyPageData struct {
	Config    *Config
	Shop      string
	HMac      string
	Timestamp string
	Locale    string
	Protocol  string
}

const messageInvalidRequest = "Invalid request. Please contact support."

func (s *Server) handleShopifyHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tmpl, err := template.ParseFiles("template/home.html")
		if err != nil {
			panic(err)
		}

		params := r.URL.Query()

		if !shopify.AuthenticateShopifyRequest(params, s.Config.ShopifyApiSecret) {
			fmt.Fprintf(w, messageInvalidRequest)
			return
		}

		data := ShopifyPageData{
			Config:    s.Config,
			Shop:      params.Get("shop"),
			HMac:      params.Get("hmac"),
			Timestamp: params.Get("timestamp"),
			Locale:    params.Get("locale"),
			Protocol:  params.Get("protocol"),
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			panic(err)
		}
	}
}

func (s *Server) handleShopifyCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()

		code := params.Get("code")
		hmac := params.Get("hmac")
		//timestamp := params.Get("timestamp")
		//state := params.Get("state")
		shop := params.Get("shop")
		//log.Printf("Install request: code: %s, hmac: %s, timestamp: %s, state: %s, shop: %s", code, hmac, timestamp, state, shop)
		params.Del("hmac")
		params.Del("signature")
		message := params.Encode()

		apiKey := s.Config.ShopifyApiKey
		apiSecret := s.Config.ShopifyApiSecret

		if shopify.VerifyRequest(hmac, message, apiSecret) {
			log.Printf("Received Valid HMAC Request for New Installation (%s)", shop)
		} else {
			log.Println("Invalid request to install app")
		}

		// now get client api key
		tokenResp := GetPermanentAccessToken(shop, apiKey, apiSecret, code)

		log.Printf("Received access token: %s\n", tokenResp.AccessToken)
		client := shopify.ShopifyClientConfig{
			Shop:         shop,
			IndexAddress: s.Config.DefaultIndexAddress,
			AuthResponse: tokenResp,
		}
		err := s.ClientsStore.Put(shop, client)
		if err != nil {
			log.Printf("Error: %s", err.Error())
		}

		// create the search collections for this customer
		// In the future, we may allocate different customers to different Solr clusters
		// once a cluster reaches full capacity
		go shopify.CreateClientCollections(s.Search, client.IndexAddress, shop)

		// redirect back to shopify admin
		redir := fmt.Sprintf("https://%s/admin/apps/%s", shop, apiKey)
		http.Redirect(w, r, redir, 301)

	}
}

func (s *Server) handleBuildIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		if !shopify.AuthenticateShopifyRequest(params, s.Config.ShopifyApiSecret) {
			fmt.Fprintf(w, messageInvalidRequest)
			return
		} else {
			shop := params.Get("shop")
			var shopClient shopify.ShopifyClientConfig
			err := s.ClientsStore.Get(shop, &shopClient)
			if err != nil {
				log.Printf("Error in handleBuildIndex() for shop: %s: %s", shop, err.Error())
				log.Println(shopClient)
			}

			// send job to indexing server
			indexWorker := s.Config.IndexingWorkerServices[0]
			cli := &http.Client{
				Timeout: time.Second * 10,
			}

			b, err := json.Marshal(shopClient)
			if err != nil {
				log.Printf("Error marshalling ShopifyClientConfig: %s", err)
			}
			req, err := http.NewRequest("POST", indexWorker+s.Config.IndexingWorkerEndpoint, bytes.NewBuffer(b))
			if err != nil {
				log.Printf("Error building http request: %s", err)
			}
			resp, err := cli.Do(req)
			if err != nil {
				log.Println("Error sending request: %s", err)
			}
			defer resp.Body.Close()

			log.Println("made it here test")
			//shopify.CrawlProducts(&shopClient, 3, "", s.Search)

		}
		redir := r.Referer()
		http.Redirect(w, r, redir, 307)
	}
}
