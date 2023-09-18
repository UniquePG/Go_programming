package main

import (
	"fmt"
	"log"
	"net/http"

	routes "github.com/uniquepg/gowithmongo/Routes"
)

func main() {

	// get instance of router
	r := routes.Routes()

	log.Fatal(http.ListenAndServe(":6000", r))
	fmt.Println("Server is started on Port 6000...")

}
