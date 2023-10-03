package controller_v2

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SendReq struct {
	Type     string `json:"type"`
	Content  string `json:"content"`
	FileType string `json:"fileType"`
}

func Send(c *gin.Context) {
	req := SendReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// 存库
	// 发送socket
}
