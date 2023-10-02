package controller

import (
	"log"
	"os"
	"path/filepath"
)

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
