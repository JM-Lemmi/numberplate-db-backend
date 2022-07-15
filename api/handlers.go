package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

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

func indexHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestLogger := log.WithFields(log.Fields{"client": GetIP(r), "profile": vars["profile"]})
	requestLogger.Infoln("New Request!")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}

func numberplateHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	requestLogger := log.WithFields(log.Fields{"client": GetIP(r)})
	requestLogger.Infoln("New Request!")

	// from https://golangdocs.com/golang-postgresql-example
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	requestLogger.Infoln("PostgreSQL connected!")

	//from https://golangdocs.com/golang-postgresql-example
	insertDynStmt := `INSERT INTO "numberplates"("plate", "owner") values($1, $2)`
	_, e := db.Exec(insertDynStmt, "HGJL1999", "Julian Lemmerich")
	CheckError(e)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("check console"))
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
