package main

import (
	"qrcode-generator/controllers"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/generate", controllers.Generate)
	r.Static("/storage", "./storage")

	r.Run(":"+os.Getenv("APP_PORT")) // listen and serve on 0.0.0.0:8080
}
