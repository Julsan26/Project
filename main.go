package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Julsan26/Project/routes"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func main() {

	//Loading Env

	dialect := os.Getenv("postgres")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	dbPassword := os.Getenv("PASSWORD")

	dbURI := fmt.Sprintf("host=%s user=%s dbName=%s dbPassword=%s port=%s", host, user, dbName, dbPassword, dbPort)

	db, err := gorm.Open(dialect, dbURI)

	// Init the mux router
	router := mux.NewRouter()

	// Route handles & endpoints

	// Get all Cars
	router.HandleFunc("/cars", routes.GetCars).Methods("GET")

	// Get a Car by ID
	router.HandleFunc("/cars/{id}", routes.CreateCarByID).Methods("GET")

	// Create a Car
	router.HandleFunc("/cars/{id}", routes.CreateCarByID).Methods("POST")

	// serve the app
	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8000", router))
}
