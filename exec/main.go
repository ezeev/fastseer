package main

import (
	"os"
	"os/signal"

	"github.com/ezeev/fastseer"
	"github.com/ezeev/fastseer/logger"
)

// @APIVersion 1.0.0
// @APITitle FastSeer Admin API
// @APIDescription My API usually works as expected.
// @Contact developer@fastseer.com
// @BasePath https://shopify-app.fastseer.com/api/
func main() {

	conf := os.Args[1]

	server, err := fastseer.NewServer(conf)
	if err != nil {
		logger.Fatal("None", err.Error())
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go server.Start()

	<-stop
	server.Shutdown()

}
