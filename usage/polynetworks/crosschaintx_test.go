package polynetworks

import (
	"fmt"
	"github.com/blockchainpro/usage/utils"
	"github.com/ontio/multi-chain/common"
	"testing"
	"time"
)

var (
	CrossChainManagerContractAddress, _ = common.AddressParseFromBytes([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03})
)

func TestCrossChainTx_SCANTX(t  *testing.T) {
	sdk := newMultiChanSdk()
	currentHeight, err := sdk.GetCurrentBlockHeight()
	if err != nil {
		fmt.Printf("loadOntCrossTxFromChain: get current block height %s", err)
		return
	}
	fmt.Printf("current block height: %d\n", currentHeight)

	//start := uint32(263867)
	start := uint32(500000)
	end := start + 100000
	for start < end {
		currentHeight, err := sdk.GetCurrentBlockHeight()
		if err != nil {
			fmt.Printf("loadOntCrossTxFromChain: get current block height %s", err)
			return
		}
		if currentHeight < start + 1 {
			time.Sleep(time.Second * 1)
			continue
		}
		start ++
		events, err := sdk.GetSmartContractEventByBlock(start)
		if err != nil {
			return
		}
		for _, event := range events {
			for _, notify := range event.Notify {
				if notify.ContractAddress != CrossChainManagerContractAddress.ToHexString() {
					continue
				}
				fmt.Printf("block height: %d, contractAddress: %s\n", start, notify.ContractAddress)
				fmt.Printf("tx hash: %s, state:%d, gas: %d\n", event.TxHash, event.State, 0)
				states := notify.States.([]interface{})
				contractMethod, _ := states[0].(string)
				if contractMethod != "makeProof" && contractMethod != "btcTxToRelay" {
					continue
				}
				fchainid := uint32(states[1].(float64))
				tchainid := uint32(states[2].(float64))

				fmt.Printf("height: %d, txHash: %s, fchainId: %d, tchainid: %d, ", start, event.TxHash, fchainid, tchainid)
				if tchainid == 1 {
					fmt.Printf("ftxhash: %s, key: %s", states[4].(string), states[3].(string))
				} else {
					if fchainid == 2 {
						fmt.Printf("ftxhash: %s", states[3].(string))
					} else if fchainid == 3 || fchainid == 1 {
						fmt.Printf("ftxhash: %s", utils.HexStringReverse(states[3].(string)))
					}
				}
			}
		}
	}
}

// ok
func TestCrossChainTx_BTC2ETH(t  *testing.T) {
	sdk := newMultiChanSdk()
	currentHeight, err := sdk.GetCurrentBlockHeight()
	if err != nil {
		fmt.Printf("loadOntCrossTxFromChain: get current block height %s", err)
		return
	}
	fmt.Printf("current block height: %d\n", currentHeight)
	end := uint32(250795)
	start := end - uint32(10)
	for start < end {
		start ++
		events, err := sdk.GetSmartContractEventByBlock(start)
		if err != nil {
			return
		}
		for _, event := range events {
			for _, notify := range event.Notify {
				fmt.Printf("block height: %d, contractAddress: %s\n", start, notify.ContractAddress)
				if notify.ContractAddress != CrossChainManagerContractAddress.ToHexString() {
					continue
				}
				states := notify.States.([]interface{})
				contractMethod, _ := states[0].(string)
				fmt.Printf("tx hash: %s, state:%d, gas: %d, method: %s\n", event.TxHash, event.State, 0, contractMethod)
				if contractMethod != "makeProof" {
					continue
				}
				fchainid := uint32(states[1].(float64))
				tchainid := uint32(states[2].(float64))

				fmt.Printf("height: %d, txHash: %s, fchainId: %d, tchainid: %d, ", start, event.TxHash, fchainid, tchainid)
				if tchainid == 1 {
					fmt.Printf("ftxhash: %s, key: %s\n", states[4].(string), states[3].(string))
				} else {
					if fchainid == 2 {
						fmt.Printf("ftxhash: %s\n", states[3].(string))
					} else {
						fmt.Printf("ftxhash: %s\n", utils.HexStringReverse(states[3].(string)))
					}
				}
			}
		}
	}
}

// ok
func TestCrossChainTx_ETH2BTC(t  *testing.T) {
	sdk := newMultiChanSdk()
	currentHeight, err := sdk.GetCurrentBlockHeight()
	if err != nil {
		fmt.Printf("loadOntCrossTxFromChain: get current block height %s", err)
		return
	}
	fmt.Printf("current block height: %d\n", currentHeight)
	end := uint32(263960)
	start := end - uint32(10)
	for start < end {
		start ++
		events, err := sdk.GetSmartContractEventByBlock(start)
		if err != nil {
			return
		}
		for _, event := range events {
			for _, notify := range event.Notify {
				if notify.ContractAddress != CrossChainManagerContractAddress.ToHexString() {
					continue
				}
				states := notify.States.([]interface{})
				contractMethod, _ := states[0].(string)
				fmt.Printf("tx hash: %s, state:%d, gas: %d, method: %s\n", event.TxHash, event.State, 0, contractMethod)
				if contractMethod != "makeProof" && contractMethod != "btcTxToRelay" {
					continue
				}
				fchainid := uint32(states[1].(float64))
				tchainid := uint32(states[2].(float64))

				fmt.Printf("height: %d, txHash: %s, fchainId: %d, tchainid: %d, ", start, event.TxHash, fchainid, tchainid)
				if tchainid == 1 {
					fmt.Printf("ftxhash: %s, key: %s", states[4].(string), states[3].(string))
				} else {
					if fchainid == 2 {
						fmt.Printf("ftxhash: %s", states[3].(string))
					} else {
						fmt.Printf("ftxhash: %s", utils.HexStringReverse(states[3].(string)))
					}
				}
			}
		}
	}
}

// ok
func TestCrossChainTx_BTC2ONT(t  *testing.T) {
	sdk := newMultiChanSdk()
	currentHeight, err := sdk.GetCurrentBlockHeight()
	if err != nil {
		fmt.Printf("loadOntCrossTxFromChain: get current block height %s", err)
		return
	}
	fmt.Printf("current block height: %d\n", currentHeight)
	end := uint32(250795)
	start := end - uint32(10)
	for start < end {
		start ++
		events, err := sdk.GetSmartContractEventByBlock(start)
		if err != nil {
			return
		}
		for _, event := range events {
			for _, notify := range event.Notify {
				//fmt.Printf("block height: %d, contractAddress: %s\n", start, notify.ContractAddress)
				if notify.ContractAddress != CrossChainManagerContractAddress.ToHexString() {
					continue
				}
				states := notify.States.([]interface{})
				contractMethod, _ := states[0].(string)
				fmt.Printf("tx hash: %s, state:%d, gas: %d， method: %s\n", event.TxHash, event.State, 0, contractMethod)
				if contractMethod != "makeProof" {
					continue
				}
				fchainid := uint32(states[1].(float64))
				tchainid := uint32(states[2].(float64))

				fmt.Printf("height: %d, txHash: %s, fchainId: %d, tchainid: %d, ", start, event.TxHash, fchainid, tchainid)
				if tchainid == 1 {
					fmt.Printf("ftxhash: %s, key: %s\n", states[4].(string), states[3].(string))
				} else {
					if fchainid == 2 {
						fmt.Printf("ftxhash: %s\n", states[3].(string))
					} else {
						fmt.Printf("ftxhash: %s\n", utils.HexStringReverse(states[3].(string)))
					}
				}
			}
		}
	}
}

// ok
func TestCrossChainTx_ONT2BTC(t  *testing.T) {
	sdk := newMultiChanSdk()
	currentHeight, err := sdk.GetCurrentBlockHeight()
	if err != nil {
		fmt.Printf("loadOntCrossTxFromChain: get current block height %s", err)
		return
	}
	fmt.Printf("current block height: %d\n", currentHeight)
	end := uint32(263940)
	start := end - uint32(10)
	for start < end {
		start ++
		events, err := sdk.GetSmartContractEventByBlock(start)
		if err != nil {
			return
		}
		for _, event := range events {
			for _, notify := range event.Notify {
				if notify.ContractAddress != CrossChainManagerContractAddress.ToHexString() {
					continue
				}
				states := notify.States.([]interface{})
				contractMethod, _ := states[0].(string)
				fmt.Printf("tx hash: %s, state:%d, gas: %d, method: %s\n", event.TxHash, event.State, 0, contractMethod)
				if contractMethod != "makeProof" && contractMethod != "btcTxToRelay" {
					continue
				}
				fchainid := uint32(states[1].(float64))
				tchainid := uint32(states[2].(float64))

				fmt.Printf("height: %d, txHash: %s, fchainId: %d, tchainid: %d, ", start, event.TxHash, fchainid, tchainid)
				if tchainid == 1 {
					if fchainid == 2 {
						fmt.Printf("ftxhash: %s, key: %s", states[4].(string), states[3].(string))
					} else {
						fmt.Printf("ftxhash: %s, key: %s", utils.HexStringReverse(states[4].(string)), states[3].(string))
					}
				} else {
					if fchainid == 2 {
						fmt.Printf("ftxhash: %s", states[3].(string))
					} else {
						fmt.Printf("ftxhash: %s", utils.HexStringReverse(states[3].(string)))
					}
				}
			}
		}
	}
}

// ok
func TestCrossChainTx_ONT2ETH(t  *testing.T) {
	sdk := newMultiChanSdk()
	currentHeight, err := sdk.GetCurrentBlockHeight()
	if err != nil {
		fmt.Printf("loadOntCrossTxFromChain: get current block height %s", err)
		return
	}
	fmt.Printf("current block height: %d\n", currentHeight)
	end := uint32(249721)
	start := end - uint32(10)
	for start < end {
		start ++
		events, err := sdk.GetSmartContractEventByBlock(start)
		if err != nil {
			return
		}
		for _, event := range events {
			for _, notify := range event.Notify {
				//fmt.Printf("block height: %d, contractAddress: %s\n", start, notify.ContractAddress)
				if notify.ContractAddress != CrossChainManagerContractAddress.ToHexString() {
					continue
				}
				states := notify.States.([]interface{})
				contractMethod, _ := states[0].(string)
				fmt.Printf("tx hash: %s, state:%d, gas: %d, method: %s\n", event.TxHash, event.State, 0, contractMethod)
				if contractMethod != "makeProof" {
					continue
				}
				fchainid := uint32(states[1].(float64))
				tchainid := uint32(states[2].(float64))

				fmt.Printf("height: %d, txHash: %s, fchainId: %d, tchainid: %d, ", start, event.TxHash, fchainid, tchainid)
				if tchainid == 1 {
					fmt.Printf("ftxhash: %s, key: %s\n", states[4].(string), states[3].(string))
				} else {
					if fchainid == 2 {
						fmt.Printf("ftxhash: %s\n", states[3].(string))
					} else {
						fmt.Printf("ftxhash: %s\n", utils.HexStringReverse(states[3].(string)))
					}
				}
			}
		}
	}
}

// ok
func TestCrossChainTx_ETH2ONT(t  *testing.T) {
	sdk := newMultiChanSdk()
	currentHeight, err := sdk.GetCurrentBlockHeight()
	if err != nil {
		fmt.Printf("loadOntCrossTxFromChain: get current block height %s", err)
		return
	}
	fmt.Printf("current block height: %d\n", currentHeight)
	end := uint32(532313)
	start := end - uint32(5)
	for start < end {
		start ++
		events, err := sdk.GetSmartContractEventByBlock(start)
		if err != nil {
			return
		}
		for _, event := range events {
			for _, notify := range event.Notify {
				//fmt.Printf("block height: %d, contractAddress: %s\n", start, notify.ContractAddress)
				if notify.ContractAddress != CrossChainManagerContractAddress.ToHexString() {
					continue
				}
				states := notify.States.([]interface{})
				contractMethod, _ := states[0].(string)
				fmt.Printf("tx hash: %s, state:%d, gas: %d, method: %s\n", event.TxHash, event.State, 0, contractMethod)
				if contractMethod != "makeProof" {
					continue
				}
				fchainid := uint32(states[1].(float64))
				tchainid := uint32(states[2].(float64))

				fmt.Printf("height: %d, txHash: %s, fchainId: %d, tchainid: %d, ", start, event.TxHash, fchainid, tchainid)
				if tchainid == 1 {
					fmt.Printf("ftxhash: %s, key: %s\n", states[4].(string), states[3].(string))
				} else {
					if fchainid == 2 {
						fmt.Printf("ftxhash: %s\n", states[3].(string))
					} else {
						fmt.Printf("ftxhash: %s\n", utils.HexStringReverse(states[3].(string)))
					}
				}
			}
		}
	}
}


// ok
func TestCrossChainTx_Cosmos2ETH(t  *testing.T) {
	sdk := newMultiChanSdk()
	currentHeight, err := sdk.GetCurrentBlockHeight()
	if err != nil {
		fmt.Printf("loadOntCrossTxFromChain: get current block height %s", err)
		return
	}
	fmt.Printf("current block height: %d\n", currentHeight)
	height := uint32(532504)

		events, err := sdk.GetSmartContractEventByBlock(height)
		if err != nil {
			return
		}
		for _, event := range events {
			for _, notify := range event.Notify {
				//fmt.Printf("block height: %d, contractAddress: %s\n", start, notify.ContractAddress)
				if notify.ContractAddress != CrossChainManagerContractAddress.ToHexString() {
					continue
				}
				states := notify.States.([]interface{})
				contractMethod, _ := states[0].(string)
				fmt.Printf("tx hash: %s, state:%d, gas: %d, method: %s\n", event.TxHash, event.State, 0, contractMethod)
				if contractMethod != "makeProof" {
					continue
				}
				fchainid := uint32(states[1].(float64))
				tchainid := uint32(states[2].(float64))

				fmt.Printf("height: %d, txHash: %s, fchainId: %d, tchainid: %d, ", height, event.TxHash, fchainid, tchainid)
				if tchainid == 1 {
					fmt.Printf("ftxhash: %s, key: %s\n", states[4].(string), states[3].(string))
				} else {
					if fchainid == 2 || fchainid == 5 {
						fmt.Printf("ftxhash: %s\n", states[3].(string))
					} else {
						fmt.Printf("ftxhash: %s\n", utils.HexStringReverse(states[3].(string)))
					}
				}
			}
		}
}