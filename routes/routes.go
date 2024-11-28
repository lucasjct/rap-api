package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasjct/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/artist", controllers.ShowAllArtist)
	r.POST("/artist", controllers.CreateNewArtist)
	r.GET("/artist/:id", controllers.SearchArtistById)
	r.DELETE("/artist/:id", controllers.DeleteArtist)
	r.PATCH("/artist/:id", controllers.UpdateArtist)
	r.GET("/artist/name/:name", controllers.SearchArtistByName)
	r.NoRoute(controllers.RouteNotFound) // use when the return is not found
	r.GET("/ping", controllers.Ping)
	r.Run()

}
