package bybit

import (
	"encoding/json"
	"fmt"
	"github.com/hirokisan/bybit"
	"testing"
)

func TestOrderBook(t *testing.T) {
	client := bybit.NewTestClient()
	orderbook, err := client.Market().OrderBook(bybit.SymbolInverse(bybit.SymbolUSDTETH))
	if err != nil {
		panic(err)
	}
	orderbookJson, _ := json.MarshalIndent(orderbook.Result, "", "    ")
	fmt.Printf("order book: %s\n", orderbookJson)
}
