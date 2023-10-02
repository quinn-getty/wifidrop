package controller

import (
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
)

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
