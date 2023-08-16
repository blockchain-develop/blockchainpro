package terra

import (
	"fmt"
	"testing"
)

func TestGetNodeInfo(t *testing.T) {
	nodeInfo, err := GetNodeInfo(MainNet)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", nodeInfo)
}

func TestGetExchangeRate(t *testing.T) {
	exchangeRate, err := GetExchangeRate(MainNet, "uusd")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", exchangeRate.Result.String())
}

func TestGetContractInfo(t *testing.T) {
	contractInfo, err := GetContractInfo(MainNet, "terra1m6ywlgn6wrjuagcmmezzz2a029gtldhey5k552")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", contractInfo.ContractInfo)
}

func TestGetSwapPool(t *testing.T) {
	swapPool, err := QueryContract(MainNet, "terra1m6ywlgn6wrjuagcmmezzz2a029gtldhey5k552")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", swapPool)
}
