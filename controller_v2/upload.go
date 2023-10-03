package controller_v2

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/quinn-getty/airdrop-go/utils"
)

func Uploads(c *gin.Context) {
	file, err := c.FormFile("raw")
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	filenameUUID := uuid.New().String()
	filename := filenameUUID + file.Filename

	uploadPath, _ := utils.GetUploadsDir()

	if err = c.SaveUploadedFile(file, filepath.Join(uploadPath, filename)); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	// 写库

	c.JSON(http.StatusOK, gin.H{
		"name": filename,
		"id":   filenameUUID,
	})

}
