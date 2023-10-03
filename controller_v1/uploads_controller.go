package controller_v1

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/quinn-getty/airdrop-go/utils"
)

func UploadsController(c *gin.Context) {
	uploadsPath, _ := utils.GetUploadsDir()

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
