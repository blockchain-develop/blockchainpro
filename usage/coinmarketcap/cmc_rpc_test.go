package coinmarketcap

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestListingsLatest(t *testing.T) {
	sdk := DefaultCoinMarketCapSdk()
	data, err := sdk.ListingsLatest()
	if err != nil {
		panic(err)
	}
	dataJson, _ := json.Marshal(data)
	fmt.Printf("%s\n", dataJson)
}

func TestQuotesLatest(t *testing.T) {
	sdk := DefaultCoinMarketCapSdk()
	data, err := sdk.QuotesLatest()
	if err != nil {
		panic(err)
	}
	dataJson, _ := json.Marshal(data)
	fmt.Printf("%s\n", dataJson)
}
