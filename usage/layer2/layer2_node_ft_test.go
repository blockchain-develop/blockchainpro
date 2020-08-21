package layer2

import (
	"fmt"
	"testing"
)

func TestLayer2OngBalance(t *testing.T) {
	// create alliance sdk
	layer2_sdk := newLayer2Sdk()
	account_user, _ := newLayer2UserAccount(layer2_sdk)
	//
	balance := getLayer2OngBalance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
}

func TestLayer2OntBalance(t *testing.T) {
	// create alliance sdk
	layer2_sdk := newLayer2Sdk()
	account_user, _ := newLayer2UserAccount(layer2_sdk)
	//
	balance := getLayer2OntBalance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
}

func TestGetLayer2CurrentHeight(t *testing.T) {
	layer2_sdk := newLayer2Sdk()
	height, err := layer2_sdk.GetCurrentBlockHeight()
	if err != nil {
		fmt.Printf("get current block height err: %s", err.Error())
	}
	fmt.Printf("current block height: %d\n", height)
}

func TestGetLayer2Block(t *testing.T) {
	layer2_sdk := newLayer2Sdk()
	block, err := layer2_sdk.GetLayer2BlockByHeight(20)
	if err != nil {
		fmt.Printf("get current block height err: %s", err.Error())
	}
	fmt.Printf("get current block successful!, State root hash: %s\n", block.Header.StateRoot.ToHexString())
}

func TestGetLayer2Transaction(t *testing.T) {
	layer2Sdk := newLayer2Sdk()
	txHash := ""
	tx, err := layer2Sdk.GetTransaction(txHash)
	if err != nil {
		panic(err)
	}
	if tx == nil {
		panic("this is no this tx!")
	}
	getTxHash := tx.Hash()
	fmt.Printf("tx hash is: %s\n", getTxHash.ToHexString())
}
