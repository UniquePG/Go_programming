package main

import (
    "fmt"
    "net/http"		// this package has all the neccessory methods to handle user that req. from browser
)

func main() {

	//* This is the handler to handle the route it takes 2 parameters: 
		//! 	1. route 		2. handler func. that handle the route
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        fmt.Fprintf(w, "Hello, go lang server is up: %s\n", r.URL.Path)
    })


    http.ListenAndServe(":80", nil)		// it set up the server on the specific port
}