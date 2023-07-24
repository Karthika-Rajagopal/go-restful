package main

import (
	"log"
	//"os"

	"github.com/gin-gonic/gin"
	//"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/Karthika-Rajagopal/go-restful/controllers"
	"github.com/Karthika-Rajagopal/go-restful/cmd/docs"
	"github.com/Karthika-Rajagopal/go-restful/middlewares"
	"github.com/Karthika-Rajagopal/go-restful/models"
	"github.com/Karthika-Rajagopal/go-restful/config"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for a RESTful API built with Golang, Gin, and Gorm.
// @host localhost:8080
// @BasePath /
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.InitDB()

	defer db.Close()

	db.AutoMigrate(&models.User{}) // Perform AutoMigration

	r := gin.Default()
	r.Use(middlewares.ErrorHandler())

	// Swagger
	docs.SwaggerInfo.Title = "API Documentation"
	docs.SwaggerInfo.Description = "API endpoints documentation"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}

		user := v1.Group("/user")
		user.Use(middlewares.Authenticate()) // JWT authentication middleware
		{
			user.GET("/profile", controllers.GetProfile)
			user.PUT("/profile", controllers.UpdateProfile)
		}
	}

	r.Run(":8080")
}
