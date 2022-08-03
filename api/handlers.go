package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	log "github.com/sirupsen/logrus"
	"path/filepath"
)

type Numberplate struct {
	Plate   string `json:"plate"`
	Country string `json:"country"`
	Owner   string `json:"owner"`
	Notes   string `json:"notes"`
}

type Meet struct {
	ID       uuid.UUID `json:"id"`
	Plate    string    `json:"plate"`
	Time     time.Time `json:"date"`
	Lat      float32   `json:"lat"`
	Lon      float32   `json:"lon"`
	Image    bool      `json:"image"`
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
	if r.Method == "GET" {
		requestLogger.Infoln("new request for GET /numberplates")

		// db query. get all rows
		rows, err := db.Query(`SELECT "plate", "country", "owner", "notes" FROM "numberplates"`)
		CheckError(err)

		// loop over rows and read into Numberplate array
		var output []Numberplate
		defer rows.Close()
		for rows.Next() {
			var entry Numberplate

			err = rows.Scan(&entry.Plate, &entry.Country, &entry.Owner, &entry.Notes)
			CheckError(err)

			output = append(output, entry)
		}

		// turn into json and respond
		response, err := json.Marshal(output)
		CheckError(err)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(response)

	} else if r.Method == "PUT" {
		requestLogger.Infoln("new request for PUT /numberplates")

		d := json.NewDecoder(r.Body)
		p := &Numberplate{}
		err := d.Decode(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		requestLogger.Debugln(p.Plate, p.Country, p.Owner, p.Notes)

		if PlateExists(p.Plate) {
			resp, e := db.Exec(`UPDATE "numberplates" SET country=$2, owner=$3, notes=$4 WHERE plate=$1`, p.Plate, p.Country, p.Owner, p.Notes)
			CheckError(e)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			respid, _ := resp.RowsAffected()
			w.Write([]byte(fmt.Sprint(respid)))
		} else {
			resp, e := db.Exec(`INSERT INTO "numberplates"("plate", "country", "owner", "notes") values($1, $2, $3, $4)`, p.Plate, p.Country, p.Owner, p.Notes)
			CheckError(e)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			respid, _ := resp.RowsAffected()
			w.Write([]byte(fmt.Sprint(respid)))
		}
	} else {
		w.Write([]byte("Disallowed Method"))
	}
}

func numberplatePlateHandler(w http.ResponseWriter, r *http.Request) {
	requestLogger := log.WithFields(log.Fields{"client": GetIP(r)})
	// get plate
	vars := mux.Vars(r)
	plate := vars["plate"]
	requestLogger.Infoln("new " + r.Method + " request for /numberplate/:plate with plate " + plate)

	if r.Method == "GET" {
		// db query. only need one row, since plate is unique
		row := db.QueryRow(`SELECT "plate", "country", "owner", "notes" FROM "numberplates" WHERE plate=$1`, plate)
		// read result into numberplate struct
		var entry Numberplate
		row.Scan(&entry.Plate, &entry.Country, &entry.Owner, &entry.Notes)

		// turn into json and respond
		response, err := json.Marshal(entry)
		CheckError(err)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(response)

	} else if r.Method == "DELETE" {
		_, err := db.Exec(`DELETE "numberplate" WHERE plate=$1`, plate)
		CheckError(err)
		w.Write([]byte("deleted " + plate))
	}
}

func meetsHandler(w http.ResponseWriter, r *http.Request) {
	requestLogger := log.WithFields(log.Fields{"client": GetIP(r)})
	if r.Method == "GET" {
		requestLogger.Infoln("new request for GET /meets")

		// db query. get all rows
		rows, err := db.Query(`SELECT "id", "plate", "lat", "lon", "time", "image" FROM "meets"`)
		CheckError(err)

		// loop over rows and read into Numberplate array
		var output []Meet
		defer rows.Close()
		for rows.Next() {
			var meet Meet

			err = rows.Scan(&meet.ID, &meet.Plate, &meet.lat, &meet.lon, &meet.Time, &meet.Image)
			CheckError(err)

			output = append(output, meet)
		}

		// turn into json and respond
		response, err := json.Marshal(output)
		CheckError(err)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(response)

	} else if r.Method == "PUT" {
		requestLogger.Infoln("new request for PUT /meets")

		d := json.NewDecoder(r.Body)
		m := &Meet{}
		err := d.Decode(m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// client generates a UUID for a new entry. If this turns out to be a mistake later, there has to be logic to handle missing ID.
		if MeetExists(m.ID) {
			resp, e := db.Exec(`UPDATE "meets" SET plate=$2, lat=$3, lon=$4, time=$5, image=$6 WHERE id=$1`, m.ID, m.Plate, m.lat, m.lon, m.Time, m.Image)
			CheckError(e)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			respid, _ := resp.RowsAffected()
			w.Write([]byte(fmt.Sprint(respid)))
		} else {
			resp, e := db.Exec(`INSERT INTO "meets"("id", "plate", "lat", "lon", "time", "image") values($1, $2, $3, $4, $5, $6)`, m.ID, m.Plate, m.lat, m.lon, m.Time, m.Image)
			CheckError(e)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			respid, _ := resp.RowsAffected()
			w.Write([]byte(fmt.Sprint(respid)))
		}

	}
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	requestLogger := log.WithFields(log.Fields{"client": GetIP(r)})
    vars := mux.Vars(r)
	id := vars["id"]
	requestLogger.Infoln("new " + r.Method + " request for uploading image for id " + id)

	if r.Method == "POST" {
		// name should be "file"
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()

		// from https://github.com/mayth/go-simple-upload-server/blob/master/server.go
		
		body, err := ioutil.ReadAll(srcFile)
		if err != nil {
			logger.WithError(err).Error("failed to read the uploaded content")
			w.WriteHeader(http.StatusInternalServerError)
			writeError(w, err)
			return
		}

		filename := id + filepath.Ext(info.Filename)

		dstPath := path.Join("/var/www/html/images/", filename) //hardcoded image path TODO make configurable
		dstFile, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			logger.WithError(err).WithField("path", dstPath).Error("failed to open the file")
			w.WriteHeader(http.StatusInternalServerError)
			writeError(w, err)
			return
		}
		defer dstFile.Close()
		if written, err := dstFile.Write(body); err != nil {
			logger.WithError(err).WithField("path", dstPath).Error("failed to write the content")
			w.WriteHeader(http.StatusInternalServerError)
			writeError(w, err)
			return
		} else if int64(written) != size {
			logger.WithFields(logrus.Fields{
				"size":    size,
				"written": written,
			}).Error("uploaded file size and written size differ")
			w.WriteHeader(http.StatusInternalServerError)
			writeError(w, fmt.Errorf("the size of uploaded content is %d, but %d bytes written", size, written))
		}

		w.WriteHeader(http.StatusOK)
	} else if r.Method == "GET" {
		// GET request returns the image
		ServeFile(w, r, "/var/www/html/images/" + id + ".jpg")
	}
}

func PlateExists(plate string) bool {
	row := db.QueryRow(`SELECT "plate" from "numberplates" WHERE plate=$1`, plate)
	temp := ""
	row.Scan(&temp)
	if temp != "" {
		return true
	}
	return false
}

func MeetExists(uuid uuid.UUID) bool {
	row := db.QueryRow(`SELECT "ID" from "meets" WHERE ID=$1`, uuid)
	temp := ""
	row.Scan(&temp)
	if temp != "" {
		return true
	}
	return false
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
