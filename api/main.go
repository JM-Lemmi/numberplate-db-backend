package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const (
	host     = "localhost"
	port     = 25432
	user     = "postgres"
	password = "password"
	dbname   = "numberplates"
)

var version = "0.1"

// create router
var router *mux.Router

// from https://golangdocs.com/golang-postgresql-example
// create database
var db, _ = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))

func main() {
	log.Info("Welcome to numberplate-db-backend, version " + version)

	// setup router
	router = mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/numberplates", numberplateHandler)
	router.HandleFunc("/numberplates/{plate}", numberplatePlateHandler)

	// listen and serve
	log.Fatal(http.ListenAndServe("0.0.0.0:80", router))
}
