package fastseer

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ezeev/fastseer/logger"
	"github.com/ezeev/fastseer/shopify"
)

func (s *Server) apiHandleBuildIndex() http.HandlerFunc {
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
			SendErrorResponse(w, r, "Unable to start index build", err)
		}
		req, err := http.NewRequest("POST", indexWorker+s.Config.IndexingWorkerEndpoint, bytes.NewBuffer(b))
		if err != nil {
			logger.Error(shop, err.Error())
			SendErrorResponse(w, r, "Unable to start index build", err)
		}
		resp, err := cli.Do(req)
		if err != nil {
			logger.Error(shop, err.Error())
			SendErrorResponse(w, r, "Unable to start index build", err)
		}
		defer resp.Body.Close()
		json.NewEncoder(w).Encode(MessageResponse{Message: "Started indexing."})

	}
}
