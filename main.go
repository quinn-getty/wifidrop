package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/quinn-getty/airdrop-go/chrome"
	"github.com/quinn-getty/airdrop-go/server"
	"github.com/quinn-getty/airdrop-go/utils"
)

func main() {
	port, err := utils.GetFreePort()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		server.Run(port)
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
