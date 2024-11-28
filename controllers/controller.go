package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasjct/database"
	"github.com/lucasjct/models"
)

// list all artist
func ShowAllArtist(c *gin.Context) {
	var artist []models.Artist
	database.DB.Find(&artist)
	c.JSON(200, artist)

}

// create a new artist
func CreateNewArtist(c *gin.Context) {
	var artist models.Artist
	if err := c.ShouldBindJSON(&artist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidateArtist(&artist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&artist)
	c.JSON(http.StatusCreated, artist)
}

// search by specifc ID related to some artist
func SearchArtistById(c *gin.Context) {
	var artist models.Artist
	id := c.Params.ByName("id")
	database.DB.First(&artist, id)

	if artist.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "artist not found"})
		return

	}

	c.JSON(http.StatusOK, artist)
}

// search artist by name
func SearchArtistByName(c *gin.Context) {
	var artist models.Artist
	name := c.Param("name")
	database.DB.Where(&models.Artist{Name: name}).First(&artist)

	if artist.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "name not found"})
		return
	}
	c.JSON(http.StatusOK, artist)
}

// delte a specific artist
func DeleteArtist(c *gin.Context) {
	var artist models.Artist
	id := c.Params.ByName("id")
	database.DB.Delete(&artist, id)
	c.JSON(http.StatusOK, gin.H{"Delete": "The artist has been deleted."})

}

// update a specific artist
func UpdateArtist(c *gin.Context) {
	var artist models.Artist
	id := c.Params.ByName("id")
	database.DB.First(&artist, id)

	if err := c.ShouldBindJSON(&artist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidateArtist(&artist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&artist).UpdateColumns(artist)
	c.JSON(http.StatusOK, artist)

}

// test api response
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong.",
	})

}

// return error to page not found - status code 404
func RouteNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
