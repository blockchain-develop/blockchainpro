package neo

import (
	"encoding/hex"
	"fmt"
	"github.com/joeqian10/neo-gogogo/rpc"
	"strconv"
	"testing"
)

func NewNeoClient() *rpc.RpcClient {
	url := "http://seed5.ngd.network:11332"
	client := rpc.NewClient(url)
	return client
}

func TestNeoCrossChainEvent_Scan(t *testing.T) {
	client := NewNeoClient()

	for i := 6024425;i < 6024426;i ++ {
		blockResp := client.GetBlockByIndex(uint32(i))
		block := blockResp.Result
		for _, tx := range block.Tx {
			if tx.Type != "InvocationTransaction" {
				continue
			}
			fmt.Printf("txid: %s\n", tx.Txid)
			fmt.Printf("txid: %s\n", tx.Txid[2:])
			logResp := client.GetApplicationLog(tx.Txid)

			if logResp.ErrorResponse.Error.Message != "" {
				fmt.Printf("GetApplicationLog err: %s\n", logResp.ErrorResponse.Error.Message)
			}

			appLog := logResp.Result
			for _, item := range appLog.Executions {
				fmt.Printf("saveNeoCrossTxsByHeight height: %d, tx contract: %s, gas: %s\n",i, item.Contract, item.GasConsumed)
				for _, notify := range item.Notifications {
					value := notify.State.Value
					if len(value) <= 0 {
						continue
					}
					method := value[0].Value
					xx, _ := hex.DecodeString(method)
					method = string(xx)

					fmt.Printf("notify, txhash : %s, contract: %s, method: %s, values: %v\n", tx.Txid, notify.Contract, method, value)
					if method == "CrossChainLockEvent" {
						fmt.Printf("xx: %s, from address: %s, contract address: %s, to chainid: %s, key: %s, param: %s\n",
							value[0].Value, value[1].Value, value[2].Value, value[3].Value, value[4].Value, value[5].Value)
						//parseNotifyData(notify.State.Value[6].Value)
					} else if method == "LockEvent" {
						fmt.Printf("xx: %s, from asset: %s, from address: %s, to chainid: %s, to asset: %s, to address: %s, amount: %s\n",
							value[0].Value, value[1].Value, value[2].Value, value[3].Value, value[4].Value, value[5].Value, value[6].Value)
						amount, _ := strconv.ParseUint(value[6].Value, 16, 32)
						fmt.Printf("==================================== %d\n", amount)
					}
				}
			}
		}
	}
}

func TestNeoCrossChainEvent(t *testing.T) {
	client := NewNeoClient()
	txhash := "b16a91d43cefe6490de6bcdd2b1b9741a455c0e4ab88717dae3c212787d04f35"
	//txhash := "21c6ae12471611d06682f47863f6771acefc0edffd9a4a5eb3fe8ca2c57c72ef"
	logResp := client.GetApplicationLog(txhash)

	if logResp.ErrorResponse.Error.Message != "" {
		fmt.Printf("GetApplicationLog err: %s\n", logResp.ErrorResponse.Error.Message)
	}

	appLog := logResp.Result
	for _, item := range appLog.Executions {
		fmt.Printf("saveNeoCrossTxsByHeight height: %d tx contract: %s, gas: %s\n", 1, item.Contract, item.GasConsumed)
		for _, notify := range item.Notifications {
			value := notify.State.Value
			method := value[0].Value
			xx, _ := hex.DecodeString(method)
			method = string(xx)
			fmt.Printf("notify, contract: %s, method: %s, values: %v\n", notify.Contract, method, value)
			if method == "CrossChainLockEvent" {
				fmt.Printf("xx: %s, from address: %s, contract address: %s, to chainid: %s, key: %s, param: %s\n",
					value[0].Value, value[1].Value, value[2].Value, value[3].Value, value[4].Value, value[5].Value)
				//parseNotifyData(notify.State.Value[6].Value)
			} else if method == "Lock" {
				fmt.Printf("xx: %s, from asset: %s, from address: %s, to chainid: %s, to asset: %s, to address: %s, amount: %s\n",
					value[0].Value, value[1].Value, value[2].Value, value[3].Value, value[4].Value, value[5].Value, value[6].Value)
			}
		}
	}
}

func TestNeoCrossChainEvent1(t *testing.T) {
	client := NewNeoClient()
	txhash := "f096666d0d1ffbe1ef039a676a31fd92b05a1516ea184c2a967a4b8074fbaffd"
	logResp := client.GetApplicationLog(txhash)

	if logResp.ErrorResponse.Error.Message != "" {
		fmt.Printf("GetApplicationLog err: %s\n", logResp.ErrorResponse.Error.Message)
	}

	appLog := logResp.Result
	for _, item := range appLog.Executions {
		fmt.Printf("saveNeoCrossTxsByHeight tx contract: %s, gas: %s\n", item.Contract, item.GasConsumed)
		for _, notify := range item.Notifications {
			value := notify.State.Value
			method := value[0].Value
			xx, _ := hex.DecodeString(method)
			method = string(xx)
			fmt.Printf("notify, contract: %s, method: %s, values: %v\n", notify.Contract, method, value)
			if method == "CrossChainUnlockEvent" {
				fmt.Printf("txhash: %s, %s, contract address: %s, to chainid: %s, to contarct: %s\n",
					notify.State.Value[1].Value, notify.State.Value[2].Value, notify.State.Value[3].Value, notify.State.Value[4].Value, notify.State.Value[5].Value)
				//parseNotifyData(notify.State.Value[6].Value)
			} else if method == "VerifyAndExecuteTxEvent" {
				value := notify.State.Value
				fmt.Printf("rtx hash: %s, token address: %s\n", value[3].Value, value[4].Value)
			}
		}
	}
}
