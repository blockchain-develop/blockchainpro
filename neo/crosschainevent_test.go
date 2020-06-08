package neo

import (
	"fmt"
	"testing"
	"github.com/joeqian10/neo-gogogo/rpc"

)

func NewNeoClient() *rpc.RpcClient {
	url := ""
	client := rpc.NewClient(url)
	return client
}

func parseNotifyData(data  string) {

}

func TestNeoCrossChainEvent(t *testing.T) {
	client := NewNeoClient()
	txhash := ""
	logResp := client.GetApplicationLog(txhash)

	if logResp.ErrorResponse.Error.Message != "" {
		fmt.Printf("GetApplicationLog err: %s\n", logResp.ErrorResponse.Error.Message)
	}

	appLog := logResp.Result
	for _, item := range appLog.Executions {
		fmt.Printf("saveNeoCrossTxsByHeight tx contract: %s, gas: %s", item.Contract, item.GasConsumed)
		for _, notify := range item.Notifications {
			fmt.Printf("notify contract: %s\n", notify.Contract)
			fmt.Printf("notify state 0: %s\n", notify.State.Value[0].Value)
			method := notify.State.Value[0].Value
			if method == "create_cross_tx_success" {
				parseNotifyData(notify.State.Value[4].Value)
			} else if method == "verifyToNeoProof" {
				fmt.Printf("rtx hash: %s, token address: %s\n", notify.State.Value[2].Value, notify.State.Value[4].Value)
			}
		}
	}
}
