package main

import (
	"coffee-delivery/database"
	"coffee-delivery/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()
	routes.InitRoutes(r)

	r.Run(":8080")
}
