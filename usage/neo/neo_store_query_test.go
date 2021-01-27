package neo

import (
	"fmt"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/rpc/models"
	"testing"
)

func TestStore(t *testing.T) {
	client := NewNeoClient()
	//arg := models.InvokeStack{}
	arg := models.InvokeStack{}
	response := client.InvokeFunction("104057f879009326250ee1f5d60e2efd925024e6", "lockProxy", helper.ZeroScriptHashString, arg)
	if response.HasError() || response.Result.State == "FAULT" {
		panic(fmt.Errorf("[GetCurrentNeoChainSyncHeight] GetCurrentHeight error: %s", "Engine faulted! "+response.Error.Message))
	}
	var address string
	s := response.Result.Stack
	if s != nil && len(s) != 0 {
		s[0].Convert()
		address = (s[0].Value.(string))
	}
	fmt.Printf("address: %s\n", address)
}


func TestStore2(t *testing.T) {
	client := NewNeoClient()
	arg := models.InvokeStack{Type: "Integer", Value: 0}
	//arg := models.InvokeStack{}
	response := client.InvokeFunction("0xe1695b1314a1331e3935481620417ed835669407", "currentSyncHeight", helper.ZeroScriptHashString, arg)
	if response.HasError() || response.Result.State == "FAULT" {
		panic(fmt.Errorf("[GetCurrentNeoChainSyncHeight] GetCurrentHeight error: %s", "Engine faulted! "+response.Error.Message))
	}
	var height uint64
	var e error
	var b []byte
	s := response.Result.Stack
	if s != nil && len(s) != 0 {
		s[0].Convert()
		b = helper.HexToBytes(s[0].Value.(string))
	}
	if len(b) == 0 {
		height = 0
	} else {
		height = helper.BytesToUInt64(b)
		if e != nil {
			panic(fmt.Errorf("[GetCurrentNeoChainSyncHeight], ParseVarInt error: %s", e))
		}
		height++ // means the next block header needs to be synced
	}
	fmt.Printf("height: %d\n", height)
}
