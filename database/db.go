package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// DB set up
func setupDB() *sql.DB {
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	dbPassword := os.Getenv("PASSWORD")

	dbURI := fmt.Sprintf("host=%s user=%s dbName=%s dbPassword=%s port=%s", host, user, dbName, dbPassword, dbPort)

	db, err := sql.Open("postgres", dbURI)

	if err != nil {
		log.Fatal(err)
	}
	return db
}
