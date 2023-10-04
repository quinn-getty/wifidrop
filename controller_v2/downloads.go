package controller_v2

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/quinn-getty/airdrop-go/utils"
)

func DownLoads(c *gin.Context) {
	uploadsPath, _ := utils.GetUploadsDir()

	if path := c.Param("path"); path != "" {
		target := filepath.Join(uploadsPath, path)
		log.Println(target)
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; filename="+path)
		c.File(target)
	} else {
		c.Status(http.StatusNotFound)
	}
}
