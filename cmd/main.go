package main

import (
	"authnz/controllers"
	"authnz/initializers"
	"authnz/internal/db"
	"authnz/middlewarre"
	"fmt"

	_ "authnz/docs"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvVariables()
	db.ConectToDB()
	db.SyncDB()
}

// @title Authnz API
// @version 1.0
// @description signup and login with jwt
// @host localhost:3000
// @BasePath /
func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middlewarre.RequireAuth, controllers.Validate)
	err := r.Run()
	if err != nil {
		fmt.Errorf("error while running gin: %v", err)
	}
}
