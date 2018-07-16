package fastseer

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ezeev/fastseer/search"
	"github.com/ezeev/fastseer/storage"
	"github.com/gorilla/mux"
)

type Server struct {
	Router       *mux.Router
	HttpServer   *http.Server
	Port         int
	Config       *Config
	ClientsStore storage.Storage
	Search       search.SearchEngine
}

// NewServer returns an initialized instance of a new server
func NewServer(confPath string) (*Server, error) {

	conf := LoadConfigFromFile(confPath)

	s := &Server{}
	s.Config = conf
	s.Port = conf.ServerPort

	// clients storage
	clientsStore, err := storage.NewStorage(s.Config.DbImpl)
	if err != nil {
		return nil, err
	}
	opts := s.Config.DbOptions
	opts["table"] = "clients"
	opts["keyField"] = "store"
	err = clientsStore.Open(opts)
	if err != nil {
		return nil, err
	}
	s.ClientsStore = clientsStore

	// search engine
	searchEngine, err := search.NewSearchEngine(s.Config.SearchImpl)
	if err != nil {
		return nil, err
	}

	s.Search = searchEngine
	s.Router = mux.NewRouter()

	// setup routes
	s.Router.HandleFunc("/ping", s.handlePing())

	// shopify routes
	s.Router.HandleFunc("/shopify/callback", s.handleShopifyCallback())
	s.Router.HandleFunc("/shopify", s.handleShopifyHome())

	s.HttpServer = &http.Server{
		Handler:      s.Router,
		Addr:         fmt.Sprintf("127.0.0.1:%d", s.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return s, nil
}

// Start http server. This method is blocking.
func (s *Server) Start() error {
	err := s.HttpServer.ListenAndServe()
	return err
}

// Shutdown stops the HTTP Server
func (s *Server) Shutdown() error {

	// clean up for shutdown here!

	//shutdown http server
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := s.HttpServer.Shutdown(ctx)
	if err != nil {
		return err
	}
	log.Println("Server shutdown gracefully")
	return nil
}

// ServerReady Blocks until the server returns a 200 status code from the /ping endpoint
func (s *Server) ServerReady() bool {

	var client = &http.Client{
		Timeout: time.Second * 5,
	}

	tries := 0
	for {
		tries++
		url := fmt.Sprintf("http://127.0.0.1:%d/ping", s.Port)
		resp, _ := client.Get(url)
		if resp.StatusCode == http.StatusOK {
			return true
		} else {
			time.Sleep(time.Second * 1)
		}

		if tries >= 5 {
			panic("Unable to connect to test HTTP server at " + url)
		}

	}

}
func (s *Server) handlePing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	}
}
