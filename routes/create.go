package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Julsan26/Project/request"
)

func CreateCarByID(w http.ResponseWriter, r *http.Request) {
	make := r.FormValue("Make")
	model := r.FormValue("Model")

	response := request.JsonResponse{}

	if make == "" || model == "" {
		response = request.JsonResponse{Type: "error", Message: "You are missing make or model parameter."}
	} else {
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
		printMessage("Inserting movie into DB")

		fmt.Println("Inserting new movie with ID: " + make + " and name: " + model)

		var lastInsertID int
		errs := db.QueryRow("INSERT INTO movies(Make, Model) VALUES($1, $2) returning id;", make, model).Scan(&lastInsertID)

		// check errors
		checkErr(errs)

		response = request.JsonResponse{Type: "success", Message: "The movie has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

// Function for handling messages
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}
