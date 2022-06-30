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

	err = ws.SubscribeMarket("USDT_BTC")
	if err != nil {
		return
	}

	go func() {
		time.Sleep(time.Second * 10)
		_ = ws.UnsubscribeMarket("USDT_BTC")
	}()

	for {
		fmt.Println(<-ws.Subs["USDT_BTC"])
	}
}
