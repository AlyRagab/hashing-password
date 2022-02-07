package provider

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

// Connecting to database
func ConnectDb() (*sql.DB, error) {
	var (
		host     = os.Getenv("DBHOST")
		port     = 5432
		dbuser   = os.Getenv("DBUSER")
		password = os.Getenv("DBPASS")
		dbname   = os.Getenv("DBNAME")
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, dbuser, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	} else {
		log.Info("Successfully connected to database")
	}
	return db, nil
}
