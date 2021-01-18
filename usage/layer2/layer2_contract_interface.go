package layer2

import (
	"encoding/hex"
	"fmt"
	"github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
)

func ontologyDeposit(ontsdk *ontology_go_sdk.OntologySdk, payer *ontology_go_sdk.Account, contract common.Address, token []byte, amount uint64) (common.Uint256, error) {
	tx, err := ontsdk.NeoVM.NewNeoVMInvokeTransaction(2500, 400000, contract, []interface{}{"deposit", []interface{}{
		payer.Address, amount, token}})
	if err != nil {
		return common.UINT256_EMPTY, fmt.Errorf("new transaction failed!")
	}
	ontsdk.SetPayer(tx, payer.Address)
	err = ontsdk.SignToTransaction(tx, payer)
	if err != nil {
		return common.UINT256_EMPTY, fmt.Errorf("SignToTransaction failed!")
	}
	txHash, err := ontsdk.SendTransaction(tx)
	if err != nil {
		return common.UINT256_EMPTY, fmt.Errorf("send transaction err: %s", err.Error())
	}
	return txHash, nil
}

func updateLayer2State(ontsdk *ontology_go_sdk.OntologySdk, payer *ontology_go_sdk.Account, contract common.Address) (common.Uint256, error) {
	depositids := make([]int, 0)
	for i := 0;i < 2;i ++ {
		depositids = append(depositids, 3 + i)
	}
	withdrawAmounts := make([]uint64, 0)
	toAddresses := make([]common.Address, 0)
	assetAddress := make([][]byte, 0)
	for i := 0;i < 1;i ++ {
		withdrawAmounts = append(withdrawAmounts, 3)
		toAddress, _ := common.AddressFromBase58("AMUGPqbVJ3TG6pe7xRpxxaeh4ai4fu9ahc")
		toAddresses = append(toAddresses,toAddress)
		tokenAddress, _ := hex.DecodeString("0000000000000000000000000000000000000002")
		assetAddress = append(assetAddress, tokenAddress)
	}
	tx, err := ontsdk.NeoVM.NewNeoVMInvokeTransaction(500, 40000, contract, []interface{}{"updateState", []interface{}{
		"0000000000000000000000000000000000000000000000000000000000000000", 6, "1",
		depositids, withdrawAmounts,toAddresses,assetAddress}})
	if err != nil {
		return common.UINT256_EMPTY, fmt.Errorf("new transaction failed!")
	}
	ontsdk.SetPayer(tx, payer.Address)
	err = ontsdk.SignToTransaction(tx, payer)
	if err != nil {
		return common.UINT256_EMPTY, fmt.Errorf("SignToTransaction failed!")
	}
	txHash, err := ontsdk.SendTransaction(tx)
	if err != nil {
		return common.UINT256_EMPTY, fmt.Errorf("send transaction err: %s", err.Error())
	}
	return txHash, nil
}

func updateLayer2StateBatch(ontsdk *ontology_go_sdk.OntologySdk, payer *ontology_go_sdk.Account, contract common.Address) (common.Uint256, error) {
	stateRootsBatch := make([]string, 0)
	heightsBatch := make([]uint32, 0)
	versionsBatch := make([]string, 0)
	depositidsBatch := make([][]int, 0)
	withdrawAmountsBatch := make([][]uint64, 0)
	toAddressesBatch := make([][]common.Address, 0)
	assetAddressBatch := make([][][]byte, 0)
	for i := 0;i < 10;i ++ {
		stateRootsBatch = append(stateRootsBatch, "0000000000000000000000000000000000000000000000000000000000000000")
		heightsBatch = append(heightsBatch, uint32(i + 1))
		versionsBatch = append(versionsBatch, "1")
		depositids := make([]int, 0)
		for j := 0; j < 2; j ++ {
			depositids = append(depositids, i * 2 + j)
		}
		depositidsBatch = append(depositidsBatch, depositids)
		withdrawAmounts := make([]uint64, 0)
		toAddresses := make([]common.Address, 0)
		assetAddress := make([][]byte, 0)
		for i := 0; i < 1; i ++ {
			withdrawAmounts = append(withdrawAmounts, 3)
			toAddress, _ := common.AddressFromBase58("AMUGPqbVJ3TG6pe7xRpxxaeh4ai4fu9ahc")
			toAddresses = append(toAddresses, toAddress)
			tokenAddress, _ := hex.DecodeString("0000000000000000000000000000000000000002")
			assetAddress = append(assetAddress, tokenAddress)
		}
		withdrawAmountsBatch = append(withdrawAmountsBatch, withdrawAmounts)
		toAddressesBatch = append(toAddressesBatch, toAddresses)
		assetAddressBatch = append(assetAddressBatch, assetAddress)
	}
	tx, err := ontsdk.NeoVM.NewNeoVMInvokeTransaction(2500, 4000000, contract, []interface{}{"updateStates", []interface{}{
		stateRootsBatch, heightsBatch, versionsBatch,
		depositidsBatch, withdrawAmountsBatch, toAddressesBatch, assetAddressBatch}})
	if err != nil {
		return common.UINT256_EMPTY, fmt.Errorf("new transaction failed!")
	}
	ontsdk.SetPayer(tx, payer.Address)
	err = ontsdk.SignToTransaction(tx, payer)
	if err != nil {
		return common.UINT256_EMPTY, fmt.Errorf("SignToTransaction failed!")
	}
	txHash, err := ontsdk.SendTransaction(tx)
	if err != nil {
		return common.UINT256_EMPTY, fmt.Errorf("send transaction err: %s", err.Error())
	}
	return txHash, nil
}

func GetCommitedLayer2StateByHeight(ontsdk *ontology_go_sdk.OntologySdk, contract common.Address, height uint32) ([]byte, uint32, error) {
	tx, err := ontsdk.NeoVM.NewNeoVMInvokeTransaction(0, 0, contract, []interface{}{"getStateRootByHeight", []interface{}{height}})
	if err != nil {
		return nil, 0, fmt.Errorf("new transaction failed!")
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
	xxx, _ := result.Result.ToString()
	fmt.Printf("xxxx : %s\n", xxx)
	return uint32(height.Uint64()), nil
}
