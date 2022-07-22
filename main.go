package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     int64   `json: "id"`
	Title  string  `json: "title"`
	Artist string  `json: "artist"`
	Price  float64 `json: "price"`
}

var albums = []Album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumID)

	router.POST("/albums", postAlbum)

	router.Run("localhost:8080")
}

/**
*	gin.Context is the most important part of Gin.
*	It carries request details, validates and serializes JSON, and more.
 */

func getAlbums(c *gin.Context) {
	/**
	*	Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.
	*	Note that you can replace Context.IndentedJSON with a call to Context.JSON to send more compact JSON.
	*	In practice, the indented form is much easier to work with when debugging and the size difference is usually small.
	 */
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumID(c *gin.Context) {
	id := c.Param("id")

	// strconv.ParseInt(s string, base int, bitSize int)
	// Convert string to int
	// base set to 10
	i, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "ID must be numeric")
		return
	}

	for _, a := range albums {
		if a.ID == i {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
