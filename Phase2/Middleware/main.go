package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouters() {

	r := mux.NewRouter() // make a new router instance

	// make different routes to handle crud operation
	r.HandleFunc("/createuser", Check(CreateUser)).Methods("POST")            // for creating user
	r.HandleFunc("/getusers", Authentication(GetUsers)).Methods("GET") // get all users
	r.HandleFunc("/getuser/{id}", GetUserById).Methods("GET")          // get user by id
	r.HandleFunc("/update/{id}", UpdateUser).Methods("PUT")            // update a user
	r.HandleFunc("/delete/{id}", DeleteUser).Methods("DELETE")         // Delete a user

	fmt.Println("Server is running..")
	// now start server
	log.Fatal(http.ListenAndServe(":9001", r))
}

func main() {

	// connect to our db and automigration
	InitialMigration()

	// call our initialize routers func here
	initializeRouters()

}
