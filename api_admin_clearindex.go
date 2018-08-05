package fastseer

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ezeev/fastseer/logger"
	"github.com/ezeev/fastseer/shopify"
)

func (s *Server) apiHandleClearIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		shop := params.Get("shop")
		shopClient, _ := shopify.ShopClientConfig(shop, s.ClientsStore)

		errMsg := "Unable to clear index"
		// send job to indexing server
		indexWorker := s.Config.IndexingWorkerServices[0]
		cli := &http.Client{
			Timeout: time.Second * 10,
		}

		b, err := json.Marshal(shopClient)
		if err != nil {
			logger.Error(shop, err.Error())
			SendErrorResponse(w, r, errMsg, err)
		}
		req, err := http.NewRequest("DELETE", indexWorker+s.Config.IndexingWorkerEndpoint, bytes.NewBuffer(b))
		if err != nil {
			logger.Error(shop, err.Error())
			SendErrorResponse(w, r, errMsg, err)
		}
		resp, err := cli.Do(req)
		if err != nil {
			logger.Error(shop, err.Error())
			SendErrorResponse(w, r, errMsg, err)
		}
		defer resp.Body.Close()
		json.NewEncoder(w).Encode(MessageResponse{Message: "Clearing index."})

	}
}
