package bybit

import (
	"encoding/json"
	"fmt"
	"github.com/hirokisan/bybit"
	"testing"
)

func TestSymbol(t *testing.T) {
	client := bybit.NewTestClient()
	symbols, err := client.Market().Symbols()
	if err != nil {
		panic(err)
	}
	symbolsJson, _ := json.MarshalIndent(symbols.Result, "", "    ")
	fmt.Printf("symbols: %s\n", symbolsJson)
}
