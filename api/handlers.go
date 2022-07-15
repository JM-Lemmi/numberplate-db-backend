package main

import (
	"net/http"

	_ "github.com/lib/pq"

	log "github.com/sirupsen/logrus"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	requestLogger := log.WithFields(log.Fields{"client": GetIP(r)})
	requestLogger.Infoln("New Request!")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}

func numberplateHandler(w http.ResponseWriter, r *http.Request) {
	requestLogger := log.WithFields(log.Fields{"client": GetIP(r)})
	requestLogger.Infoln("New Request!")

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
	requestLogger.Infoln("New Request!")

	//from https://golangdocs.com/golang-postgresql-example
	insertDynStmt := `INSERT INTO "numberplates"("plate", "country", "owner") values($1, $2, $3)`
	_, e := db.Exec(insertDynStmt, "HGJL1999", "DEU", "Julian Lemmerich")
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
