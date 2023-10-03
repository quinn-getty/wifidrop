package file

import (
	"log"
	"os"
	"path/filepath"
	"time"
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

	uploads := filepath.Join(dir, "uploads/"+time.Now().Format(time.DateOnly))
	err = os.MkdirAll(uploads, os.ModePerm)
	if err != nil {
		return "", err
	}
	return uploads, nil
}

func GetFileList() ([]string, error) {
	dir, _ := GetUploadsDir()
	return filepath.Glob(dir)
}

func RemoveLogFile(fileName string) error {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(exe)
	if err != nil {
		log.Fatal(err)
	}

	return os.Remove(filepath.Join(dir, fileName))
}
