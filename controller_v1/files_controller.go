package controller_v1

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/quinn-getty/airdrop-go/utils"
)

func FilesController(c *gin.Context) {
	file, err := c.FormFile("raw")
	log.Print("Filename", file.Filename)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	filename := uuid.New().String() + file.Filename

	uploadPath, _ := utils.GetUploadsDir()

	if err = c.SaveUploadedFile(file, filepath.Join(uploadPath, filename)); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"url": "/api/v1/uploads/" + filename,
	})

}
