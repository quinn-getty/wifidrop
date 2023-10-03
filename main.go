package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/quinn-getty/airdrop-go/lorca"
	"github.com/quinn-getty/airdrop-go/server"
	"github.com/quinn-getty/airdrop-go/utils"
)

func main() {
	var ui lorca.UI
	port, _ := utils.GetFreePort()

	go server.Run(port)
	ui, _ = lorca.New(fmt.Sprintf("http://127.0.0.1:%d/static", port), "", 480, 320, "--remote-allow-origins=*")
	chSignal := listenToInterpt()

	select {
	case <-chSignal:
	case <-ui.Done():
	}
}

func listenToInterpt() chan os.Signal {
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)
	return chSignal
}
