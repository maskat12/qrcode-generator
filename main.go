package main

import (
	"qrcode-generator/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/generate", controllers.Generate)
	r.Static("/storage", "./storage")

	r.Run() // listen and serve on 0.0.0.0:8080
}
