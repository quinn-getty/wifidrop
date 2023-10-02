package main

import (
	"os"
	"os/signal"

	"github.com/quinn-getty/airdrop-go/server"
	"github.com/quinn-getty/airdrop-go/utils"
)

func main() {
	port, _ := utils.GetFreePort()
	go server.Run(port)
	// cmd := chrome.Open(fmt.Sprintf("http://127.0.0.1:%d/static", port))
	chSignal := listenToInterpt()

	select {
	case <-chSignal:
		// cmd.Process.Kill()
	}
}

func listenToInterpt() chan os.Signal {
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)
	return chSignal
}
