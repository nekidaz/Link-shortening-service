package main

import (
	"github.com/gin-gonic/gin"
	"jwt-auth/controllers"
	"jwt-auth/initializers"
	"jwt-auth/middleware"
	"net/http"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "404 Not Found"})
	})

	//authorization
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)

	go router.GET("/:shorten", controllers.RedirectUrl) //routing to redirect links by id

	//router group  with auth
	auth := router.Group("/")
	auth.Use(middleware.RequireAuth) //middleware to auth jwtTokens

	// routings where authorization is needed
	{
		auth.GET("/home", controllers.Home)
		auth.POST("/cut", controllers.CutUrl)

	}
	router.Run()
}
