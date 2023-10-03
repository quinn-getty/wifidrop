package server

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/quinn-getty/airdrop-go/controller_v1"
	"github.com/quinn-getty/airdrop-go/server/ws"
)

//go:embed dist/*
var FS embed.FS

// 定义中间
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Print("【IP】", c.ClientIP())
		c.Next()
	}
}

func initApiV1(r *gin.RouterGroup, hub *ws.Hub) {
	r.GET("/api/v1/uploads/:path", controller_v1.UploadsController)
	r.POST("/api/v1/files", controller_v1.FilesController)
	r.GET("/api/v1/qrcodes", controller_v1.QrcodeController)
	r.POST("/api/v1/texts", controller_v1.TextController)
	r.GET("/api/v1/addresses", controller_v1.AddressesController)
	r.GET("/api/v1/ws", func(ctx *gin.Context) {
		ws.HttpController(ctx, hub)
	})
}

func Run(port int) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	hub := ws.NewHub()
	go hub.Run()
	staticFiles, _ := fs.Sub(FS, "dist")

	r.Use(MiddleWare())

	apiV1 := r.Group("/api/v1")
	initApiV1(apiV1, hub)

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
