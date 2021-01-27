package ontology

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"
	"github.com/ontio/ontology/common"
)

const (
	_ont_crosschainlock = "makeFromOntProof"
	_ont_crosschainunlock = "verifyToOntProof"
	_ont_lock               = "lock"
	_ont_unlock               = "unlock"
)

func parseOntolofyMethod(v string) string {
	xx, _ := hex.DecodeString(v)
	return string(xx)
}

func Hash2Address(value string) string {
	value = HexStringReverse(value)
	addr, _ := common.AddressFromHexString(value)
	return addr.ToBase58()
}

func HexReverse(arr []byte) []byte {
	l := len(arr)
	x := make([]byte, 0)
	for i := l - 1; i >= 0; i-- {
		x = append(x, arr[i])
	}
	return x
}

func HexStringReverse(value string) string {
	aa, _ := hex.DecodeString(value)
	bb := HexReverse(aa)
	return hex.EncodeToString(bb)
}



func TestCrossChainNotify(t *testing.T) {
	sdk := newOntologySdk()
	localHeight := uint32(13965800)
	events, err := sdk.GetSmartContractEventByBlock(localHeight)
	if err != nil {
		panic(err)
	}
	for _, event := range events {
		for _, notify := range event.Notify {
			if notify.ContractAddress != "0900000000000000000000000000000000000000" {
				continue
			}
			states := notify.States.([]interface{})
			contractMethod, _ := states[0].(string)
			switch contractMethod {
			case _ont_crosschainlock:
				fmt.Printf("from chain: %s, txhash: %s\n", "ontology", event.TxHash)
				for _, notifynew := range event.Notify {
					statesnew := notifynew.States.([]interface{})
					method, ok := statesnew[0].(string)
					if !ok {
						continue
					}
					contractMethodNew := parseOntolofyMethod(method)
					if contractMethodNew == _ont_lock {
						//
						From := Hash2Address(statesnew[2].(string))
						To := Hash2Address(states[5].(string))
						Asset := HexStringReverse(statesnew[1].(string))
						amount, _ := strconv.ParseUint(HexStringReverse(statesnew[6].(string)), 16, 64)
						Amount := uint64(amount)
						toChain, _ := strconv.ParseUint(statesnew[3].(string), 16, 32)
						ToChain := uint32(toChain)
						ToAsset := statesnew[4].(string)
						ToUser := Hash2Address(statesnew[5].(string))
						fmt.Printf("From: %s, To: %s, Asset: %s, Amount: %d, ToChain: %d, ToAsset: %s, ToUser: %s\n", From, To, Asset, Amount, ToChain, ToAsset, ToUser)
					}
				}

				TChain := uint32(states[2].(float64))
				Contract := HexStringReverse(states[5].(string))
				Key := states[4].(string)
				Param := states[6].(string)
				fmt.Printf("TChain: %d, Contract: %s, Key: %s, Param:%s\n", TChain, Contract, Key, Param)
			case _ont_crosschainunlock:
				fmt.Printf("to chain: %s, txhash: %s\n", "ontology", event.TxHash)
				for _, notifynew := range event.Notify {
					statesnew := notifynew.States.([]interface{})
					method, ok := statesnew[0].(string)
					if !ok {
						continue
					}
					contractMethodNew := parseOntolofyMethod(method)
					if contractMethodNew == _ont_unlock {
						//
						From := Hash2Address(states[5].(string))
						To := Hash2Address(statesnew[2].(string))
						Asset := HexStringReverse(statesnew[1].(string))
						amount, _ := strconv.ParseUint(HexStringReverse(statesnew[3].(string)), 16, 64)
						Amount := amount
						fmt.Printf("From: %s, To: %s, Asset: %s, Amount: %d\n", From, To, Asset, Amount)
					}
				}
				FChain := uint32(states[3].(float64))
				Contract := HexStringReverse(states[5].(string))
				RTxHash := HexStringReverse(states[1].(string))
				fmt.Printf("FChain: %d, Contract: %s, RTxHash: %s\n", FChain, Contract, RTxHash)
			}
		}
	}
}