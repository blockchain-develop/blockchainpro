package ontology

import (
	"fmt"
	"github.com/ontio/ontology-crypto/keypair"
	"github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/utils"
	ontology_common "github.com/ontio/ontology/common"
	ontology_types "github.com/ontio/ontology/core/types"
)

const (
	LAYER2_CONTRACT = "0aad0408c6e4615b2f3f90c0c8c912649619a379"
	STORE_CONTRACT = "7680bc3227089ee6ac790be698e88bcd0be04609"
)

func newOntologySdk() *ontology_go_sdk.OntologySdk {
	ontSdk := ontology_go_sdk.NewOntologySdk(utils.ONTOLOGY_SDK)
	ontSdk.NewRpcClient(utils.ONTOLOGY_SDK).SetAddress("http://polaris5.ont.io:20336")
	return ontSdk
}

func newOntologyOperatorAccount(ontsdk *ontology_go_sdk.OntologySdk) (*ontology_go_sdk.Account, error) {
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

func newOntologyUserAccount(ontsdk *ontology_go_sdk.OntologySdk) (*ontology_go_sdk.Account, error) {
	// AScExXzLbkZV32tDFdV7Uoq7ZhCT1bRCGp
	privateKey, err := keypair.WIF2Key([]byte("KyxsqZ45MCx3t2UbuG9P8h96TzyrzTXGRQnfs9nZKFx6YkjTfHqb"))
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

func ontologyContractInvoke(ontsdk *ontology_go_sdk.OntologySdk, payer *ontology_go_sdk.Account, contract ontology_common.Address, token []byte, amount uint64) (ontology_common.Uint256, error) {
	tx, err := ontsdk.NeoVM.NewNeoVMInvokeTransaction(2500, 400000, contract, []interface{}{"deposit", []interface{}{
		payer.Address, amount, token}})
	if err != nil {
		fmt.Printf("new transaction failed!")
	}
	ontsdk.SetPayer(tx, payer.Address)
	err = ontsdk.SignToTransaction(tx, payer)
	if err != nil {
		fmt.Printf("SignToTransaction failed!")
	}
	txHash, err := ontsdk.SendTransaction(tx)
	if err != nil {
		fmt.Printf("SignToTransaction failed! err: %s", err.Error())
	}
	return txHash, nil
}

func createOntologyAccount() {
	// create alliance sdk
	ontsdk := newOntologySdk()
	var wallet *ontology_go_sdk.Wallet
	var err error
	if !ontology_common.FileExisted("./wallet_ontology_new.dat") {
		wallet, err = ontsdk.CreateWallet("./wallet_ontology_new.dat")
		if err != nil {
			return
		}
	} else {
		wallet, err = ontsdk.OpenWallet("./wallet_ontology_new.dat")
		if err != nil {
			fmt.Errorf("createOntologyAccount - wallet open error: %s", err.Error())
			return
		}
	}
	signer, err := wallet.GetDefaultAccount([]byte("1"))
	if err != nil || signer == nil {
		signer, err = wallet.NewDefaultSettingAccount([]byte("1"))
		if err != nil {
			fmt.Errorf("createOntologyAccount - wallet password error")
			return
		}
		err = wallet.Save()
		if err != nil {
			return
		}
	}
	pri_key, _ := keypair.Key2WIF(signer.PrivateKey)
	addr := signer.Address.ToBase58()
	fmt.Printf("private key: %s, address: %s %s\n", string(pri_key), addr, signer.Address.ToHexString())
}

func getOntologyOngBalance(ontSdk *ontology_go_sdk.OntologySdk, addr ontology_common.Address) uint64 {
	amount, err := ontSdk.Native.Ong.BalanceOf(addr)
	if err != nil {
		fmt.Printf("getOntologyBalance err: %s", err.Error())
	}
	return amount
}
