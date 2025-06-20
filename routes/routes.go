package routes

import (
	"coffee-delivery/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.POST("/coffee-create", controllers.CreateCoffee)
	r.GET("/coffees", controllers.GetCoffees)
	r.DELETE("/coffees/:id", controllers.DeleteCoffee)
}
