package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// fileServer := http.FileServer(http.Dir("/static")) //get Static Files
	http.Handle("/", http.FileServer(http.Dir("./static"))) // Serving Static Files : index.html
	http.HandleFunc("/hello", handleHello)                  //calling function when this route is called
	http.HandleFunc("/form", handleFunc)
	fmt.Println("Server Starting at 1552")
	err := http.ListenAndServe(":1552", nil) // Starting Server at 1552 and assigning nil to error if erverything is ook
	if err != nil {
		log.Fatal(err) //Log if any err
	}
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	// check for Parsing error
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm( error) err:%v", err)
		return
	}
	fmt.Println("Form data Recieved Succesfully")
	name := r.FormValue("name")
	phoneNumber := r.FormValue("phone")

	fmt.Fprintf(w, "Name : %v", name)
	fmt.Fprintf(w, "Phone Number : %v", phoneNumber)
}

func handleHello(w http.ResponseWriter, r *http.Request) {

	//Checking if Path is not correct
	if r.URL.Path != "/hello" {
		http.Error(w, "404! Path Not Found", http.StatusNotFound)
		return
	}
	//Checking if Method is not correct
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusNotAcceptable)
		return
	}
	fmt.Println(w, "Hello , Function ran Succesfully", http.StatusOK)
	fmt.Fprintf(w, "Hello")

}
