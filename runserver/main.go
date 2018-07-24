package main

import (
	"os"
	"os/signal"

	"github.com/ezeev/fastseer"
	"github.com/ezeev/fastseer/logger"
)

func main() {
	server, err := fastseer.NewServer("config.yaml")
	if err != nil {
		logger.Fatal("None", err.Error())
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go server.Start()

	<-stop
	server.Shutdown()

}
