package routes

import (
	"encoding/json"
	"github.com/Julsan26/Project/database"
	"github.com/Julsan26/Project/request"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

const (
	query = "SELECT id,Make, Model,Package,Color,Year,Category,Mileage,Price FROM cars WHERE id=?;"
)

// response and request handlers
func GetCarsByID(w http.ResponseWriter, r *http.Request) {
	var car request.Car

	var dbid int64
	var Make string
	var model string
	var Package string
	var color string
	var year string
	var catagory string
	var mileage string
	var price string

	vars := mux.Vars(r)
	id := vars["id"]
	newID, _ := strconv.ParseInt(id, 10, 64)
	sqlStatement := `SELECT id,make,model,package,color,year,mileage,price FROM cars WHERE id=$1;`
	db := database.SetupDB()

	row := db.QueryRow(sqlStatement, newID)
	err := row.Scan(&dbid, &Make, &model, &Package, &color, &year, &mileage, &price)
	if err != nil {
		return
	}

	car = request.Car{
		ID:       dbid,
		Make:     Make,
		Model:    model,
		Package:  Package,
		Color:    color,
		Year:     year,
		Catagory: catagory,
		Mileage:  mileage,
		Price:    price,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)

}

// Function for handling errors
