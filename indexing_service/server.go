package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ezeev/fastseer/logger"
	"github.com/ezeev/fastseer/search"
	"github.com/gorilla/mux"
)

type Server struct {
	Router     *mux.Router
	HttpServer *http.Server
	Port       int
	Search     search.SearchEngine
}

func NewServer(port int, searchImpl string) (*Server, error) {

	s := &Server{}
	s.Port = port

	searchEngine, err := search.NewSearchEngine(searchImpl)
	if err != nil {
		return nil, err
	}
	s.Search = searchEngine

	// routes
	s.Routes()

	s.HttpServer = &http.Server{
		Handler:      s.Router,
		Addr:         fmt.Sprintf("127.0.0.1:%d", s.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return s, nil
}

func (s *Server) Start() error {
	err := s.HttpServer.ListenAndServe()
	return err
}

func (s *Server) Shutdown() error {

	// clean up for shutdown here!

	//shutdown http server
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := s.HttpServer.Shutdown(ctx)
	if err != nil {
		return err
	}
	logger.Info("None", "Server shutdown gracefully")
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
