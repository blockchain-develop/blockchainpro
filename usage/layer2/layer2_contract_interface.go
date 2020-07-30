package layer2

import (
	"fmt"
	"github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
)

func GetCommitedLayer2StateByHeight(ontsdk *ontology_go_sdk.OntologySdk, contract common.Address, height uint32) ([]byte, uint32, error) {
	tx, err := ontsdk.NeoVM.NewNeoVMInvokeTransaction(0, 0, contract, []interface{}{"getStateRootByHeight", []interface{}{height}})
	if err != nil {
		fmt.Printf("new transaction failed!")
	}
	result, err := ontsdk.PreExecTransaction(tx)
	if err != nil {
		fmt.Printf("PreExecTransaction failed! err: %s", err.Error())
		return nil, 0, err
	}
	if result == nil {
		fmt.Printf("can not find the result")
		return nil, 0, fmt.Errorf("can not find state of heigh: %d", height)
	}
	tt, _ := result.Result.ToArray()
	if len(tt) != 3 {
		fmt.Printf("result is not right")
		return nil, 0, fmt.Errorf("result is not right, height: %d", height)
	}
	item0,_ := tt[0].ToString()
	item1,_ := tt[1].ToInteger()
	item2,_ := tt[2].ToInteger()
	fmt.Printf("item0: %s, item1: %d, item2: %d\n", item0, item1, item2)
	stateRoot, err := common.Uint256FromHexString(item0)
	if err != nil {
		return nil, 0, fmt.Errorf("state hash is not right, height: %d", height)
	}
	return stateRoot.ToArray(), uint32(item1.Uint64()), nil
}

func GetCommitedLayer2Height(ontsdk *ontology_go_sdk.OntologySdk, contract common.Address) (uint32, error) {
	tx, err := ontsdk.NeoVM.NewNeoVMInvokeTransaction(0, 0, contract, []interface{}{"getCurrentHeight", []interface{}{}})
	if err != nil {
		return 0, err
	}
	result, err := ontsdk.PreExecTransaction(tx)
	if err != nil {
		fmt.Printf("PreExecTransaction failed! err: %s", err.Error())
		return 0, err
	}
	if result == nil {
		fmt.Printf("can not find the result")
		return 0, fmt.Errorf("can not find current height!")
	}
	height, err := result.Result.ToInteger()
	if err != nil {
		return 0, fmt.Errorf("current height is not right!")
	}
	return uint32(height.Uint64()), nil
}
