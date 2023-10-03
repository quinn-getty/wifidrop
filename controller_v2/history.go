package controller_v2

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quinn-getty/airdrop-go/file"
)

func History(c *gin.Context) {
	list, err := file.GetHistory()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"list": "[]",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"list": list,
	})
}
