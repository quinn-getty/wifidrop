package controller

import (
	"log"
	"os"
	"path/filepath"
)

func GetUploadsDir() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(exe)
	if err != nil {
		log.Fatal(err)
	}
	uploads := filepath.Join(dir, "uploads")
	err = os.MkdirAll(uploads, os.ModePerm)
	if err != nil {
		return "", err
	}
	return uploads, nil
}
