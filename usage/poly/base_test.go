package poly

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetTransactionByHash(t *testing.T) {
	url := "http://138.91.6.125:20336"
	hash := "8acdadd77bcbd8337f836afda02fc1813fdc2d29285579948f43d55ee2cbb762"
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
	url := "http://138.91.6.125:20336"
	sdk := NewSdk(url)
	height, err := sdk.GetCurrentBlockHeight()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", height)
}

func TestCreatePolyAccount(t *testing.T) {
	createPolyAccount()
}

