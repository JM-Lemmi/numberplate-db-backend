package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	log "github.com/sirupsen/logrus"
)

type Numberplate struct {
	Plate   string `json:"plate"`
	Country string `json:"country"`
	Owner   string `json:"owner"`
	Notes   string `json:"notes"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	requestLogger := log.WithFields(log.Fields{"client": GetIP(r)})
	requestLogger.Infoln("new request for /")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}

func numberplateHandler(w http.ResponseWriter, r *http.Request) {
	requestLogger := log.WithFields(log.Fields{"client": GetIP(r)})
	requestLogger.Infoln("new request for /numberplates")

	//from https://golangdocs.com/golang-postgresql-example
	rows, err := db.Query(`SELECT "plate", "owner" FROM "numberplates"`)
	CheckError(err)

	output := ""
	defer rows.Close()
	for rows.Next() {
		var plate string
		var owner string

		err = rows.Scan(&plate, &owner)
		CheckError(err)

		output = (output + plate + ", " + owner + "\n")
	}

	CheckError(err)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(output))
}

func numberplateNewHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	requestLogger := log.WithFields(log.Fields{"client": GetIP(r)})
	requestLogger.Infoln("new request for /numberplates/new")

	d := json.NewDecoder(r.Body)
	p := &Numberplate{}
	err := d.Decode(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	requestLogger.Infoln(p.Plate, p.Country, p.Owner, p.Notes)
	resp, e := db.Exec(`INSERT INTO "numberplates"("plate", "country", "owner", "notes") values($1, $2, $3, $4)`, p.Plate, p.Country, p.Owner, p.Notes)
	CheckError(e)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	respid, _ := resp.LastInsertId()
	w.Write([]byte(fmt.Sprint(respid)))
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
