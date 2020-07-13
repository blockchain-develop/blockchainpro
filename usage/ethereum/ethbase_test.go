package ethereum

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
	ethcommon "github.com/ethereum/go-ethereum/common"
)

func TestAddress1(t *testing.T) {
	addr := ethcommon.HexToAddress("344cfc3b8635f72f14200aaf2168d9f75df86fd3")
	fmt.Printf("address: %s\n", addr.String())
}

func TestGetTransaction(t *testing.T) {
	url := "http://18.139.17.85:10331"
	ethclient, err := ethclient.Dial(url)
	if err != nil {
		fmt.Printf("getmocktokeninfo - cannot dial sync node, err: %s", err)
		return
	}
	tx, _, err := ethclient.TransactionByHash(nil, ethcommon.HexToHash("0xfd40b4f7ebf6bebe5f7cab25ed2ff4938c6ba68d055fecaff78470aaf170f09e"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("tx: %v\n", tx)
}
