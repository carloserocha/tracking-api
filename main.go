package main

import (
	"net/http"

	"scraper/correios"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	router := mux.NewRouter()

	router.HandleFunc("/tracking/{id}", correios.getDetails).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
