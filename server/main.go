package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	
	"catbook.com/auth"
)

func main() {
	fmt.Println("Hello Air!")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/signin", func(c *gin.Context) {
		userCreds := auth.UserCredentials
		c.BindJSON(&userCreds)
	})

	r.POST("/register", func *gin.Context) {
		userCreds := auth.UserCredentials
		c.BindJSON(&userCreds)
	}

	r.Run()
}