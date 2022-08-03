package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"remine-api/database/mssql"

	log "github.com/sirupsen/logrus"
)

func main() {
	//connect database
	configPath := flag.String("config", "configure", "set configs path, default as: 'configure'")
	mssql.ConnectDB(*configPath)

	//start http route
	portApp := flag.String("portApp", "8080", "portApp number")

	flag.Parse()
	log.Infof("Run on port : %+v", *portApp)

	//start routes
	r := handler.Routes{}
	handlerRoute := r.InitTransactionRoute()
	AppServer := &http.Server{
		Addr:    fmt.Sprint(":", *portApp),
		Handler: handlerRoute,
	}

	//start http app server
	go func() {
		if err := AppServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panicf("Transaction listen : %s\n", err)
		} else if err != nil {
			log.Panicf("Transaction listen error: %s\n", err)
		}
		log.Infof("Transaction listen at port: %s", *portApp)
	}()

	//wait signals
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals //wait for SIGINT
}
