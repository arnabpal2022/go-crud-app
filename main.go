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