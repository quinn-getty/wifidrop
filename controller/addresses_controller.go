package controller

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddressesController(c *gin.Context) {
	addresses, _ := net.InterfaceAddrs()
	var result []string
	for _, address := range addresses {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				result = append(result, ipnet.IP.String())
			}
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"addresses": result,
	})
}
