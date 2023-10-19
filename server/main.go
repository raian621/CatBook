package main

import (
	"fmt"
	"net/http"
	"os"

	"catbook.com/auth"
	"catbook.com/db"
	"catbook.com/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	util.LoadEnvVars(".env")

	dbParams := &db.DatabaseParams{
		Provider: os.Getenv("DB_PROV"),
		Hostname: os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSL"),
	}

	util.DeleteEnvVars()

	database, err := db.ConnectToDB(dbParams)
	dbParams = nil
	if err != nil {
		fmt.Println("Database offline!")
		panic(err)
	}
	err = db.CreateTables(database)
	if err != nil {
		panic(err)
	}

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
