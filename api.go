package main

import (
	"net/http"

	"bodotech.net/fizzbuzz/controllers"
	log "github.com/sirupsen/logrus"
)

func initLogger() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	log.SetLevel(log.InfoLevel)
}

func main() {
	initLogger()
	log.Print("Starting API service")

	handler := controllers.HandleRequests()
	log.Fatal(http.ListenAndServe(":7001", handler))
}
