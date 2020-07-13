package neo

import (
	"encoding/hex"
	"fmt"
	"github.com/joeqian10/neo-gogogo/rpc"
	mccommon "github.com/ontio/multi-chain/common"
	"testing"
)

func NewNeoClient() *rpc.RpcClient {
	url := "http://seed1.ngd.network:20332"
	client := rpc.NewClient(url)
	return client
}

func parseNotifyData(data  string) {
	value, _ := hex.DecodeString(data)
	source := mccommon.NewZeroCopySource(value)
	txHash, _ := source.NextVarBytes()
	crosschainid, _ := source.NextVarBytes()
	token, _ := source.NextVarBytes()
	tochainid, _ := source.NextUint64()
	tocontract, _ := source.NextVarBytes()
	method, _ := source.NextVarBytes()
	args, _ := source.NextVarBytes()

	fmt.Printf("txhash: %s %s, token: %s, to chainid: %d, to contract: %s, mthod: %s\n",
		hex.EncodeToString(txHash), hex.EncodeToString(crosschainid), hex.EncodeToString(token), tochainid,
		hex.EncodeToString(tocontract), hex.EncodeToString(method))
	ParseNeoCrossTransfer(tochainid, args)
}

func ParseNeoCrossTransfer(toChainId uint64, args []byte) {
	source := mccommon.NewZeroCopySource(args)
	assethash, _ := source.NextVarBytes()
	toaddress, _ := source.NextVarBytes()
	amount, _ := source.NextUint64()
	fmt.Printf("asset hash: %s, to address: %s, amount: %d\n",
		hex.EncodeToString(assethash), hex.EncodeToString(toaddress), amount)
}


func TestNeoCrossChainEvent_Scan(t *testing.T) {
	client := NewNeoClient()

	for i := 4468200;i < 4489361;i ++ {
		blockResp := client.GetBlockByIndex(uint32(i))
		block := blockResp.Result
		for _, tx := range block.Tx {
			if tx.Type != "InvocationTransaction" {
				continue
			}
			logResp := client.GetApplicationLog(tx.Txid)

			if logResp.ErrorResponse.Error.Message != "" {
				fmt.Printf("GetApplicationLog err: %s\n", logResp.ErrorResponse.Error.Message)
			}

			appLog := logResp.Result
			for _, item := range appLog.Executions {
				fmt.Printf("saveNeoCrossTxsByHeight height: %d, tx contract: %s, gas: %s\n",i, item.Contract, item.GasConsumed)
				for _, notify := range item.Notifications {
					value := notify.State.Value
					method := value[0].Value
					xx, _ := hex.DecodeString(method)
					method = string(xx)
					fmt.Printf("notify, txhash : %s, contract: %s, method: %s, values: %v\n", tx.Txid, notify.Contract, method, value)
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
	}
}

func TestNeoCrossChainEvent(t *testing.T) {
	client := NewNeoClient()
	txhash := "0b66417146be074489a3556a1093c8e500d11584383c7c25bfc0d4aed7785b7d"
	//txhash := "21c6ae12471611d06682f47863f6771acefc0edffd9a4a5eb3fe8ca2c57c72ef"
	logResp := client.GetApplicationLog(txhash)

	if logResp.ErrorResponse.Error.Message != "" {
		fmt.Printf("GetApplicationLog err: %s\n", logResp.ErrorResponse.Error.Message)
	}

	appLog := logResp.Result
	for _, item := range appLog.Executions {
		fmt.Printf("saveNeoCrossTxsByHeight height: %d tx contract: %s, gas: %s\n", item.Contract, item.GasConsumed)
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
				parseNotifyData(notify.State.Value[6].Value)
			} else if method == "VerifyAndExecuteTxEvent" {
				value := notify.State.Value
				fmt.Printf("rtx hash: %s, token address: %s\n", value[3].Value, value[4].Value)
			}
		}
	}
}
