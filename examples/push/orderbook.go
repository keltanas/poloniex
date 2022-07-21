// the following code shows
// how to access OrderBook fields.
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

	var m polo.WSOrderBook

	for {
		receive := <-ws.Subs["USDT_BTC"]
		updates := receive.([]polo.MarketUpdate)
		for _, v := range updates {
			if v.TypeUpdate == "OrderBookRemove" || v.TypeUpdate == "OrderBookModify" {
				m = v.Data.(polo.WSOrderBook)

				fmt.Printf("Rate:%f, Type:%s, Amount:%f\n",
					m.Rate, m.TypeOrder, m.Amount)
			}
		}
	}
}
