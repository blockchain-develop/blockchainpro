package layer2

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/core/types"
	"github.com/tendermint/iavl"
	"testing"
	"time"
)

func TestGetProof(t *testing.T) {
	sdk := ontology_go_sdk.NewOntologySdk(ontology_go_sdk.LAYER2_SDK)
	sdk.NewRpcClient().SetAddress("http://127.0.0.1:20336")

	key_str := "040dac0b6a91ac2fd5203ff2c5156fa4b4f9dc3902"
	key, _ := hex.DecodeString(key_str)
	store, err := sdk.GetStoreProof("xxxx", key)
	if err  != nil {
		panic(err)
	}

	fmt.Printf("value: %s, proof: %s, height: %d\n", store.Value, store.Proof, store.Height)

	proof_byte, _ := hex.DecodeString(store.Proof)

	source := common.NewZeroCopySource(proof_byte)
	proof := new(types.StoreProof)
	err = proof.Deserialization(source)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 30)

	block, err := sdk.GetBlockByHeight(store.Height)
	if err != nil {
		panic(err)
	}

	fmt.Printf("block height: %d, state root: %s\n", block.Header.Height, block.Header.StateRoot.ToHexString())
	fmt.Printf("block height: %d, state root: %s\n", block.Header.Height, hex.EncodeToString(block.Header.StateRoot.ToArray()))

	/*
	root_str := ""
	root, _ := hex.DecodeString(root_str)
	*/
	proof_iavl := iavl.RangeProof(*proof)
	proof_iavl_json, _ := json.Marshal(proof_iavl)
	fmt.Printf("proof json is: %s\n", string(proof_iavl_json))
	err = proof_iavl.Verify(block.Header.StateRoot.ToArray())
	if err != nil {
		panic(err)
	}

	fmt.Printf("verify successful!\n")
}


