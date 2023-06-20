package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Karthika-Rajagopal/go-restful/controllers"
	"github.com/Karthika-Rajagopal/go-restful/middleware"
)

func InitUserRoutes(router *gin.RouterGroup) {
	user := router.Group("/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/profile", controllers.GetProfile)
		user.PUT("/profile", controllers.UpdateProfile)
	}
}
