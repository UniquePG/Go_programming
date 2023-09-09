package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Dbvar *gorm.DB
var err error
var DNS = "root:12345@tcp(127.0.0.1:3306)/go_ecommerce?charset=utf8mb4&parseTime=True&loc=Local"

type User struct {
	gorm.Model

	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func DbConnection() {

	Dbvar, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error(), http.StatusInternalServerError)
		return

	}

	fmt.Println("Database connected...")

	Dbvar.AutoMigrate(&User{})

}

func main() {

	DbConnection()

	router := mux.NewRouter()

	router.HandleFunc("/signup", SignupHandler).Methods("POST")
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.HandleFunc("/logout", LogoutHandler).Methods("GET")
	router.HandleFunc("/profile", ProfileHandler).Methods("GET")

	fmt.Println("server started...")
	log.Fatal(http.ListenAndServe(":5000", router))
}
