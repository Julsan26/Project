package routes

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Julsan26/Project/model"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "julsanmagaju"
)

func CreateCar(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	newID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err.Error())
		log.Println(err)
		return
	}

	var car model.Car
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err.Error())
		log.Println(err)
		return
	}

	decodeErr := json.NewDecoder(bytes.NewReader(body)).Decode(&car)
	if decodeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(decodeErr.Error())
		log.Println(decodeErr)
		return
	}
	if newID != car.ID {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("The id the request and body are not identical")
		log.Println("The id the request and body are not identical")
		return
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(err.Error())
		log.Println(err.Error())
		return
	}

	printMessage("Inserting car into DB")

	var lastInsertID int
	errs := db.QueryRow("INSERT INTO cars(id ,Make, Model,Package,Color,Year,Category,Mileage,Price) VALUES($1, $2,$3,$4,$5,$6,$7,$8,$9) returning id;", car.ID, car.Make, car.Model, car.Package, car.Color, car.Year, car.Category, car.Mileage, car.Price).Scan(&lastInsertID)
	if errs != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(errs)
		log.Println(errs)
		return
	}
	printMessage("Database is sucesfully updated")

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode("New Car Created")

	return

}

// Function for handling messages
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}
