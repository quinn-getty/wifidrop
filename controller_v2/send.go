package controller_v2

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/quinn-getty/airdrop-go/file"
)

type SendReq struct {
	Content string `json:"content"`
}

func Send(c *gin.Context) {
	req := SendReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := file.WhiteHistory(file.HistoryTiem{
		Type:    "string",
		Content: req.Content,
		Time:    time.Now().Unix(),
		IP:      c.ClientIP(),
	}); err != nil {
		log.Print(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)

}
