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

type Movie struct {
	ID string `json:"id"`
	ISBN string `json:"isbn"`
	Name string `json:"name"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var movies []Movie;

func getallmovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json");
	json.NewEncoder(w).Encode(movies)
}

func deletemovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json");
	params := mux.Vars(r);
	for index, item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break;
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getmovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json");
	params := mux.Vars(r);
	for _, item := range movies {
		if item.ID == params["id"] { 
			json.NewEncoder(w).Encode(item);
			return;
		}
	}
}

func createmovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json");
	var movie Movie;
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updatemovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json");
	params := mux.Vars(r);

	for index, item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]... );
			var movie Movie;
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"];
			movies = append(movies, movie);
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", ISBN: "100001", Name: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", ISBN: "100002", Name: "Movie Two", Director: &Director{FirstName: "John", LastName: "K"}})

	r.HandleFunc("/movies", getallmovies).Methods("GET");
	r.HandleFunc("/movies/{id}", getmovie).Methods("GET");
	r.HandleFunc("/movies", createmovie).Methods("POST");
	r.HandleFunc("/movies/{id}", updatemovie).Methods("PUT");
	r.HandleFunc("/movies/{id}", deletemovie).Methods("DELETE");

	fmt.Printf("Starting Server at Port 8000\n");
	log.Fatal(http.ListenAndServe(":8000",r))
}