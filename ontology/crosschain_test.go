package ontology

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joeqian10/neo-gogogo/helper"
	mccommon "github.com/ontio/multi-chain/common"
	sdk "github.com/ontio/ontology-go-sdk"
	"testing"
)

var (
	CrossChainManagerContractAddress, _ = mccommon.AddressParseFromBytes([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x09})
)

func NewClient() *sdk.OntologySdk {
	rawsdk := sdk.NewOntologySdk()
	rawsdk.NewRpcClient().SetAddress("http://polaris1.ont.io:20336")
	return rawsdk
}

func ParserOntCrossChainValue(toChainId uint32, value string) {
	args, err := hex.DecodeString(value)
	if err != nil {
		fmt.Printf("hex.DecodeString error: %v\n", err)
	}
	source := mccommon.NewZeroCopySource(args)
	assethash, _ := source.NextVarBytes()
	toaddress, _ := source.NextVarBytes()
	amount, _ := source.NextUint64()

	var assetaddress2 string
	var toaddress2 string
	if toChainId == 1 {
		assetaddress2 = "0000000000000000000000000000000000000011"
		toaddress2 = base58.Encode(toaddress)
	} else if toChainId == 2 {
		assetAddress1 := common.BytesToAddress(assethash)
		toaddress1 := common.BytesToAddress(toaddress)
		assetaddress2 = assetAddress1.String()
		toaddress2 = toaddress1.String()
	} else if toChainId == 4 {
		assetaddress1, _ := helper.UInt160FromBytes(assethash)
		toaddress1, _ := helper.UInt160FromBytes(toaddress)
		assetaddress2 = helper.ScriptHashToAddress(assetaddress1)
		toaddress2 = helper.ScriptHashToAddress(toaddress1)
	} else {
		assetaddress2 = ""
		toaddress2 = ""
	}

	fmt.Printf("to chain id: %s, to asset address: %s, to address: %s, toamount: %d\n", toChainId, assetaddress2, toaddress2, amount)
}

func TestCrossChainEvent_BTC2ONT(t *testing.T) {
	height := uint32(12575841)
	sdk := NewClient()
	events, _ := sdk.GetSmartContractEventByBlock(height)
	fmt.Printf("ontology, block height: %d, events num: %d\n", height, len(events))
	for _, event := range events {
		fmt.Printf("saveOntCrossTxsByHeight tx hash: %s, state:%d, gas: %d\n", event.TxHash, event.State, event.GasConsumed)
		for _, notify := range event.Notify {
			{
				states := notify.States.([]interface{})
				contractMethod, _ := states[0].(string)
				if contractMethod == "lock" {
					sourceAssetAddress, _ := states[1].(string)
					tochainid, _ := states[2].(uint64)
					targetAssetAddress, _ := states[3].(string)
					fromAddress, _ := states[4].(string)
					toAddress, _ := states[5].(string)
					amount, _ := states[6].(uint64)
					fmt.Printf("source asset address: %s, tochainid: %d, targetassetaddress: %s, fromaddress: %s, toaddress: %s, amount:%d\n",
						sourceAssetAddress, tochainid, targetAssetAddress, fromAddress, toAddress, amount)
				}
			}
			if notify.ContractAddress != CrossChainManagerContractAddress.ToHexString() {
				continue
			}

			states := notify.States.([]interface{})
			contractMethod, _ := states[0].(string)
			fmt.Printf("contract method: %s, ", contractMethod)
			switch contractMethod {
			case "makeFromOntProof":
				Key := states[4].(string)
				TokenAddress := states[5].(string)
				Contract := notify.ContractAddress
				Value := states[6].(string)
				TChain := uint32(states[2].(float64))
				fmt.Printf("from ont, key: %s, token address: %s, contract: %s, value: %s, tchain: %d\n", Key, TokenAddress, Contract, Value, TChain)
				ParserOntCrossChainValue(TChain, Value)
			case "verifyToOntProof":
				FChain := uint32(states[3].(float64))
				Contract := notify.ContractAddress
				RTxHash := (states[1].(string))
				TokenAddress := states[5].(string)
				fmt.Printf("to ont, FChain: %d, Contract: %s, RTxhash: %s, token address: %s\n", FChain, Contract, RTxHash, TokenAddress)
			}
		}
	}
}


func TestCrossChainEvent_ONT2ETH(t *testing.T) {
	height := uint32(12575504)
	sdk := NewClient()
	events, _ := sdk.GetSmartContractEventByBlock(height)
	fmt.Printf("ontology, block height: %d, events num: %d\n", height, len(events))
	for _, event := range events {
		fmt.Printf("saveOntCrossTxsByHeight tx hash: %s, state:%d, gas: %d\n", event.TxHash, event.State, event.GasConsumed)
		for _, notify := range event.Notify {
			{
				states := notify.States.([]interface{})
				contractMethod, _ := states[0].(string)
				xxxx, _ := hex.DecodeString(contractMethod)
				if string(xxxx) == "lock" {
					sourceAssetAddress, _ := states[1].(string)
					tochainid, _ := states[2].(uint64)
					targetAssetAddress, _ := states[3].(string)
					fromAddress, _ := states[4].(string)
					toAddress, _ := states[5].(string)
					amount, _ := states[6].(uint64)
					fmt.Printf("source asset address: %s, tochainid: %d, targetassetaddress: %s, fromaddress: %s, toaddress: %s, amount:%d\n",
						sourceAssetAddress, tochainid, targetAssetAddress, fromAddress, toAddress, amount)
				}
			}

			if notify.ContractAddress != CrossChainManagerContractAddress.ToHexString() {
				continue
			}

			states := notify.States.([]interface{})
			contractMethod, _ := states[0].(string)
			fmt.Printf("contract method: %s, ", contractMethod)
			switch contractMethod {
			case "makeFromOntProof":
				Key := states[4].(string)
				TokenAddress := states[5].(string)
				Contract := notify.ContractAddress
				Value := states[6].(string)
				TChain := uint32(states[2].(float64))
				fmt.Printf("from ont, key: %s, token address: %s, contract: %s, value: %s, tchain: %d\n", Key, TokenAddress, Contract, Value, TChain)
				ParserOntCrossChainValue(TChain, Value)
			case "verifyToOntProof":
				FChain := uint32(states[3].(float64))
				Contract := notify.ContractAddress
				RTxHash := (states[1].(string))
				TokenAddress := states[5].(string)
				fmt.Printf("to ont, FChain: %d, Contract: %s, RTxhash: %s, token address: %s\n", FChain, Contract, RTxHash, TokenAddress)
			}
		}
	}
}


func TestCrossChainEvent_ONT2BTC(t *testing.T) {
	height := uint32(12579198)
	sdk := NewClient()
	events, _ := sdk.GetSmartContractEventByBlock(height)
	fmt.Printf("ontology, block height: %d, events num: %d\n", height, len(events))
	for _, event := range events {
		fmt.Printf("saveOntCrossTxsByHeight tx hash: %s, state:%d, gas: %d\n", event.TxHash, event.State, event.GasConsumed)
		for _, notify := range event.Notify {
			{
				states := notify.States.([]interface{})
				contractMethod, _ := states[0].(string)
				xxxx, _ := hex.DecodeString(contractMethod)
				if string(xxxx) == "lock" {
					sourceAssetAddress, _ := states[1].(string)
					tochainid, _ := states[2].(uint64)
					targetAssetAddress, _ := states[3].(string)
					fromAddress, _ := states[4].(string)
					toAddress, _ := states[5].(string)
					amount, _ := states[6].(uint64)
					fmt.Printf("source asset address: %s, tochainid: %d, targetassetaddress: %s, fromaddress: %s, toaddress: %s, amount:%d\n",
						sourceAssetAddress, tochainid, targetAssetAddress, fromAddress, toAddress, amount)
				}
			}

			if notify.ContractAddress != CrossChainManagerContractAddress.ToHexString() {
				continue
			}

			states := notify.States.([]interface{})
			contractMethod, _ := states[0].(string)
			fmt.Printf("contract method: %s, ", contractMethod)
			switch contractMethod {
			case "makeFromOntProof":
				Key := states[4].(string)
				TokenAddress := states[5].(string)
				Contract := notify.ContractAddress
				Value := states[6].(string)
				TChain := uint32(states[2].(float64))
				fmt.Printf("from ont, key: %s, token address: %s, contract: %s, value: %s, tchain: %d\n", Key, TokenAddress, Contract, Value, TChain)
				ParserOntCrossChainValue(TChain, Value)
			case "verifyToOntProof":
				FChain := uint32(states[3].(float64))
				Contract := notify.ContractAddress
				RTxHash := (states[1].(string))
				TokenAddress := states[5].(string)
				fmt.Printf("to ont, FChain: %d, Contract: %s, RTxhash: %s, token address: %s\n", FChain, Contract, RTxHash, TokenAddress)
			}
		}
	}
}
