package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var key = []byte("This-is-secret-key") // here we used sequence of bytes to more secure our secret key

var store = sessions.NewCookieStore(key) // this is our secret key to secure our session management

// make our handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to our website")

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	// first we get the session from the store
	session, err := store.Get(r, "login-session") //*we get the session -> it gives us session obj OR err

	// handler err
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // give the error if found
		return
	}

	// Now set the values in our session
	session.Values["username"] = "Pavan" // username is a string
	session.Values["IsLoggedIn"] = true  // it is bool (authentication)

	// now save the session
	err = session.Save(r, w) // it returns error (if occure)

	// check the error while saving the session
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError) // give the error if found
		return
	}

	// print if session set successfully
	w.Write([]byte("Session set successfully! \n"))
	fmt.Fprintf(w, "Session set Successfully... You are logged in")

}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {

	// first get the session using same name as we used while setting the session
	session, err := store.Get(r, "login-session")

	//handle error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // give the error if found
		return
	}

	// Now retrive session values for authentication
	username, ok := session.Values["username"].(string) // specify value type(string)

	// if username not found in sesion then it gives false in 'ok' variable
	if !ok {
		http.Error(w, "Username not found in session", http.StatusInternalServerError)
		return
	}

	// now retrive isLoggedin value for authentication of login
	isLoggedin, _ok := session.Values["IsLoggedIn"].(bool)

	// if username not found in sesion then it gives false in 'ok' variable
	if !_ok {
		http.Error(w, "You are not logged in, Please login! \n", http.StatusInternalServerError)
		return
	}

	// Now use our session value as we want
	if isLoggedin {

		w.Write([]byte("Welcome " + username + ". You are logged in \n"))

	} else {
		w.Write([]byte("Welcome guest. You are not logged in \n"))
	}

}
func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	// first get the session using same name as we used while setting the session
	session, err := store.Get(r, "login-session")

	//handle error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // give the error if found
		return
	}

//! addition checking -> we will not going to allow log out if user is not logged in
	// get the isloggedin value from session and check it
	isLoggedin, _ := session.Values["IsLoggedIn"].(bool)

	if !isLoggedin {
		w.Write([]byte("You are not loggedIn so You can not logout now. First have to logged in before logout \n"))

		return
	}

	// now invoke the session value
	// Set isLoggedin as false for logout
	session.Values["IsLoggedIn"] = false

	session.Save(r, w)

	fmt.Fprintf(w, "You are successfully logged out \n")
}
