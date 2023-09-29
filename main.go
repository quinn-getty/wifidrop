package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/quinn-getty/airdrop-go/chrome"
)

func main() {
	go func() {
		gin.SetMode(gin.ReleaseMode)
		r := gin.Default()

		r.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "<h1>hi</h1>")
		})
		r.Run(":8080")
	}()

	cmd := chrome.Open("http://127.0.0.1:8080")
	log.Println(cmd)

	chSignal := make(chan os.Signal, 1)
	select {
	case <-chSignal:
		signal.Notify(chSignal, os.Interrupt)
		cmd.Process.Kill()
	}
}
