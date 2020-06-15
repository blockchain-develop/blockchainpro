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

func TestNeoCrossChainEvent(t *testing.T) {
	client := NewNeoClient()
	txhash := "46fbe0e121168cd318a690b663d0781732fc4150cd3bbb15f6cc859dd6661904"
	logResp := client.GetApplicationLog(txhash)

	if logResp.ErrorResponse.Error.Message != "" {
		fmt.Printf("GetApplicationLog err: %s\n", logResp.ErrorResponse.Error.Message)
	}

	appLog := logResp.Result
	for _, item := range appLog.Executions {
		fmt.Printf("saveNeoCrossTxsByHeight tx contract: %s, gas: %s\n", item.Contract, item.GasConsumed)
		for _, notify := range item.Notifications {
			fmt.Printf("notify contract: %s\n", notify.Contract)
			method := notify.State.Value[0].Value
			xx, _ := hex.DecodeString(method)
			method = string(xx)
			fmt.Printf("notify, method: %s\n", method)
			if method == "create_cross_tx_success" {
				parseNotifyData(notify.State.Value[4].Value)
			} else if method == "verifyToNeoProof" {
				fmt.Printf("rtx hash: %s, token address: %s\n", notify.State.Value[2].Value, notify.State.Value[4].Value)
			}
		}
	}
}

func TestNeoCrossChainEvent1(t *testing.T) {
	client := NewNeoClient()
	txhash := "4cb48acce66c48a451f516dbcf8d525a84d661dbe1610be882a97e29c87b46fa"
	logResp := client.GetApplicationLog(txhash)

	if logResp.ErrorResponse.Error.Message != "" {
		fmt.Printf("GetApplicationLog err: %s\n", logResp.ErrorResponse.Error.Message)
	}

	appLog := logResp.Result
	for _, item := range appLog.Executions {
		fmt.Printf("saveNeoCrossTxsByHeight tx contract: %s, gas: %s\n", item.Contract, item.GasConsumed)
		for _, notify := range item.Notifications {
			fmt.Printf("notify contract: %s\n", notify.Contract)
			method := notify.State.Value[0].Value
			xx, _ := hex.DecodeString(method)
			method = string(xx)
			fmt.Printf("notify, method: %s\n", method)
			if method == "transfer" {
				value := notify.State.Value
				fmt.Printf("from address: %s, to address: %s\n",value[1], value[2])
			}
			if notify.Contract[2:] != "02d9290db5ff0ce5242727fbdbdf01aacc6656f5" {
				continue
			}
			if method == "CrossChainEvent" {
				fmt.Printf("txhash: %s, %s, contract address: %s, to chainid: %s, to contarct: %s\n",
					notify.State.Value[1].Value, notify.State.Value[2].Value, notify.State.Value[3].Value, notify.State.Value[4].Value, notify.State.Value[5].Value)
				parseNotifyData(notify.State.Value[6].Value)
			} else if method == "verifyToNeoProof" {
				fmt.Printf("rtx hash: %s, token address: %s\n", notify.State.Value[2].Value, notify.State.Value[4].Value)
			}
		}
	}
}
