package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// We're not using a Database in Case of This, So we'll be going through making a Struct 
// that will be encoded in json while making API Requests

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Name     string    `json:"name"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

// The function `getallmovies` sets the response content type to JSON and encodes a list of movies to
// the response writer.
func getallmovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// The function `deletemovie` deletes a movie from a list of movies based on the provided ID.
func deletemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Vars returns the route variables for the current request, if any.
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // Used to Delete an Item
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// The function retrieves a specific movie by its ID and returns it as JSON in a HTTP response.
func getmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Vars returns the route variables for the current request, if any.
	// Showing only the "id"th item
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// The function `createmovie` decodes a JSON request body into a Movie struct, assigns a random ID to
// the movie, adds it to a slice of movies, and then encodes and returns the movie as JSON in the
// response.
func createmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000)) // Random ID Generation
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// The function `updatemovie` updates a movie in a list based on the provided ID in a Go web
// application.
func updatemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Logic is Delete that record and Create a Record, Just the ID is Same.
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", ISBN: "100001", Name: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", ISBN: "100002", Name: "Movie Two", Director: &Director{FirstName: "John", LastName: "K"}})

	r.HandleFunc("/movies", getallmovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getmovie).Methods("GET")
	r.HandleFunc("/movies", createmovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updatemovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deletemovie).Methods("DELETE")

	fmt.Printf("Starting Server at Port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
