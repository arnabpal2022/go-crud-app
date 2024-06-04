package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
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

// The function `getallmovies` in Go returns all movies in JSON format using the Gin framework.
func getallmovies(c *gin.Context){
	c.JSON(http.StatusOK, movies)
}

// The `deletemovie` function in Go deletes a movie from a collection based on its ID and returns the
// updated collection.
func deletemovie(c *gin.Context) {
	id := c.Param("id") // Param returns the value of the URL param.
	for index, item := range movies {
		if item.ID == id {
			movies = append(movies[:index], movies[index+1:]...) // Delete Logic
			break
		}
	}
	c.JSON(http.StatusOK, movies)
}

// The function `getmovie` retrieves a movie by its ID and returns it as JSON, or a "movie not found"
// error if the movie is not found.
func getmovie(c *gin.Context) {
    id := c.Param("id")
    for _, item := range movies {
        if item.ID == id {
            c.JSON(http.StatusOK, item)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "movie not found"})
}

// The function `createmovie` in Go creates a new movie object from JSON input, assigns a random ID,
// and adds it to a list of movies.
func createmovie(c *gin.Context) {
	var movie Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	movie.ID = strconv.Itoa(rand.Intn(10000000)) // Random ID Generation
	movies = append(movies, movie)
	c.JSON(http.StatusCreated, movie)
}

// The function `updatemovie` updates a movie in a collection based on the provided ID.
func updatemovie(c *gin.Context) {
    id := c.Param("id") // Param returns the value of the URL param.
    var updatedMovie Movie
    if err := c.ShouldBindJSON(&updatedMovie); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    for index, item := range movies {
        if item.ID == id {
            movies[index] = updatedMovie
            movies[index].ID = id // Preserve original ID
            c.JSON(http.StatusOK, movies[index])
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "movie not found"})
}

func main() {
	r := gin.Default()

	movies = append(movies, Movie{ID: "1", ISBN: "100001", Name: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", ISBN: "100002", Name: "Movie Two", Director: &Director{FirstName: "John", LastName: "K"}})

	r.GET("/movies", getallmovies)
	r.GET("/movies/:id", getmovie)
	r.POST("/movies", createmovie)
	r.PUT("/movies/:id", updatemovie)
	r.DELETE("/movies/:id", deletemovie)

	fmt.Printf("Starting Server at Port 8000\n")
	log.Fatal(r.Run(":8080"))
}
