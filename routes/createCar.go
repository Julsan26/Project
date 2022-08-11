package routes

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Julsan26/Project/request"
	"io/ioutil"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "julsanmagaju"
)

func CreateCar(w http.ResponseWriter, r *http.Request) {
	var car request.Car
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	decodeErr := json.NewDecoder(bytes.NewReader(body)).Decode(&car)
	if decodeErr != nil {
		return
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	printMessage("Inserting movie into DB")

	var lastInsertID int
	errs := db.QueryRow("INSERT INTO cars(id ,Make, Model,Package,Color,Year,Category,Mileage,Price) VALUES($1, $2,$3,$4,$5,$6,$7,$8,$9) returning id;", car.ID, car.Make, car.Model, car.Package, car.Color, car.Year, car.Catagory, car.Mileage, car.Price).Scan(&lastInsertID)
	if err == nil {
		printMessage("Database is sucesfully updated")

	}
	checkErr(errs)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(r)

	return

}

// Function for handling messages
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}
