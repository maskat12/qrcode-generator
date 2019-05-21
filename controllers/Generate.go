package controllers

import (
	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

/**
* Generate QR Codes URL 
* @return JSON image_blob(Blob)	
*/
func Generate(c *gin.Context) {
	text := c.Query("text")
	png, _ := qrcode.Encode(text, qrcode.Medium, 256)

	c.JSON(200, gin.H{
		"image_blob": png,
	})
}
