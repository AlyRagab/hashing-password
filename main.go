package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/AlyRagab/hashing-password/provider"
	log "github.com/sirupsen/logrus"
	bcrypt "golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
)

func init() {
	var err error
	provider.DB, err = provider.ConnectDb()
	if err != nil {
		log.Fatalln(err)
	}
}

// HTTP 404 NotFound
func notfound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "404 Page not found")
}

func password(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	password := string(string(reqBody))
	// Hashing the passowrd using Salt with SALT_SECRET env var
	salt := []byte(password + os.Getenv("SALT_SECRET"))
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(salt), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}
	hashedPassword := string(hashedPasswordBytes)

	// Inserting into database
	insertStatement := (`INSERT INTO hashed_password(password) VALUES($1)`)
	_, err = provider.DB.Exec(insertStatement, hashedPassword)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		err := provider.DB.Ping()
		if err != nil {
			fmt.Fprintf(w, "Not healthy %d", http.StatusInternalServerError)
			log.Error("database connection not heathy")
		} else {
			fmt.Fprintf(w, "Healthy %d", http.StatusOK)
		}
	})
	log.Info("Starting Server on 0.0.0.0:8080")
	r.HandleFunc("/password", password).Methods("POST")
	r.NotFoundHandler = http.HandlerFunc(notfound)
	http.ListenAndServe(":8080", r)
}
