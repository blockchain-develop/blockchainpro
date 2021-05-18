package poly

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetTransactionByHash(t *testing.T) {
	url := "http://13.92.155.62:20336"
	hash := "2e7c2708e6753f74f7fbf6aac75274198221656e57b98523a0178bcd73577157"
	sdk := NewSdk(url)
	tx, err := sdk.GetTransaction(hash)
	if err != nil {
		panic(err)
	}
	json, _ := json.Marshal(tx)
	fmt.Printf("%s\n", json)
	height, err := sdk.GetBlockHeightByTxHash(hash)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", height)
}

func TestGetLatestHeight(t *testing.T) {
	url := "http://beta1.poly.network:20336"
	//url := "http://138.91.6.125:20336"
	sdk := NewSdk(url)
	height, err := sdk.GetCurrentBlockHeight()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", height)
}

func TestGetTransactionHeight(t *testing.T) {
	url := "http://124.156.226.204:20336"
	//url := "http://13.92.155.62:20336"
	hash := "a7b6b94dfe36fe33445d28a25c43e27cfffda88b58af6629ed1306ba6ce08440"
	sdk := NewSdk(url)
	height, err := sdk.GetBlockHeightByTxHash(hash)
	if err != nil {
		panic(err)
	}
	fmt.Printf("height: %d\n", height)
}

func TestGetBlock(t *testing.T) {
	url := "http://13.92.155.62:20336"
	//url := "http://13.92.155.62:20336"
	sdk := NewSdk(url)
	block, err := sdk.GetBlockByHeight(1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", block)
}

func TestCreatePolyAccount(t *testing.T) {
	createPolyAccount()
}

