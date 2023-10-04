package utils

import (
	"log"
	"net"
	"os"
	"path/filepath"
)

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	// return l.Addr().(*net.TCPAddr).Port, nil
	return 8888, nil

}

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
