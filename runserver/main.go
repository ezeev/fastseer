package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/ezeev/fastseer"
)

func main() {
	server, err := fastseer.NewServer("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go server.Start()

	<-stop
	server.Shutdown()

}
