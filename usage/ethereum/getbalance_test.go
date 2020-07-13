package ethereum

import (
	"fmt"
	"github.com/blockchainpro/usage/ethereum/ethtools/btcx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func TestGetERC20Balance(t *testing.T) {
	url := "http://18.139.17.85:10331"
	ethclient, err := ethclient.Dial(url)
	if err != nil {
		fmt.Printf("getmocktokeninfo - cannot dial sync node, err: %s", err)
		return
	}
	btcxContractAddress := "0x700CA49ccA3803316124D2A8a44498ABB3E9cF51"
	contractAddress := common.HexToAddress(btcxContractAddress)
	instance, err := btcx_abi.NewBTCX(contractAddress, ethclient)
	if err != nil {
		fmt.Printf("getmocktokeninfo - new eth cross chain failed: %s", err.Error())
		return
	}
	userAddress := "0x344cFc3B8635f72F14200aAf2168d9f75df86FD3"
	balance, err := instance.BalanceOf(nil, common.HexToAddress(userAddress))
	fmt.Printf("getmocktokeninfo - balance of %s: %d\n", userAddress, balance.Uint64())
}

