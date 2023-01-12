package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type error struct {
	Error string `json: "error"`
}

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, error{"bad_request"})
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, albums)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, item := range albums {
		if item.ID == id {
			c.IndentedJSON(http.StatusOK, item)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"massage": "album not found"})
}

func removeAlbumById(c *gin.Context) {
	id := c.Param("id")

	for i, item := range albums {
		if item.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusNoContent, item)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"massage": "album not found"})

}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)
	router.DELETE("/albums/:id", removeAlbumById)
	router.Run("localhost:8082")
}
