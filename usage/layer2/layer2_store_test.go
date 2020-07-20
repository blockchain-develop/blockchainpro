package layer2

import (
	"encoding/hex"
	"fmt"
	"github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/core/types"
	"github.com/tendermint/iavl"
	"testing"
)

func TestGetProof(t *testing.T) {
	sdk := ontology_go_sdk.NewOntologySdk()
	sdk.NewRpcClient().SetAddress("http://127.0.0.1:20336")

	key_str := "040dac0b6a91ac2fd5203ff2c5156fa4b4f9dc3902"
	key, _ := hex.DecodeString(key_str)
	proof, err := sdk.GetStoreProof("xxxx", key)
	if err  != nil {
		panic(err)
	}

	fmt.Printf("value: %s, proof: %s\n", proof.Value, proof.Proof)

	newProof, _ := hex.DecodeString(proof.Proof)

	source := common.NewZeroCopySource(newProof)
	xx := new(types.StoreProof)
	err = xx.Deserialization(source)
	if err != nil {
		panic(err)
	}

	root_str := ""
	root, _ := hex.DecodeString(root_str)
	bb := iavl.RangeProof(*xx)
	err = bb.Verify(root)
	if err != nil {
		panic(err)
	}
}
