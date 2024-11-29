package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucasjct/database"
	"github.com/lucasjct/models"
)

// list all artist
//
//	@Summary	List all artist
//	@Success	200
//	@Router		/artist  [get]
//	@Tags		Artists
func ShowAllArtist(c *gin.Context) {

	// define query parameter
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)

	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid page parameter",
		})
		return
	}

	const pageSize = 5
	offset := (page - 1) * pageSize
	var totalRows int64
	var artist []models.Artist

	// data limit per page
	database.DB.Limit(5).Offset(offset).Find(&artist)

	// count total page
	database.DB.Model(&models.Artist{}).Count(&totalRows)

	totalPages := (totalRows + int64(pageSize) - 1) / int64(pageSize)
	// show total page on headers
	c.Header("Total-Pages", fmt.Sprintf("%d", totalPages))
	// show page on headers
	c.Header("Current-Page", pageStr)
	c.JSON(http.StatusOK, artist)
}

// Create a new artist
//
//	@Summary	Create a new artist
//	@Tags		Artists
//	@Accept		json
//	@Produce	json
//	@Param		artist	body		models.Artist	true	"Artist details"
//	@Param		page	query		int				false	"Page number"	default(1)	example(1)
//	@Success	201		{object}	models.Artist	"Artist Created Successfully"
//	@Router		/artist [post]
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
//	@Param		id	query	integer	true	"ID of the artist"	example("25")
//	@Tags		Artists
//	@Success	200
//	@Router		/artist/specific    [get]
func SearchArtistById(c *gin.Context) {

	id := c.DefaultQuery("id", "")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID query parameter is required",
		})
		return

	}
	var artist models.Artist

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
//	@Param		name	query	string	true	"Name of the artist"	example("Emicida")
//	@Success	200
//	@Tags		Artists
//	@Router		/artist/name    [get]
func SearchArtistByName(c *gin.Context) {
	var artist models.Artist
	name := c.DefaultQuery("name", "")

	if name == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "name not found",
		})
	}

	database.DB.Where(&models.Artist{Name: name}).First(&artist)
	c.JSON(http.StatusOK, artist)
}

// delete a specific artist

// @Param		id	query	integer	true	"ID of the artist"	example("25")
// @Summary	Delete a specific artist
// @Tags		Artists
// @Success	200
// @Router		/artist   [delete]
func DeleteArtist(c *gin.Context) {
	var artist models.Artist
	id := c.DefaultQuery("id", "")
	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "id not found",
		})
	}

	database.DB.Delete(&artist, id)
	c.JSON(http.StatusOK, gin.H{"Delete": "The artist has been deleted."})
}

// Update a specific artist
//
//	@Summary	Update a specific artist
//	@Tags		Artists
//	@Accept		json
//	@Produce	json
//	@Param		id		query		integer			true	"ID of the artist"	example("25")
//	@Param		artist	body		models.Artist	true	"Artist details"
//	@Success	200		{object}	models.Artist	"Artist updated successfully"
//	@Router		/artist   [patch]
func UpdateArtist(c *gin.Context) {
	var artist models.Artist
	id := c.DefaultQuery("id", "")

	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "id not found",
		})
	}
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
