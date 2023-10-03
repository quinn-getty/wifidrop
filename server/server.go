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
	"github.com/quinn-getty/airdrop-go/controller_v2"
	"github.com/quinn-getty/airdrop-go/server/ws"
)

//go:embed dist/*
var FS embed.FS

//go:embed .chat/*
var ChatFS embed.FS

// 定义中间
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Print("【IP】", c.ClientIP())
		c.Next()
	}
}

func initApiV1(r *gin.RouterGroup, hub *ws.Hub) {
	r.GET("/uploads/:path", controller_v1.UploadsController)
	r.POST("/files", controller_v1.FilesController)
	r.GET("/qrcodes", controller_v1.QrcodeController)
	r.POST("/texts", controller_v1.TextController)
	r.GET("/addresses", controller_v1.AddressesController)
	r.GET("/ws", func(ctx *gin.Context) {
		ws.HttpController(ctx, hub)
	})
}
func initApiV2(r *gin.RouterGroup, hub *ws.Hub) {
	r.POST("/send", controller_v2.Send)
	r.GET("/history", controller_v2.History)
}

func Run(port int) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	hub := ws.NewHub()
	go hub.Run()
	staticFiles, _ := fs.Sub(FS, "dist")
	chatStaticFiles, _ := fs.Sub(ChatFS, ".chat")

	r.Use(MiddleWare())

	initApiV1(r.Group("/api/v1"), hub)
	initApiV2(r.Group("/api/v2"), hub)

	r.StaticFS("/static", http.FS(staticFiles))
	r.StaticFS("/chat", http.FS(chatStaticFiles))
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
			ctx.Status(http.StatusInternalServerError)
			// ctx.Redirect(http.StatusFound, "/index.html")
		}
	})
	log.Println("开发服务: ", fmt.Sprintf("http://127.0.0.1:%d/static", port))
	r.Run(fmt.Sprintf(":%d", port))
}
