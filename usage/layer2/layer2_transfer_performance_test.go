package layer2

import (
	"fmt"
	"testing"
	"time"
)

func TestLayer2TransferPerformance(t *testing.T) {
	// create alliance sdk
	layer2_sdk := newLayer2Sdk()
	account_operator, _ := newLayer2OperatorAccount(layer2_sdk)
	account_user, _ := newLayer2UserAccount(layer2_sdk)
	//
	{
		balance := getLayer2OngBalance(layer2_sdk, account_user.Address)
		fmt.Printf("amount of address1 %s is: %d\n", account_user.Address.ToBase58(), balance)
		balance1 := getLayer2OngBalance(layer2_sdk, account_operator.Address)
		fmt.Printf("amount of address2 %s is: %d\n", account_operator.Address.ToBase58(), balance1)
	}
	//
	var amount uint64 = 100
	var txCounter int = 1000000;

	start := time.Now().Unix()
	for i := 0;i < txCounter;i ++ {
		txhash, err := layer2TransferOng(layer2_sdk, account_user, account_user.Address, account_operator.Address, amount)
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
		balance := getLayer2OngBalance(layer2_sdk, account_user.Address)
		fmt.Printf("amount of address1 %s is: %d\n", account_user.Address.ToBase58(), balance)
		balance1 := getLayer2OngBalance(layer2_sdk, account_operator.Address)
		fmt.Printf("amount of address2 %s is: %d\n", account_operator.Address.ToBase58(), balance1)
	}
	fmt.Printf("tx counter: %d, time: %d\n", txCounter, end - start)
}
