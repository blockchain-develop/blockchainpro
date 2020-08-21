package layer2

import (
	"fmt"
	"github.com/ontio/ontology/common"
	"testing"
)

func TestLayer2DepositOng_ut(t *testing.T) {
	// create alliance sdk
	layer2_sdk := newLayer2Sdk()
	account_operator, _ := newLayer2OperatorAccount(layer2_sdk)
	account_user, _ := newLayer2UserAccount(layer2_sdk)
	//
	balance := getLayer2OngBalance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
	//
	var amount uint64 = 100000000
	txhash, err := layer2DepositTransferOng(layer2_sdk, account_operator, account_user.Address, amount)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
	} else {
		fmt.Printf("tx hash: %s\n", txhash.ToHexString())
	}
	//
	balance = getLayer2OngBalance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
}

func TestLayer2WithdrawOng_ut(t *testing.T) {
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

func TestLayer2TransferOng_ut(t *testing.T) {
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

func TestCommitLayer2State2Ontology_ut(t *testing.T) {
	//
	ontSdk := newOntologySdk()
	contractAddress, _ := common.AddressFromHexString(LAYER2_CONTRACT)
	account_operator, err := newOntologyOperatorAccount(ontSdk)
	if err != nil {
		fmt.Printf("newOntologyAccount failed!")
	}
	txHash, err := updateLayer2State(ontSdk, account_operator, contractAddress)
	if err != nil {
		panic(err)
	}
	fmt.Printf("hash: %s", txHash.ToHexString())
}

func TestCommitLayer2State2OntologyBatch_ut(t *testing.T) {
	//
	ontSdk := newOntologySdk()
	contractAddress, _ := common.AddressFromHexString(LAYER2_CONTRACT)
	account_operator, err := newOntologyOperatorAccount(ontSdk)
	if err != nil {
		fmt.Printf("newOntologyAccount failed!")
	}
	txHash, err := updateLayer2StateBatch(ontSdk, account_operator, contractAddress)
	if err != nil {
		panic(err)
	}
	fmt.Printf("hash: %s", txHash.ToHexString())
}

func TestCreateLayer2Account_ut(t *testing.T) {
	createLayer2Account()
}

func TestCreateOntologyAccount_ut(t *testing.T) {
	createOntologyAccount()
}
