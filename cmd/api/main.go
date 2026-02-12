// @title User API
// @version 1.0
// @description Simple User API with Gin + PostgreSQL
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"log"

	"github.com/joaodemutti/user-api/internal/router"

	"github.com/joho/godotenv"

	_ "github.com/joaodemutti/user-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	r := router.SetupRouter()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
