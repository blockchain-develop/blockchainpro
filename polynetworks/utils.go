package polynetworks

import (
	"github.com/ontio/multi-chain-go-sdk"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type HeaderWithDifficultySum struct {
	Header        types.Header `json:"header"`
	DifficultySum *big.Int     `json:"difficultySum"`
}

func newMultiChanSdk() *multi_chain_go_sdk.MultiChainSdk {
	mcSdk := multi_chain_go_sdk.NewMultiChainSdk()
	// test net
	mcSdk.NewRpcClient().SetAddress("http://40.115.182.238:40336")
	//mcSdk.NewRpcClient().SetAddress("http://172.168.3.73:40336")
	return mcSdk
}
