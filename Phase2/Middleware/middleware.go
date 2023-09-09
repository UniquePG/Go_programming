package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Authentication(originalHandler http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// authenticate token
		if !(r.Header.Get("token") == "auth") {

			fmt.Fprintln(w, "Token is not valid. please login")
			return
		}

		// send the original handler -> it is like Next() func in node.js
		originalHandler.ServeHTTP(w, r)

	}

}

// middleware to check form property

func Check(f http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var user User

		// now get data from the body and decode it into our user struct
		json.NewDecoder(r.Body).Decode(&user) // pass user(struct) must be reference type

		if user.FirstName == "" || user.LastName == "" || user.Email == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Please provide all details...")
			return
		}

		f.ServeHTTP(w, r)
	}
}
