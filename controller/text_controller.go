package controller

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TextReq struct {
	Raw string
}

func TextController(c *gin.Context) {
	var req = TextReq{}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uploadPath, err := GetUploadsDir()
	if err != nil {
		log.Fatal(err)
	}

	filename := uuid.New().String() + ".txt"

	err = os.WriteFile(filepath.Join(uploadPath, filename), []byte(req.Raw), 0644)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"url": "/api/v1/uploads/" + filename})
}
