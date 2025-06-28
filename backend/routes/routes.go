package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"backend/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	r.POST("/api/register", controllers.Register)
	r.POST("/api/login", controllers.Login)
	r.GET("/api/dashboard", controllers.Dashboard)

	return r
}
