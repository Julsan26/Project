package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Julsan26/Project/routes"
	"github.com/gorilla/mux"
)

func main() {

	// Init the mux router
	router := mux.NewRouter()

	// Route handles & endpoints

	// Get all Cars
	router.HandleFunc("/cars", routes.GetCars).Methods("GET")

	// Get a Car by ID
	router.HandleFunc("/cars/{id}", routes.GetCarsByID).Methods("GET")

	// Create a Car
	router.HandleFunc("/cars/{id}", routes.CreateCarByID).Methods("POST")

	// serve the app
	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8000", router))
}
