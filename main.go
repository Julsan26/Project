package main

import (
	"fmt"
	"github.com/Julsan26/Project/database"
	"log"
	"net/http"

	"github.com/Julsan26/Project/routes"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db := database.SetupDB()
	fmt.Println(db)
	fmt.Println("Successfully connected!")

	router := mux.NewRouter()

	// Route handles & endpoints

	// Get all Cars
	router.HandleFunc("/cars", routes.GetAllCars).Methods("GET")

	// Get a Car by ID
	router.HandleFunc("/cars/{id}", routes.GetCarsByID).Methods("GET")

	// Create a Car
	router.HandleFunc("/cars/{id}", routes.CreateCar).Methods("POST")

	// serve the app
	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8000", router))
}
