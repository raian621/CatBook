package main

import (
	"fmt"
	"net/http"

	"catbook.com/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}

	r.Use(cors.New(corsConfig))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/signin", func(c *gin.Context) {
		userCreds := auth.UserCredentials{}
		c.BindJSON(&userCreds)
		fmt.Println(userCreds)
	})

	r.POST("/register", func(c *gin.Context) {
		userCreds := auth.UserCredentials{}
		c.BindJSON(&userCreds)
		fmt.Println(userCreds)
	})

	r.Run()
}
