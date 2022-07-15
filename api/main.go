package main

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var version = "0.1"

var router *mux.Router

var templateBox *rice.Box

func main() {
	log.Info("Welcome to numberplate-db-backend, version " + version)

	// setup router
	router = mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/numberplates/new", numberplateNewHandler)
	router.HandleFunc("/numberplates", numberplateHandler)

	// listen and serve
	log.Fatal(http.ListenAndServe("0.0.0.0:80", router))
}
