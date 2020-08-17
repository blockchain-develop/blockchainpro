package layer2

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common"
	"github.com/tendermint/iavl"
	"testing"
	"time"
)


func TestDeployContract(t *testing.T) {
	layer2_sdk := newLayer2Sdk()
	account_operator, _ := newLayer2OperatorAccount(layer2_sdk)
	code := "52c56b0568656c6c6f6a00527ac46c59c56b6a00527ac46a51527ac46a52527ac46a51c30548656c6c6f7d9c7c756422006a52c300c36a53527ac46a53c3516a00c3065400000000006e6c7566620300006c756657c56b6a00527ac46a51527ac46a52527ac46203006a52c36a00c300c3681953797374656d2e53746f726167652e476574436f6e74657874681253797374656d2e53746f726167652e5075746a52c3681553797374656d2e52756e74696d652e4e6f74696679516c7566"
	hash, err := layer2_sdk.NeoVM.DeployNeoVMSmartContract_Layer2(0, 20000000, account_operator, true,
		code, "hello", "1.0.0", "hello", "hello", "hello")

	if err != nil {
		panic(err)
	}
	fmt.Printf("deploy hash: %s\n", hash.ToHexString())
}

func TestInvokeContract(t *testing.T) {
	layer2_sdk := newLayer2Sdk()
	account_operator, _ := newLayer2OperatorAccount(layer2_sdk)
	contractAddress, _ := common.AddressFromHexString(STORE_CONTRACT)
	tx, err := layer2_sdk.NeoVM.NewNeoVMInvokeTransaction(0, 200000, contractAddress, []interface{}{"Hello", []interface{}{"this is example"}})
	if err != nil {
		panic(err)
	}
	layer2_sdk.SetPayer(tx, account_operator.Address)
	err = layer2_sdk.SignToLayer2Transaction(tx, account_operator)
	if err != nil {
		panic(err)
	}

	txHash, err := layer2_sdk.SendTransaction(tx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("layer2 state commit transaction hash: %s", txHash.ToHexString())
}

func TestGetProof(t *testing.T) {
	sdk := newLayer2Sdk()
	key_str := "040dac0b6a91ac2fd5203ff2c5156fa4b4f9dc3902"
	key, _ := hex.DecodeString(key_str)
	store, err := sdk.GetStoreProof(key)
	if err  != nil {
		panic(err)
	}
	fmt.Printf("value: %s, proof: %s, height: %d\n", store.Value, store.Proof, store.Height)

	proof_byte, _ := hex.DecodeString(store.Proof)
	source := common.NewZeroCopySource(proof_byte)
	proof := new(utils.StoreProof)
	err = proof.Deserialization(source)
	if err != nil {
		panic(err)
	}

	curHeight, err := sdk.GetCurrentBlockHeight()
	if err != nil {
		panic(err)
	}
	for curHeight < store.Height {
		time.Sleep(time.Second * 1)
		curHeight, err = sdk.GetCurrentBlockHeight()
		if err != nil {
			panic(err)
		}
	}
	block, err := sdk.GetLayer2BlockByHeight(curHeight)
	if err != nil {
		panic(err)
	}

	fmt.Printf("block height: %d, state root: %s\n", block.Header.Height, block.Header.StateRoot.ToHexString())
	fmt.Printf("block height: %d, state root: %s\n", block.Header.Height, hex.EncodeToString(block.Header.StateRoot.ToArray()))

	proof_iavl := iavl.RangeProof(*proof)
	proof_iavl_json, _ := json.Marshal(proof_iavl)
	fmt.Printf("proof json is: %s\n", string(proof_iavl_json))
	err = proof_iavl.Verify(block.Header.StateRoot.ToArray())
	if err != nil {
		panic(err)
	}
	fmt.Printf("verify successful!\n")
}

func TestGetContractStoreProof(t *testing.T) {
	sdk := newLayer2Sdk()
	key, _ := sdk.GetStoreKey(STORE_CONTRACT, []byte("hello"))
	store, err := sdk.GetStoreProof(key)
	if err  != nil {
		panic(err)
	}
	fmt.Printf("key: %s\n", hex.EncodeToString(key))
	fmt.Printf("value: %s, proof: %s, height: %d\n", store.Value, store.Proof, store.Height)

	proof_byte, _ := hex.DecodeString(store.Proof)
	source := common.NewZeroCopySource(proof_byte)
	proof := new(utils.StoreProof)
	err = proof.Deserialization(source)
	if err != nil {
		panic(err)
	}

	curHeight, err := sdk.GetCurrentBlockHeight()
	if err != nil {
		panic(err)
	}
	for curHeight < store.Height {
		time.Sleep(time.Second * 1)
		curHeight, err = sdk.GetCurrentBlockHeight()
		if err != nil {
			panic(err)
		}
	}
	block, err := sdk.GetLayer2BlockByHeight(curHeight)
	if err != nil {
		panic(err)
	}

	fmt.Printf("block height: %d, state root: %s\n", block.Header.Height, block.Header.StateRoot.ToHexString())
	fmt.Printf("block height: %d, state root: %s\n", block.Header.Height, hex.EncodeToString(block.Header.StateRoot.ToArray()))

	proof_iavl := iavl.RangeProof(*proof)
	proof_iavl_json, _ := json.Marshal(proof_iavl)
	fmt.Printf("proof json is: %s\n", string(proof_iavl_json))
	err = proof_iavl.Verify(block.Header.StateRoot.ToArray())
	if err != nil {
		panic(err)
	}
	fmt.Printf("verify successful!\n")
}

func TestVerifyContractStore(t *testing.T) {
	sdk := newLayer2Sdk()
	key, _ := sdk.GetStoreKey(STORE_CONTRACT, []byte("hello"))
	store, err := sdk.GetStoreProof(key)
	if err  != nil {
		panic(err)
	}
	fmt.Printf("value: %s, proof: %s, height: %d\n", store.Value, store.Proof, store.Height)

	proof_byte, _ := hex.DecodeString(store.Proof)
	source := common.NewZeroCopySource(proof_byte)
	proof := new(utils.StoreProof)
	err = proof.Deserialization(source)
	if err != nil {
		panic(err)
	}

	ont_sdk := newOntologySdk()
	contractAddress, _ := common.AddressFromHexString(LAYER2_CONTRACT)
	curHeight, err := GetCommitedLayer2Height(ont_sdk, contractAddress)
	if err != nil {
		panic(err)
	}

	for curHeight < store.Height {
		time.Sleep(time.Second * 1)
		curHeight, err = GetCommitedLayer2Height(ont_sdk, contractAddress)
		if err != nil {
			panic(err)
		}
	}

	stateRoot, height, err := GetCommitedLayer2StateByHeight(ont_sdk, contractAddress, store.Height)
	if err != nil {
		panic(err)
	}
	fmt.Printf("state root: %s, height: %d\n", hex.EncodeToString(stateRoot), height)

	proof_iavl := iavl.RangeProof(*proof)
	proof_iavl_json, _ := json.Marshal(proof_iavl)
	fmt.Printf("proof json is: %s\n", string(proof_iavl_json))
	err = proof_iavl.Verify(stateRoot)
	if err != nil {
		panic(err)
	}
	value_bytes, _ := hex.DecodeString(store.Value)
	err = proof_iavl.VerifyItem(key, value_bytes)
	if err != nil {
		panic(err)
	}

	fmt.Printf("verify successful!\n")
}

func TestVerifyContractStore1(t *testing.T) {
	sdk := newLayer2Sdk()
	key, _ := sdk.GetStoreKey(STORE_CONTRACT, []byte("hello"))
	store, err := sdk.GetStoreProof(key)
	if err  != nil {
		panic(err)
	}
	fmt.Printf("value: %s, proof: %s, height: %d\n", store.Value, store.Proof, store.Height)

	ont_sdk := newOntologySdk()
	contractAddress, _ := common.AddressFromHexString(LAYER2_CONTRACT)
	curHeight, err := GetCommitedLayer2Height(ont_sdk, contractAddress)
	if err != nil {
		panic(err)
	}
	for curHeight < store.Height {
		time.Sleep(time.Second * 1)
		curHeight, err = GetCommitedLayer2Height(ont_sdk, contractAddress)
		if err != nil {
			panic(err)
		}
	}

	stateRoot, height, err := GetCommitedLayer2StateByHeight(ont_sdk, contractAddress, store.Height)
	if err != nil {
		panic(err)
	}
	fmt.Printf("state root: %s, height: %d\n", hex.EncodeToString(stateRoot), height)

	proof_byte, _ := hex.DecodeString(store.Proof)
	value_bytes, _ := hex.DecodeString(store.Value)
	result, err := sdk.VerifyStoreProof(key, value_bytes, proof_byte, stateRoot)
	if err != nil {
		panic(err)
	}
	if result {
		fmt.Printf("verify successful!\n")
	} else {
		fmt.Printf("verify failed!\n")
	}
}


