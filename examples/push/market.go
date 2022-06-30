package main

import (
	"fmt"

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

	for {
		fmt.Println(<-ws.Subs["USDT_BTC"])
	}
}
