package routes

import (
	"encoding/json"
	"github.com/Julsan26/Project/database"
	mod "github.com/Julsan26/Project/model"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// GetCarsByID will Get Cars from the database based on id from the request.
func GetCarsByID(w http.ResponseWriter, r *http.Request) {
	var car mod.Car

	var dbID int64
	var Make string
	var model string
	var Package string
	var color string
	var year string
	var category string
	var mileage string
	var price string

	vars := mux.Vars(r)
	id := vars["id"]
	newID, _ := strconv.ParseInt(id, 10, 64)
	sqlStatement := `SELECT id,make,model,package,color,year,mileage,price FROM cars WHERE id=$1;`
	db := database.SetupDB()

	row := db.QueryRow(sqlStatement, newID)
	err := row.Scan(&dbID, &Make, &model, &Package, &color, &year, &mileage, &price)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}
	car = mod.Car{
		ID:       dbID,
		Make:     Make,
		Model:    model,
		Package:  Package,
		Color:    color,
		Year:     year,
		Category: category,
		Mileage:  mileage,
		Price:    price,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)

}
