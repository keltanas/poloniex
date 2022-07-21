package main

import (
	"fmt"
	"time"

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
	log.Info("Subscribed to USDT_BTC channel.")
	go func() {
		for {
			fmt.Println(<-ws.Subs["USDT_BTC"], ws.Subs)
		}
	}()
	time.Sleep(time.Second * 10)

	err = ws.UnsubscribeMarket("USDT_BTC")
	if err != nil {
		return
	}
	log.Info("Unsubscribed from USDT_BTC channel.")
	time.Sleep(time.Second * 10)

	err = ws.SubscribeMarket("USDT_BTC")
	if err != nil {
		panic(err)
		return
	}
	log.Info("Subscribed to USDT_BTC channel.")
	time.Sleep(time.Second * 50)
}
