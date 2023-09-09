package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// declair variables for databse connection
var Dbvar *gorm.DB                                                                           // this var point to the gorm db
var err error                                                                                // this is for error
var DNS = "root:12345@tcp(127.0.0.1:3306)/gotestdb?charset=utf8mb4&parseTime=True&loc=Local" // this is the url of the db

// make struct for the user
type User struct {
	//* initilize this struct as the gorm model
	gorm.Model //* with that this struct treated as model for database

	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

// make a func to connect the db (Migration for db)
func InitialMigration() {
	// connect to db
	Dbvar, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})

	// check for errors
	if err != nil {
		fmt.Println(err.Error())
		panic("Can not connect to db")
	}

	// now automigrate our struct(user) to db -> it will make table automatically in the db
	Dbvar.AutoMigrate(&User{}) // reference to our User

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	//* first we set the content type to the header
	w.Header().Set("Content-Type", "application/json")

	// now make varible of the user struct
	var user User

	// now get data from the body and decode it into our user struct
	json.NewDecoder(r.Body).Decode(&user) 

	// save the data
	Dbvar.Create(&user) // pass our user reference

	// now pass the data base to the browser
	json.NewEncoder(w).Encode(user)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// make slice to store the user info
	var users []User

	Dbvar.Find(&users) // find the all users from the db

	// encode the data and  sent back to clint as response
	json.NewEncoder(w).Encode(users) // encode the json format

}

func GetUserById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// get the id from the params	-> Vars(r) it returns a map[ id: <id> ]
	params := mux.Vars(r) // 	pass the request(r) -> we get all data(params) from the url

	fmt.Println("params: ", params)
	// make variable to store the user
	var user User

	// find the user based on the id passed in the url
	Dbvar.First(&user, params["id"]) // get the id from params and find user in db

	json.NewEncoder(w).Encode(user) // seth the response of the user

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	// for updation first we will find the user by id and then we get the data from body and then save it
	w.Header().Set("Content-Type", "application/json")

	// get the id from the params
	params := mux.Vars(r) // 	pass the request(r) -> we get all data(params) from the url

	// make variable to store the user
	var user User

	// find the user based on the id passed in the url
	Dbvar.First(&user, params["id"]) // get the id from params and find user in db

	// now we get the updated data from the body
	json.NewDecoder(r.Body).Decode(&user)

	Dbvar.Save(&user) //save the new user (updated)

	fmt.Println(user) // print the updated user

	json.NewEncoder(w).Encode(user) // seth the response of the user

}
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// get the id from the params
	params := mux.Vars(r) // 	pass the request(r) -> we get all data(params) from the url

	fmt.Println("params: ", params)
	// make variable to store the user
	var user User

	// find the user based on the id passed in the url
	Dbvar.First(&user, params["id"]) // get the id from params and find user in db

	Dbvar.Delete(&user) // delete the user of a perticular id

	// now sent back the response to the client
	json.NewEncoder(w).Encode("User deleted successfully!")
}
