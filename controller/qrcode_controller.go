package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func QrcodeController(c *gin.Context) {
	content := c.Query("content")
	if content == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	log.Print("content:", content)
	png, err := qrcode.Encode(content, qrcode.Medium, 256)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.Data(http.StatusOK, "image/png", png)
}
