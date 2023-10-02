package controller

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func FilesController(c *gin.Context) {
	file, err := c.FormFile("raw")
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	dir, _ := GetUploadsDir()
	fullPath := filepath.Join(dir, filepath.Join(dir, filepath.Ext(file.Filename)))
	if err = c.SaveUploadedFile(file, fullPath); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"url": "/" + fullPath,
	})

}
