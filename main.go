package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/AlyRagab/hashing-password/pkg/config"
	"github.com/gorilla/mux"
)

func init() {
	db, err := config.ConnectDb()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
}

// HTTP 404 NotFound
func notfound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "404 Page not found")
}

func password(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Hello Password")
	db, err := config.ConnectDb()
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	insertStatement := (`INSERT INTO hashed_password(password) VALUES($1)`)
	_, err = db.Exec(insertStatement, string(reqBody))
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		db, err := config.ConnectDb()
		if err != nil {
			fmt.Fprintf(w, "Not healthy %d", http.StatusInternalServerError)
			log.Error("database connection not heathy")
		} else {
			fmt.Fprintf(w, "Healthy %d", http.StatusOK)
		}
		defer db.Close()
	})
	log.Info("Starting Server on 0.0.0.0:8080")
	r.HandleFunc("/password", password).Methods("POST")
	r.NotFoundHandler = http.HandlerFunc(notfound)
	http.ListenAndServe(":8080", r)
}
