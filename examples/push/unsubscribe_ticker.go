package main

import (
	"fmt"
	"time"

	polo "github.com/keltanas/poloniex"
)

func main() {

	ws, err := polo.NewWSClient()
	if err != nil {
		return
	}

	err = ws.SubscribeTicker()
	if err != nil {
		return
	}

	go func() {
		time.Sleep(time.Second * 10)
		_ = ws.UnsubscribeTicker()
	}()

	for {
		fmt.Println(<-ws.Subs["TICKER"])
	}
}
