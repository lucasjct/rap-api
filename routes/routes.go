package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasjct/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/artist", controllers.ShowAllArtist)
	r.GET("/artist/:id", controllers.SearchArtistById)
	r.GET("/artist/name/:name", controllers.SearchArtistByName)
	r.GET("/ping", controllers.Ping)
	r.DELETE("/artist/:id", controllers.DeleteArtist)
	r.POST("/artist", controllers.CreateNewArtist)
	r.PATCH("/artist/:id", controllers.UpdateArtist)
	r.NoRoute(controllers.RouteNotFound) // use it when the return will be not found
	r.Run()

}
