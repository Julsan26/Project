package routes

import (
	"encoding/json"
	"github.com/Julsan26/Project/database"
	"github.com/Julsan26/Project/model"
	"net/http"
)

// Get all Cars

// response and model handlers
func GetAllCars(w http.ResponseWriter, r *http.Request) {

	db := database.SetupDB()

	printMessage("Getting cars...")

	// Get all cars from cars table
	rows, err := db.Query("SELECT * FROM cars")
	if err != nil {
		panic(err)
	}

	var cars []model.Car

	// Foreach car
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
		if err != nil {
			panic(err)
		}

		cars = append(cars, model.Car{
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
