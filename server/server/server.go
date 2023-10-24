package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"catbook.com/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start(database *sql.DB) {
	r := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	r.Use(cors.New(corsConfig))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/signin", func(c *gin.Context) {
		var userCreds *auth.UserCredentials
		c.BindJSON(userCreds)
		fmt.Println(userCreds)
	})

	r.POST("/register", func(c *gin.Context) {
		var userRegInfo *auth.UserRegistrationInfo
		c.BindJSON(userRegInfo)
		fmt.Println(userRegInfo)
	})

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if host == "" || port == "" {
		r.Run()
	} else {
		r.Run(fmt.Sprintf("%s:%s", host, port))
	}
}
