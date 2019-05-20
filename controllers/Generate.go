package controllers

import (
	"fmt"
	"os"

	"github.com/gofrs/uuid"

	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

// Generate QR Codes URL
func Generate(c *gin.Context) {
	url := c.Query("url")
	storagePath := "storage/"
	FileName, _ := uuid.NewV4()
	fullpath := storagePath + FileName.String() + ".png"

	if err := qrcode.WriteFile(url, qrcode.Medium, 256, fullpath); err != nil {
		if err = os.Remove(fullpath); err != nil {
			fmt.Println(err)
		}
	}

	c.JSON(200, gin.H{
		"image_url": fullpath,
	})
}
