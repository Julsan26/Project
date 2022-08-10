package routes

import (
	"net/http"
)

// Get all Cars

// response and request handlers
func GetCars(w http.ResponseWriter, r *http.Request) {

	// db := database.SetupDB()

	// // Get all cars from cars "
	// rows, err := db.Query("SELECT * FROM cars")

	// // check errors
	// checkErr(err)

	// // var response []JsonResponse
	// var cars []request.Car

	// // Foreach movie
	// for rows.Next() {
	// 	var id int
	// 	var movieID string
	// 	var movieName string

	// 	err = rows.Scan(&id, &movieID, &movieName)

	// 	// check errors
	// 	checkErr(err)

	// 	movies = append(movies, Movie{MovieID: movieID, MovieName: movieName})
	// }

	// var response = JsonResponse{Type: "success", Data: movies}

	// json.NewEncoder(w).Encode(response)
	return
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
