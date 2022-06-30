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

	err = ws.SubscribeTicker()
	if err != nil {
		return
	}

	for {
		fmt.Println(<-ws.Subs["TICKER"])
	}
}
