package main

import (
	"github.com/lucasjct/database"
	"github.com/lucasjct/routes"
)

func main() {
	database.DataBaseConnect()
	routes.HandleRequest()
}
