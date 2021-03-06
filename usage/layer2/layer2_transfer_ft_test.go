package layer2

import (
	"encoding/hex"
	"fmt"
	"github.com/ontio/ontology/common"
	"testing"
)

func TestOntologyDeposit2Layer2_ft(t *testing.T) {
	ontSdk := newOntologySdk()
	contractAddress, _ := common.AddressFromHexString(LAYER2_CONTRACT)
	account_user, err := newOntologyUserAccount(ontSdk)
	if err != nil {
		fmt.Printf("ontology account err: %s", err.Error())
		return
	}
	tokenAddress, _ := hex.DecodeString("0000000000000000000000000000000000000002")
	txHash, err := ontologyDeposit(ontSdk, account_user, contractAddress, tokenAddress, 3000000000)
	if err != nil {
		panic(err)
	}
	fmt.Printf("hash: %s", txHash.ToHexString())
}

func TestLayer2WithdrawOng_ft(t *testing.T) {
	// create alliance sdk
	layer2_sdk := newLayer2Sdk()
	account_user, _ := newLayer2UserAccount(layer2_sdk)
	//
	balance := getLayer2OngBalance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
	//
	var amount uint64 = 200000000
	txhash, err := layer2WithdrawTransferOng(layer2_sdk, account_user, account_user.Address, amount)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
	} else {
		fmt.Printf("tx hash: %s\n", txhash.ToHexString())
	}
	//
	balance = getLayer2OngBalance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
}

func TestLayer2TransferOng_ft(t *testing.T) {
	// create alliance sdk
	layer2_sdk := newLayer2Sdk()
	account_operator, _ := newLayer2OperatorAccount(layer2_sdk)
	account_user, _ := newLayer2UserAccount(layer2_sdk)
	//
	balance := getLayer2OngBalance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
	//
	var amount uint64 = 200000000
	txhash, err := layer2TransferOng(layer2_sdk, account_user, account_user.Address, account_operator.Address, amount)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
	} else {
		fmt.Printf("tx hash: %s\n", txhash.ToHexString())
	}
	//
	balance = getLayer2OngBalance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
}

func TestGetCommitedLayer2StateByHeight_ft(t *testing.T) {
	//
	ontSdk := newOntologySdk()
	contractAddress, _ := common.AddressFromHexString(LAYER2_CONTRACT)
	stateRoot, height, err := GetCommitedLayer2StateByHeight(ontSdk, contractAddress, 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("state root hash: %s, height: %d\n", hex.EncodeToString(stateRoot), height)
}

func TestGetCurrentCommitedLayer2Height_ft(t *testing.T) {
	//
	ontSdk := newOntologySdk()
	contractAddress, _ := common.AddressFromHexString(LAYER2_CONTRACT)
	height, err := GetCommitedLayer2Height(ontSdk, contractAddress)
	if err != nil {
		panic(err)
	}
	fmt.Printf("current layer2 height: %d\n", height)
}
