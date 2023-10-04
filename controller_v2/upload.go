package controller_v2

import (
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/quinn-getty/airdrop-go/file"
	"github.com/quinn-getty/airdrop-go/utils"
)

func Uploads(c *gin.Context) {
	data, err := c.FormFile("raw")
	fileType := c.PostForm("type")
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	filenameUUID := uuid.New().String()
	filename := filenameUUID + "_" + data.Filename
	// data.Size

	uploadPath, _ := utils.GetUploadsDir()

	if err = c.SaveUploadedFile(data, filepath.Join(uploadPath, filename)); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	if err := file.WhiteHistory(file.HistoryTiem{
		Type:    fileType,
		Content: filename,
		Time:    time.Now().Unix(),
		IP:      c.ClientIP(),
	}); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name": filename,
		"id":   filenameUUID,
	})

}
