// @APIVersion 1.0.0
// @APITitle Teamwork Desk
// @APIDescription Bend Teamwork Desk to your will using these read and write endpoints
// @Contact support@teamwork.com
// @TermsOfServiceUrl https://www.teamwork.com/termsofservice
// @License BSD
// @LicenseUrl http://opensource.org/licenses/BSD-2-Clause
package main

import (
	"os"
	"os/signal"

	"github.com/ezeev/fastseer"
	"github.com/ezeev/fastseer/logger"
)

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
