package neo

import (
	"fmt"
	"testing"
)

func TestNodeStable(t *testing.T) {
	start := uint32(0)
	end := uint32(0)
	client := NewNeoClient()
	for i := start;i < end;i ++ {
		block := client.GetBlockByIndex(i)
		if block.ErrorResponse.Error.Message != "" {
			fmt.Printf("get block err: %s, height: %d\n", block.ErrorResponse.Error.Message, i)
			continue
		}

		for _, tx := range block.Result.Tx {
			if tx.Type != "InvocationTransaction" {
				continue
			}
			log := client.GetApplicationLog(tx.Txid)
			if log.ErrorResponse.Error.Message != "" {
				fmt.Printf("get application log err: %s, height: %d\n", block.ErrorResponse.Error.Message, i)
				continue
			}
		}
	}
}
