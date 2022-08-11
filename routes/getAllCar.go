package routes

import (
	"encoding/json"
	"github.com/Julsan26/Project/database"
	"github.com/Julsan26/Project/request"
	"net/http"
)

// Get all Cars

// response and request handlers
func GetAllCars(w http.ResponseWriter, r *http.Request) {

	db := database.SetupDB()

	printMessage("Getting cars...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM cars")

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var cars []request.Car

	// Foreach movie
	for rows.Next() {
		var id int64
		var make string
		var model string
		var Package string
		var color string
		var year string
		var catagory string
		var mileage string
		var price string

		err = rows.Scan(&id, &make, &model, &Package, &color, &year, &catagory, &mileage, &price)
		checkErr(err)

		cars = append(cars, request.Car{
			ID:       id,
			Make:     make,
			Model:    model,
			Package:  Package,
			Color:    color,
			Year:     year,
			Catagory: catagory,
			Mileage:  mileage,
			Price:    price,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)

	return
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
