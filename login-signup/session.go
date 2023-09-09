package main

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var key = []byte("secret-key")
var store = sessions.NewCookieStore(key)

func CreateSession(res http.ResponseWriter, req *http.Request, email string) (*sessions.Session, error) {

	session, err := store.Get(req, "login-session")

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError) // give the error if found
		return nil, err
	}

	session.Values["email"] = email
	session.Values["IsLoggedIn"] = true

	return session, nil

}


// get session
func getSession(res http.ResponseWriter, req *http.Request) (*sessions.Session, error) {

	session, err := store.Get(req, "login-session")

	if err != nil {
		http.Error(res, "Session not found", http.StatusInternalServerError)
		return session, err
	}

	return session, nil

}
