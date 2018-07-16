package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	server, err := NewServer(8083, "solr")
	if err != nil {
		log.Fatal(err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go server.Start()

	<-stop
	server.Shutdown()

}
