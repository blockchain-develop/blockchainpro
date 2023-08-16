package sui

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/serde"
	"testing"
)

func TestTransactionData_newTransferSui(t *testing.T) {
	recipient := "115d569aa1e7bd2527f48efa1eecf03df3bc3105bedcd729250b6384eb3b9d97"
	amount := uint64(100000000000)
	pt := ProgrammableTransaction{}
	//
	var amountBytes [8]uint8
	binary.LittleEndian.PutUint64(amountBytes[:], amount)
	var amountArg CallArg__Pure = amountBytes[:]
	recipientBytes, _ := hex.DecodeString(recipient)
	var recipientArg CallArg__Pure = recipientBytes
	inputs := make([]CallArg, 0)
	inputs = append(inputs, &amountArg)
	inputs = append(inputs, &recipientArg)
	//
	var gasCoin Argument__GasCoin
	var splitCoin Argument__Input = 0
	var splitCoinCommand Command__SplitCoins
	splitCoinCommand.Field0 = &gasCoin
	splitCoinCommand.Field1 = []Argument{&splitCoin}
	//
	var transferObjectsCommand Command__TransferObjects
	var srcObjects Argument__NestedResult
	srcObjects.Field0 = 0
	srcObjects.Field1 = 0
	transferObjectsCommand.Field0 = []Argument{&srcObjects}
	var dest Argument__Input = 1
	transferObjectsCommand.Field1 = &dest
	//
	se := serde.NewBinarySerializer(100)
	err := pt.Serialize(se)
	if err != nil {
		panic(err)
	}
}
