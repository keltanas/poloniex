package main

import (
	"fmt"

	"go.uber.org/zap"

	polo "github.com/keltanas/poloniex"
)

func main() {

	log, _ := zap.NewDevelopment()
	ws, err := polo.NewWSClient(log)
	if err != nil {
		return
	}

	err = ws.SubscribeMarket("USDT_BTC")
	if err != nil {
		return
	}

	for {
		fmt.Println(<-ws.Subs["USDT_BTC"])
	}
}
