package layer2

import (
	"encoding/hex"
	"fmt"
	"github.com/ontio/ontology/common"
	"testing"
	"time"
)

func TestLayer2Deposit(t *testing.T) {
	// create alliance sdk
	layer2_sdk := newLayer2Sdk()
	account_operator, _ := newLayer2OperatorAccount(layer2_sdk)
	account_user, _ := newLayer2UserAccount(layer2_sdk)
	//
	balance := getLayer2Balance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
	//
	var amount uint64 = 100000000
	txhash, err := layer2DepositTransfer(layer2_sdk, account_operator, account_user.Address, amount)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
	} else {
		fmt.Printf("tx hash: %s\n", txhash.ToHexString())
	}
	//
	balance = getLayer2Balance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
}

func TestLayer2Withdraw(t *testing.T) {
	// create alliance sdk
	layer2_sdk := newLayer2Sdk()
	account_user, _ := newLayer2UserAccount(layer2_sdk)
	//
	balance := getLayer2Balance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
	//
	var amount uint64 = 200000000
	txhash, err := layer2WithdrawTransfer(layer2_sdk, account_user, account_user.Address, amount)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
	} else {
		fmt.Printf("tx hash: %s\n", txhash.ToHexString())
	}
	//
	balance = getLayer2Balance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
}

func TestLayer2Transfer(t *testing.T) {
	// create alliance sdk
	layer2_sdk := newLayer2Sdk()
	account_operator, _ := newLayer2OperatorAccount(layer2_sdk)
	account_user, _ := newLayer2UserAccount(layer2_sdk)
	//
	balance := getLayer2Balance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
	//
	var amount uint64 = 200000000
	txhash, err := layer2Transfer(layer2_sdk, account_user, account_user.Address, account_operator.Address, amount)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
	} else {
		fmt.Printf("tx hash: %s\n", txhash.ToHexString())
	}
	//
	balance = getLayer2Balance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
}

func TestLayer2TransferPerformance(t *testing.T) {
	// create alliance sdk
	layer2_sdk := newLayer2Sdk()
	account_operator, _ := newLayer2OperatorAccount(layer2_sdk)
	account_user, _ := newLayer2UserAccount(layer2_sdk)
	//
	{
		balance := getLayer2Balance(layer2_sdk, account_user.Address)
		fmt.Printf("amount of address1 %s is: %d\n", account_user.Address.ToBase58(), balance)
		balance1 := getLayer2Balance(layer2_sdk, account_operator.Address)
		fmt.Printf("amount of address2 %s is: %d\n", account_operator.Address.ToBase58(), balance1)
	}
	//
	var amount uint64 = 100
	var txCounter int = 1000000;

	start := time.Now().Unix()
	for i := 0;i < txCounter;i ++ {
		txhash, err := layer2Transfer(layer2_sdk, account_user, account_user.Address, account_operator.Address, amount)
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
		} else {
			fmt.Printf("tx hash: %s\n", txhash.ToHexString())
		}
	}
	end := time.Now().Unix()
	//
	time.Sleep(time.Second * 5)

	//
	{
		balance := getLayer2Balance(layer2_sdk, account_user.Address)
		fmt.Printf("amount of address1 %s is: %d\n", account_user.Address.ToBase58(), balance)
		balance1 := getLayer2Balance(layer2_sdk, account_operator.Address)
		fmt.Printf("amount of address2 %s is: %d\n", account_operator.Address.ToBase58(), balance1)
	}
	fmt.Printf("tx counter: %d, time: %d\n", txCounter, end - start)
}

func TestLayer2Balance(t *testing.T) {
	// create alliance sdk
	layer2_sdk := newLayer2Sdk()
	account_user, _ := newLayer2UserAccount(layer2_sdk)
	//
	balance := getLayer2Balance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
}

func TestGetCurrentHeight(t *testing.T) {
	layer2_sdk := newLayer2Sdk()
	height, err := layer2_sdk.GetCurrentBlockHeight()
	if err != nil {
		fmt.Printf("get current block height err: %s", err.Error())
	}
	fmt.Printf("current block height: %d\n", height)
}

func TestGetBlock(t *testing.T) {
	layer2_sdk := newLayer2Sdk()
	block, err := layer2_sdk.GetBlockByHeight(20)
	if err != nil {
		fmt.Printf("get current block height err: %s", err.Error())
	}
	fmt.Printf("get current block successful!, State root hash: %s\n", block.Header.StateRoot.ToHexString())
}

func TestCommitLayer2State2Ontology(t *testing.T) {
	//
	ontSdk := newOntologySdk()
	contractAddress, _ := common.AddressFromHexString(LAYER2_CONTRACT)
	depositids := make([]int, 0)
	for i := 0;i < 2;i ++ {
		depositids = append(depositids, 3 + i)
	}
	withdrawAmounts := make([]uint64, 0)
	toAddresses := make([]common.Address, 0)
	assetAddress := make([][]byte, 0)
	for i := 0;i < 1;i ++ {
		withdrawAmounts = append(withdrawAmounts, 3)
		toAddress, _ := common.AddressFromBase58("AMUGPqbVJ3TG6pe7xRpxxaeh4ai4fu9ahc")
		toAddresses = append(toAddresses,toAddress)
		tokenAddress, _ := hex.DecodeString("0000000000000000000000000000000000000002")
		assetAddress = append(assetAddress, tokenAddress)
	}
	tx, err := ontSdk.NeoVM.NewNeoVMInvokeTransaction(500, 40000, contractAddress, []interface{}{"updateState", []interface{}{
		"0000000000000000000000000000000000000000000000000000000000000000", 6, "1.0.0",
		depositids, []interface{}{},[]interface{}{},[]interface{}{}}})
	if err != nil {
		fmt.Printf("new transaction failed!")
	}
	account_operator, err := newOntologyOperatorAccount(ontSdk)
	if err != nil {
		fmt.Printf("newOntologyAccount failed!")
	}
	ontSdk.SetPayer(tx, account_operator.Address)
	err = ontSdk.SignToTransaction(tx, account_operator)
	if err != nil {
		fmt.Printf("SignToTransaction failed!")
	}
	txHash, err := ontSdk.SendTransaction(tx)
	if err != nil {
		fmt.Printf("SignToTransaction failed! err: %s", err.Error())
	}
	fmt.Printf("hash: %s", txHash.ToHexString())
}

func TestOntologyDeposit(t *testing.T) {
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

func TestGetLayer2StateByHeight(t *testing.T) {
	//
	ontSdk := newOntologySdk()
	contractAddress, _ := common.AddressFromHexString(LAYER2_CONTRACT)
	tx, err := ontSdk.NeoVM.NewNeoVMInvokeTransaction(0, 0, contractAddress, []interface{}{"getStateRootByHeight", []interface{}{1}})
	if err != nil {
		fmt.Printf("new transaction failed!")
	}
	result, err := ontSdk.PreExecTransaction(tx)
	if err != nil {
		fmt.Printf("PreExecTransaction failed! err: %s", err.Error())
		return
	}
	if result == nil {
		fmt.Printf("can not find the result")
		return
	}
	tt, _ := result.Result.ToArray()
	if len(tt) != 3 {
		fmt.Printf("result is not right")
		return
	}
	item0,_ := tt[0].ToString()
	item1,_ := tt[1].ToInteger()
	item2,_ := tt[2].ToInteger()
	fmt.Printf("item0: %s, item1: %d, item2: %d\n", item0, item1, item2)
}

func TestGetLayer2CurrentHeight(t *testing.T) {
	//
	ontSdk := newOntologySdk()
	contractAddress, _ := common.AddressFromHexString(LAYER2_CONTRACT)
	tx, err := ontSdk.NeoVM.NewNeoVMInvokeTransaction(0, 0, contractAddress, []interface{}{"getCurrentHeight", []interface{}{}})
	if err != nil {
		fmt.Printf("new transaction failed!")
	}
	result, err := ontSdk.PreExecTransaction(tx)
	if err != nil {
		fmt.Printf("PreExecTransaction failed! err: %s", err.Error())
		return
	}
	if result == nil {
		fmt.Printf("can not find the result")
		return
	}
	height, _ := result.Result.ToInteger()
	fmt.Printf("height: %d\n", height.Uint64())
}

func TestOntologyBalance(t *testing.T) {
	// create alliance sdk
	ontSdk := newOntologySdk()
	account_user, _ := newOntologyUserAccount(ontSdk)
	//
	balance := getOntologyBalance(ontSdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
}

func TestCreateLayer2Account(t *testing.T) {
	createLayer2Account()
}

func TestCreateOntologyAccount(t *testing.T) {
	createOntologyAccount()
}
