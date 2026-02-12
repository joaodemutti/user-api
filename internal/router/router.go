package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaodemutti/user-api/internal/database"
	"github.com/joaodemutti/user-api/internal/middleware"
	"github.com/joaodemutti/user-api/internal/user"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	db := database.Connect()

	// Auto migrate table
	db.AutoMigrate(&user.User{})

	repo := user.NewRepository(db)
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	r.GET("/users", middleware.AuthMiddleware(), handler.GetUsers)
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	return r
}
