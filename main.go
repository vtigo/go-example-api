package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// We will be storing data in a slice for now. That means that data will be stored in memory and lost every time the program stops

// Struct tags such as json:"artist" specify what a field’s name should be when the struct’s contents are serialized into JSON. Without them, the JSON would use the struct’s capitalized field names – a style not as common in JSON
type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  int64  `json:"price"`
}

// Our albums slice with some starting data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 5699},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 1799},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 3999},
}

func main() {
	// Gin initializatin and route definitions
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// CRUD functions
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	
	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

