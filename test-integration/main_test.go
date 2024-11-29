package testintegration

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

var artistID float64 // Global variable to store the ID

func BaseUrl() string {
	URL := "http://localhost:8080"
	return URL
}

// create the instance to handle with http requests.

func HTTPexpectinstance(t *testing.T) *httpexpect.Expect {
	return httpexpect.Default(t, BaseUrl())
}

// create an artist and get the id value to storage
// on global variable.
func TestCreatArtist(t *testing.T) {
	e := HTTPexpectinstance(t)

	payload := map[string]interface{}{
		"name":     "Racionais MC's",
		"album":    "Sobrevivendo no Inferno",
		"released": "1997-12-20",
		"country":  "Brazil",
	}

	test := e.POST("/artist").
		WithJSON(payload).
		Expect().
		Status(http.StatusCreated)

	// storage the id vaalue to global variable.
	// it will be user as a parameter on the next requests.
	artistID = test.JSON().Object().Value("ID").Number().Raw()
	fmt.Println(artistID)
}

// list the artist by ID.
func TestListArtistByID(t *testing.T) {
	e := HTTPexpectinstance(t)
	content := "Racionais"
	test := e.GET("/artist").
		WithQuery("id", artistID).
		Expect().
		Status(http.StatusOK)

	test.Body().Contains(content)
}

// list all artist created.
func TestListAllArtists(t *testing.T) {
	e := HTTPexpectinstance(t)
	content := "Racionais"
	test := e.GET("/artist").
		WithQuery("page", 1).
		Expect().
		Status(http.StatusOK)

	test.Body().Contains(content)
}

// list artist by name.
func TestListArtistsByName(t *testing.T) {
	e := HTTPexpectinstance(t)
	content := "Racionais"
	test := e.GET("/artist").
		WithQuery("id", content).
		Expect().
		Status(http.StatusOK)

	test.Body().Contains(content)
}

// update artist data.
func TestUpdateArtists(t *testing.T) {
	e := HTTPexpectinstance(t)
	content := "Brasil"

	payload := map[string]interface{}{
		"country": content,
	}
	test := e.PATCH("/artist").
		WithQuery("id", artistID).
		WithJSON(payload).
		Expect().
		Status(http.StatusOK)

	test.Body().Contains(content)
}

// delete artist.
func TestDeleteArtist(t *testing.T) {
	e := HTTPexpectinstance(t)

	test := e.DELETE("/artist").
		WithQuery("id", artistID).
		Expect().
		Status(http.StatusOK)
	test.Body().Contains("The artist has been deleted.")
}
