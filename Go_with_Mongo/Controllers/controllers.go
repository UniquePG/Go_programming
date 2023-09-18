package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	helpers "github.com/uniquepg/gowithmongo/Helpers"
	model "github.com/uniquepg/gowithmongo/Model"
)

// ! now write our actual controllers
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/x-www-form-urlencode")

	allMovies := helpers.GetAllMovies()

	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST") // allow only post method

	//crete movie instance
	var movie model.Netflix

	// accept data from body and decode it into movie struct
	_ = json.NewDecoder(r.Body).Decode(&movie)

	// now call our insertOne helper method
	helpers.InsertOne(movie)

	// now send back the response with Encode
	json.NewEncoder(w).Encode(movie) // encode to JSON format

}

// update movie
func UpdateMovieAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/x-www-form-urlencode")
	// w.Header().Set("Allow-Control-Allow-Methods", "POST") // allow only post method

	// take id from params
	params := mux.Vars(r)

	// now call our update method
	// helpers.UpdateValue(params["id"])

	//* here we take updated values from the body
	var movie model.Netflix

	json.NewDecoder(r.Body).Decode(&movie)

	helpers.UpdateValue(params["id"], movie) // update our movie

	// now send back response with id
	json.NewEncoder(w).Encode(movie)
}

// * delete one movie
func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/x-www-form-urlencode")

	// take id form the params
	params := mux.Vars(r)

	helpers.DeleteOneRecord(params["id"])

	fmt.Fprintf(w, "Successfully deleted one record...")
}

// * delete all records
func DeleteAllRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/x-www-form-urlencode")

	// now call the delete all documents helper
	helpers.DeleteAllDocuments()

	fmt.Fprintf(w, "All documents deleted successfully...")

}
