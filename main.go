package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/quinn-getty/airdrop-go/chrome"
	"github.com/quinn-getty/airdrop-go/utils"
)

//go:embed frontend/dist/*
var FS embed.FS

func main() {
	port, err := utils.GetFreePort()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		gin.SetMode(gin.ReleaseMode)
		r := gin.Default()

		staticFiles, _ := fs.Sub(FS, "frontend/dist")
		r.POST("/api/v1/texts", TextController)
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
		r.Run(fmt.Sprintf(":%d", port))
	}()

	log.Println(fmt.Sprintf("http://127.0.0.1:%d/static", port))

	cmd := chrome.Open(fmt.Sprintf("http://127.0.0.1:%d/static", port))

	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)

	select {
	case <-chSignal:
		cmd.Process.Kill()
	}
}

type TextReq struct {
	Raw string
}

func TextController(c *gin.Context) {
	var req = TextReq{}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(exe)
	if err != nil {
		log.Fatal(err)
	}
	filename := uuid.New().String()
	uploads := filepath.Join(dir, "uploads")
	err = os.MkdirAll(uploads, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	fullpath := path.Join("uploads", filename+".txt")
	err = os.WriteFile(filepath.Join(dir, fullpath), []byte(req.Raw), 0644)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"url": "/" + fullpath})
}
