package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/quinn-getty/airdrop-go/chrome"
	"github.com/quinn-getty/airdrop-go/lorca"
	"github.com/quinn-getty/airdrop-go/server"
	"github.com/quinn-getty/airdrop-go/utils"
)

func test() {
	ui, err := lorca.New("data:text/html,"+url.PathEscape(`
	<html>
		<head><title>Hello</title></head>
		<body><h1>Hello, world!</h1></body>
	</html>
	`), "", 480, 320, "--remote-allow-origins=*")
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}

func main() {
	test()
	port, _ := utils.GetFreePort()
	go server.Run(port)
	cmd := chrome.Open(fmt.Sprintf("http://127.0.0.1:%d/static", port))
	chSignal := listenToInterpt()
	go func() {
		for {
			time.Sleep(2 * time.Second) // 2 秒间隔
			log.Print("-----", cmd.ProcessState)
		}
	}()

	select {
	case <-chSignal:
		cmd.Process.Kill()
	}
}

func listenToInterpt() chan os.Signal {
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)
	return chSignal
}
