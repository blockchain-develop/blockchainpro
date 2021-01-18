package binance

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBinanceSdk(t *testing.T) {
	sdk := DefaultBinanceSdk()
	tickers, err := sdk.QuotesLatest()
	if err != nil {
		panic(err)
	}
	res, _ := json.Marshal(tickers)
	fmt.Printf("result: %s\n", string(res))
}
