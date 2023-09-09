package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func SignupHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	// check Route
	if req.URL.Path != "/signup" {
		http.Error(res, "Route Not Found...", http.StatusBadRequest)
		return
	}

	// check method
	if req.Method != "POST" {

		http.Error(res, "This is wrong route method...", http.StatusBadRequest)
		return
	}

	var user User

	json.NewDecoder(req.Body).Decode(&user)

	// hash our password
	pass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	user.Password = string(pass)

	Dbvar.Create(&user)

	json.NewEncoder(res).Encode(user)

	res.Write([]byte("User Register successfully \n"))

}

func LoginHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	// check Route
	if req.URL.Path != "/login" {
		http.Error(res, "Route Not Found...", http.StatusBadRequest)
		return
	}

	// check method
	if req.Method != "POST" {

		http.Error(res, "This is wrong route method...", http.StatusBadRequest)
		return
	}

	var user User

	json.NewDecoder(req.Body).Decode(&user)

	// implement a func that gives user with provided email id
	userDb, e := findUserByEmail(user.Email)	

	if e != nil {
		http.Error(res, "User not found with provided email address \n", http.StatusNotFound)
		return
	}


	err := bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(user.Password))

	if err != nil {
		http.Error(res, "Wrong Password, try again \n", http.StatusNotFound)
		return
	}


	session, er := CreateSession(res, req, user.Email)

	if er != nil{
		fmt.Fprintf(res, "Session values not set")
		return
	}

	err = session.Save(req, res)

	if err != nil {
		http.Error(res, "Error while saving session...", http.StatusInternalServerError)
		return
	}


	
	json.NewEncoder(res).Encode(userDb)
	
	fmt.Fprintf(res, "Session set successfully\n")
	fmt.Fprintf(res, "User logged in successfully... Password matched\n")

	// redirect to other route
	// http.Redirect(res, req, "/profile", http.StatusSeeOther)
}


func LogoutHandler(res http.ResponseWriter, req *http.Request){
	
	// get the session
	session, err := getSession(res, req);

	if err != nil {
		http.Error(res, "Session not found...", http.StatusNotFound)
	}

	session.Values["IsLoggedIn"] = false
	session.Values["email"] = ""

	err = session.Save(req, res)

	if err != nil {
		http.Error(res, "Error while saving session...", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(res, "You are logged out successfully!");

}



func ProfileHandler(res http.ResponseWriter, req *http.Request){

	// get the session

	session, err := getSession(res, req)

	if err != nil {
		http.Error(res, "Session not found...", http.StatusNotFound)
		return
	}

	isLoggedin := session.Values["IsLoggedIn"].(bool)
	email := session.Values["email"].(string)

	if !isLoggedin{
		http.Error(res,"Not Logged In!", 401 )
		return
	}
	
	fmt.Fprintf(res, "Welcome "+ email + " in Your profile! ");
}
