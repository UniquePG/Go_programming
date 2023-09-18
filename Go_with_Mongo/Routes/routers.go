package routes

import (
	"github.com/gorilla/mux"
	controllers "github.com/uniquepg/gowithmongo/Controllers"
)

func Routes() *mux.Router {
	// create new router instance
	router := mux.NewRouter();

	router.HandleFunc("/getmovies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/create", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/update/{id}", controllers.UpdateMovieAsWatched).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/deleteall", controllers.DeleteAllRecords).Methods("DELETE")

	return router
}