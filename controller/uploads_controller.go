package controller

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadsController(c *gin.Context) {
	uploadsPath, _ := GetUploadsDir()
	if path := c.Param("path"); path != "" {
		target := filepath.Join(uploadsPath, path)
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; filename="+path)
		c.File(target)
	} else {
		c.Status(http.StatusNotFound)
	}
}
