package main

import (
	"fmt"
	"net/http"
	"log"
)


func HelloHandler(w http.ResponseWriter, r *http.Request) {

		// handle the error if path route not match
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

		// check method
	if r.Method != "GET" {
		http.Error(w,"method is wrong", http.StatusNotFound ) // internal server error
		return
	}
	
	fmt.Fprintf(w, "This is hello from server")

}

func FormHandler(w http.ResponseWriter, r *http.Request) {

		// handle the error if path route not match
	if r.URL.Path != "/form"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

		// check method
	if r.Method != "POST" {
		http.Error(w,"method is wrong", http.StatusNotFound ) // internal server error
		return
	}

	// now parse the form and handle errors
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "error in parseForm func: %v \n", err)
		return
	}

	fmt.Fprintf(w, "Post request successful sent \n")

	// now fetch form values and response on them
	name := r.FormValue("name")
	number := r.FormValue("mobile")
	address := r.FormValue("address")

		// response on form values and sent to the browser as a response
	fmt.Fprintf(w, "Name: %s \n", name)	
	fmt.Fprintf(w, "Number: %v \n", number)
	fmt.Fprintf(w, "Address: %s \n", address)


}

func main() {

	// to serve static files like- html,css,js,img....
	fs := http.FileServer(http.Dir("./static")) //* dir is the path of folder where we have our static file

	http.Handle("/", fs)	// handle all staic files inside static folder

	// now handle specific routes
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/form", FormHandler)

	fmt.Printf("server started at port 8080.....")

	if err := http.ListenAndServe(":8080", nil); err != nil {

		log.Fatal(err)
	}

	

}

