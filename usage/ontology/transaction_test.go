package ontology

import (
	"fmt"
	"github.com/ontio/ontology-crypto/keypair"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	ontology_types "github.com/ontio/ontology/core/types"
	"testing"
)

func newOntologyed25519Account(ontsdk *ontology_go_sdk.OntologySdk) (*ontology_go_sdk.Account, error) {
	// AMUGPqbVJ3TG6pe7xRpxxaeh4ai4fu9ahc
	privateKey, err := keypair.WIF2Key([]byte("L5CKUdMTnHQNeBtCzdoEZ1hyRpaCsc7LaesZWvFhfpKbzQV1v7pk"))
	if err != nil {
		return nil, fmt.Errorf("decrypt privateKey error:%s", err)
	}
	pub := privateKey.Public()
	address := ontology_types.AddressFromPubKey(pub)
	fmt.Printf("address: %s\n", address.ToBase58())
	return &ontology_go_sdk.Account{
		PrivateKey: privateKey,
		PublicKey:  pub,
		Address:    address,
	}, nil
}

func TestTransferOng(t *testing.T) {
	ontSdk := newOntologySdk()
	fromAddr, err := common.AddressFromBase58("AHqqNcWJ75XRgz3RLhFE5nJApnAwPfZXNG")
	if err != nil {
		panic(err)
	}
	toAddr, err := common.AddressFromBase58("Ad9ggp73Gm8bshrMqse9PGbv5wsCB1diTS")
	if err != nil {
		panic(err)
	}
	ontologyTx, err := ontSdk.Native.Ong.NewTransferTransaction(2500, 30000, fromAddr, toAddr, 2000000000)
	if err != nil {
		panic(err)
	}
	ontSdk.SetPayer(ontologyTx, fromAddr)
	fakeAccount, err := newOntologyed25519Account(ontSdk)
	if err != nil {
		panic(err)
	}
	ontSdk.SignToTransaction(ontologyTx, fakeAccount)
}
