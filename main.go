package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net"
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
	"github.com/skip2/go-qrcode"
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
		r.GET("/uploads/:path", UploadsController)
		r.POST("/api/v1/files", FilesController)
		r.GET("/api/v1/qrcodes", QrcodeController)
		r.POST("/api/v1/texts", TextController)
		r.GET("/api/v1/addresses", AddressesController)
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

func FilesController(c *gin.Context) {
	file, err := c.FormFile("raw")
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	dir, _ := GetUploadsDir()
	fullPath := filepath.Join(dir, filepath.Join(dir, filepath.Ext(file.Filename)))
	if err = c.SaveUploadedFile(file, fullPath); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"url": "/" + fullPath,
	})

}

func QrcodeController(c *gin.Context) {
	content := c.Query("content")
	if content == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	log.Print("content:", content)
	png, err := qrcode.Encode(content, qrcode.Medium, 256)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.Data(http.StatusOK, "image/png", png)
}

func GetUploadsDir() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	dir := filepath.Dir(exe)
	uploads := filepath.Join(dir, "uploads")
	log.Print("uploads----------:", uploads)
	return uploads, nil
}

func UploadsController(c *gin.Context) {
	uploadsPath, _ := GetUploadsDir()
	if path := c.Param("path"); path != "" {
		target := filepath.Join(uploadsPath, path)
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; filename="+path)
		c.File(target)
	} else {
		c.Status(http.StatusNotFound)
	}
}
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
