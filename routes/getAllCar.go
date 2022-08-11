package routes

import (
	"encoding/json"
	"github.com/Julsan26/Project/database"
	mod "github.com/Julsan26/Project/model"
	"log"
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
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(err.Error())
		log.Println(err.Error())
	}

	var cars []mod.Car

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
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(err.Error())
			log.Println(err.Error())
			return
		}

		cars = append(cars, mod.Car{
			ID:       id,
			Make:     make,
			Model:    model,
			Package:  Package,
			Color:    color,
			Year:     year,
			Category: catagory,
			Mileage:  mileage,
			Price:    price,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)

	return
}
