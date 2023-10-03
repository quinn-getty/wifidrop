package server

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/quinn-getty/airdrop-go/controller"
	"github.com/quinn-getty/airdrop-go/server/ws"
)

//go:embed frontend/dist/*
var FS embed.FS

func Run(port int) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	hub := ws.NewHub()
	go hub.Run()

	staticFiles, _ := fs.Sub(FS, "frontend/dist")
	r.GET("/uploads/:path", controller.UploadsController)
	r.POST("/api/v1/files", controller.FilesController)
	r.GET("/api/v1/qrcodes", controller.QrcodeController)
	r.POST("/api/v1/texts", controller.TextController)
	r.GET("/api/v1/addresses", controller.AddressesController)
	r.GET("/ws", func(ctx *gin.Context) {
		ws.HttpController(ctx, hub)
	})
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
	log.Println("开发服务: ", fmt.Sprintf("http://127.0.0.1:%d/static", port))
	r.Run(fmt.Sprintf(":%d", port))
}
