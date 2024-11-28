package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasjct/database"
	"github.com/lucasjct/models"
)

// list all artist
//
//	@Summary	List all artist
//	@Success	200
//	@Router		/artist [get]
//	@Tags		Artists
func ShowAllArtist(c *gin.Context) {
	var artist []models.Artist
	database.DB.Find(&artist)
	c.JSON(200, artist)

}

// Create a new artist
//
//	@Summary		Create a new artist
//	@Description	This endpoint allows the creation of a new artist by providing artist details in JSON format.
//	@Tags			Artists
//	@Accept			json
//	@Produce		json
//	@Param			artist	body		models.Artist	true	"Artist details"
//	@Success		201		{object}	models.Artist	"Artist Created Successfully"
//	@Router			/artist [post]
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
//
//	@Summary	List an artist by ID
//	@Param		id	path	integer	true	"Artist ID"
//	@Tags		Artists
//	@Success	200
//	@Router		/artist/{id}    [get]s
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
//
//	@Summary	Search artist by name
//	@Param		name	path		string			true	"Artist Name"
//	@Param	    artist	body		models.Artist	true	"Artist details"
//	@Success	200		{object}	models.Artist	"Artist object"
//	@Tags		Artists
//	@Router		/artist/{name}   [get]
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

// delete a specific artist

// @Param		id	path	integer	true	"Artist ID"
// @Summary	Delete a specific artist
// @Tags		Artists
// @Success	200
// @Router		/artist/{id}   [delete]
func DeleteArtist(c *gin.Context) {
	var artist models.Artist
	id := c.Params.ByName("id")
	database.DB.Delete(&artist, id)
	c.JSON(http.StatusOK, gin.H{"Delete": "The artist has been deleted."})
}

// update a specific artist
//
//	@Summary	Update a specific artist
//	@Param		id	path	integer	true	"Artist ID"
//	@Tags		Artists
//	@Produce	json
//	@Success	200	{object}	models.Artist	"Artist Updated Successfully"
//	@Router		/artist/{id}   [patch]
//	@Accept		json
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
//
//	@Summary	Test API connection
//	@Schemes
//	@Description	do ping
//	@Tags			Test API connection
//	@Accept			json
//	@Produce		json
//	@Router			/ping [get]
//	@Success		200	{string}	pong
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong.",
	})

}

// return error to page not found - status code 404
func RouteNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
