package polynetworks

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ontio/multi-chain-go-sdk"
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

func revertBytes(data []byte) []byte {
	n := len(data)
	newdata := make([]byte, n)
	for i := 0;i < n;i ++ {
		newdata[i] = data[n - 1 - i]
	}
	return newdata
}

func HexStringReverse(value string) string {
	aa, _ := hex.DecodeString(value)
	bb := revertBytes(aa)
	return hex.EncodeToString(bb)
}
