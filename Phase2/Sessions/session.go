package main

import (
	"fmt"
	"net/http"
)

func main() {

	// make routes
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/profile", ProfileHandler)
	http.HandleFunc("/logout", LogoutHandler)

	fmt.Println("Server started...")

	http.ListenAndServe(":3002", nil)
}
