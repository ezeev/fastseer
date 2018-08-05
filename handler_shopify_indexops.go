package fastseer

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ezeev/fastseer/logger"
	"github.com/ezeev/fastseer/shopify"
)

func (s *Server) handleBuildIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		shop := params.Get("shop")
		shopClient, _ := shopify.ShopClientConfig(shop, s.ClientsStore)

		// send job to indexing server
		indexWorker := s.Config.IndexingWorkerServices[0]
		cli := &http.Client{
			Timeout: time.Second * 10,
		}

		b, err := json.Marshal(shopClient)
		if err != nil {
			logger.Error(shop, err.Error())
		}
		req, err := http.NewRequest("POST", indexWorker+s.Config.IndexingWorkerEndpoint, bytes.NewBuffer(b))
		if err != nil {
			logger.Error(shop, err.Error())
		}
		resp, err := cli.Do(req)
		if err != nil {
			logger.Error(shop, err.Error())
		}
		defer resp.Body.Close()

		s.SetFlashMessage(w, r, "Builing Index...")

		//redir := r.Referer()
		//http.Redirect(w, r, redir, 307)
		s.RedirectToHome(w, r)
	}
}

func (s *Server) handleClearIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		shop := params.Get("shop")
		shopClient, _ := shopify.ShopClientConfig(shop, s.ClientsStore)

		// send job to indexing server
		indexWorker := s.Config.IndexingWorkerServices[0]
		cli := &http.Client{
			Timeout: time.Second * 10,
		}

		b, err := json.Marshal(shopClient)
		if err != nil {
			logger.Error(shop, err.Error())
		}
		req, err := http.NewRequest("DELETE", indexWorker+s.Config.IndexingWorkerEndpoint, bytes.NewBuffer(b))
		if err != nil {
			logger.Error(shop, err.Error())
		}
		resp, err := cli.Do(req)
		if err != nil {
			logger.Error(shop, err.Error())
		}
		defer resp.Body.Close()

		s.SetFlashMessage(w, r, "Clearing Index...")
		//redir := r.Referer()
		//http.Redirect(w, r, redir, 307)
		s.RedirectToHome(w, r)

	}
}
