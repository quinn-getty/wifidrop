package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/quinn-getty/airdrop-go/chrome"
)

//go:embed frontend/dist/*
var FS embed.FS

func main() {
	go func() {
		gin.SetMode(gin.ReleaseMode)
		r := gin.Default()

		staticFiles, _ := fs.Sub(FS, "frontend/dist")
		r.StaticFS("/static", http.FS(staticFiles))
		r.NoRoute(func(ctx *gin.Context) {
			path := ctx.Request.URL.Path
			if strings.HasPrefix(path, "/static/") {
				reader, err := staticFiles.Open("index.html")
				if err != nil {
					log.Fatal("err")
				}
				defer reader.Close()
				stat, err := reader.Stat()
				if err != nil {
					log.Fatal("err")
				}
				ctx.DataFromReader(http.StatusOK, stat.Size(), "text/html", reader, nil)
			} else {
				ctx.Redirect(http.StatusFound, "/index.html")
			}
		})
		r.Run(":8080")
	}()

	cmd := chrome.Open("http://127.0.0.1:8080/static")

	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)

	select {
	case <-chSignal:
		cmd.Process.Kill()
	}
}
