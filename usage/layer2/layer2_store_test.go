package layer2

import (
	"encoding/hex"
	"fmt"
	"github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/core/types"
	"testing"
)

func TestGetProof(t *testing.T) {
	sdk := ontology_go_sdk.NewOntologySdk()
	sdk.NewRpcClient().SetAddress("127.0.0.1:30334")

	key_str := ""
	key, _ := hex.DecodeString(key_str)
	proof, err := sdk.GetStoreProof("", key)
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
	err = xx.Verify(root)
	if err != nil {
		panic(err)
	}
}
