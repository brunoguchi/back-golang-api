package routes

import (
	"github.com/brunoguchi/back-golang-api/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("api")
	{
		books := main.Group("books")
		{
			books.GET("/", controllers.GetBooks)
			books.GET("/:id", controllers.GetBookById)
		}
	}

	return router
}
