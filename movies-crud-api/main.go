package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Creating Map for storing and Sending Data
type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"` //* is a pointer to struct Type Director
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// List of type Movie
var movies []Movie

func main() {
	r := mux.NewRouter() //importing Router from mux

	movies = append(movies, Movie{Id: "1", Isbn: "2002", Title: "Gangs", Director: &Director{Firstname: "Rounak", Lastname: "Jha"}})
	movies = append(movies, Movie{Id: "2", Isbn: "2052", Title: "Tabahi", Director: &Director{Firstname: "Hansu", Lastname: "Cemal"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	//Starting Server
	fmt.Printf("Starting server at http://localhost:1552")
	log.Fatal(http.ListenAndServe(":1552", r))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	//Method is GET
	//setting headers
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies) // Converting Struct type to json for sending full list of movies
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	//Method is GET
	//setting headers
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item) //send found movie
			return
		} else {
			http.Error(w, "No Matching ENtry", http.StatusNotFound)
		}
	}
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	//Method is POST
	//setting headers
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie) // Decodes Request body and saves in variable movie
	movie.Id = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)

}
func updateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // getting params from the request
	for index, item := range movies {
		if item.Id == params["id"] { //seaching for entry of that specific ID
			movies = append(movies[:index], movies[index+1:]...) //Appending all movies except for found ig
			var movie Movie
			movie.Id = params["id"]
			movies := append(movies, movie)
			json.NewEncoder(w).Encode(movies)
		}
	}
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	//Method is DELETE
	//setting headers
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // getting params from the request
	for index, item := range movies {
		if item.Id == params["id"] { //seaching for entry of that specific ID
			movies = append(movies[:index], movies[index+1:]...) //Appending all movies except for found ig
			break
		}
	}
	json.NewEncoder(w).Encode(movies)

}
