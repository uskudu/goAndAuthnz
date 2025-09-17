package main

import (
	"authnz/initializers"
	"authnz/internal/db"
	"authnz/internal/handlers"
	"authnz/internal/middlewarre"
	"authnz/internal/userService"
	"fmt"

	_ "authnz/docs"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Authnz API
// @version 1.0
// @description signup and login with jwt
// @host localhost:3000
// @BasePath /
func main() {
	initializers.LoadEnvVariables()
	database, err := db.Connect()
	if err != nil {
		panic(err)
	}
	r := gin.Default()

	usrRepo := userService.NewUserRepository(database)
	usrService := userService.NewUserService(usrRepo)
	usrHandlers := handlers.NewUserHandler(usrService)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/signup", usrHandlers.Signup)
	r.POST("/login", usrHandlers.Login)
	r.GET("/validate", middlewarre.RequireAuth, usrHandlers.Validate)
	err = r.Run()
	if err != nil {
		fmt.Errorf("error while running gin: %v", err)
	}
}
