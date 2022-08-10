package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Julsan26/Project/database"
	"github.com/Julsan26/Project/request"
)

func CreateCarByID(w http.ResponseWriter, r *http.Request) {
	make := r.FormValue("Make")
	model := r.FormValue("Model")

	var response = request.JsonResponse{}

	if make == "" || model == "" {
		response = request.JsonResponse{Type: "error", Message: "You are missing make or model parameter."}
	} else {
		db := database.SetupDB()

		printMessage("Inserting movie into DB")

		fmt.Println("Inserting new movie with ID: " + make + " and name: " + model)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO movies(Make, Model) VALUES($1, $2) returning id;", make, model).Scan(&lastInsertID)

		// check errors
		checkErr(err)

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
