package layer2

import (
	"fmt"
	"github.com/ontio/ontology-crypto/keypair"
	"github.com/ontio/ontology-go-sdk"
	ontology_common "github.com/ontio/ontology/common"
	ontology_types "github.com/ontio/ontology/core/types"
)

const (
	LAYER2_CONTRACT = "0aad0408c6e4615b2f3f90c0c8c912649619a379"
	STORE_CONTRACT = "7680bc3227089ee6ac790be698e88bcd0be04609"
)

func newLayer2Sdk() *ontology_go_sdk.Layer2Sdk {
	// create alliance sdk
	layer2_sdk := ontology_go_sdk.NewLayer2Sdk()
	layer2_sdk.NewRpcClient().SetAddress("http://127.0.0.1:20336")
	//layer2_sdk.NewWebSocketClient().Connect("ws://localhost:40335")
	return layer2_sdk
}

func newOntologySdk() *ontology_go_sdk.OntologySdk {
	ontSdk := ontology_go_sdk.NewOntologySdk()
	ontSdk.NewRpcClient().SetAddress("http://polaris4.ont.io:20336")
	return ontSdk
}

func newLayer2OperatorAccount( layer2Sdk *ontology_go_sdk.Layer2Sdk) (*ontology_go_sdk.Account, error) {
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

func newLayer2UserAccount(ontsdk *ontology_go_sdk.Layer2Sdk) (*ontology_go_sdk.Account, error) {
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

func layer2DepositTransferOng(ontsdk *ontology_go_sdk.Layer2Sdk, payer *ontology_go_sdk.Account, to ontology_common.Address, amount uint64) (ontology_common.Uint256, error) {
	tx, err := ontsdk.Native.Ong.NewTransferTransaction(0, 20000, ontology_common.ADDRESS_EMPTY, to, amount)
	if err != nil {
		return ontology_common.UINT256_EMPTY, err
	}
	if payer != nil {
		ontsdk.SetPayer(tx, payer.Address)
		err = ontsdk.SignToTransaction(tx, payer)
		if err != nil {
			return ontology_common.UINT256_EMPTY, err
		}
	}
	return ontsdk.SendTransaction(tx)
}

func layer2WithdrawTransferOng(ontsdk *ontology_go_sdk.Layer2Sdk, payer *ontology_go_sdk.Account, from ontology_common.Address, amount uint64) (ontology_common.Uint256, error) {
	tx, err := ontsdk.Native.Ong.NewTransferTransaction(0, 20000, from, ontology_common.ADDRESS_EMPTY, amount)
	if err != nil {
		return ontology_common.UINT256_EMPTY, err
	}
	if payer != nil {
		ontsdk.SetPayer(tx, payer.Address)
		err = ontsdk.SignToTransaction(tx, payer)
		if err != nil {
			return ontology_common.UINT256_EMPTY, err
		}
	}
	return ontsdk.SendTransaction(tx)
}

func layer2TransferOng(ontsdk *ontology_go_sdk.Layer2Sdk, payer *ontology_go_sdk.Account, from ontology_common.Address, to ontology_common.Address, amount uint64) (ontology_common.Uint256, error) {
	tx, err := ontsdk.Native.Ong.NewTransferTransaction(0, 20000, from, to, amount)
	if err != nil {
		return ontology_common.UINT256_EMPTY, err
	}
	if payer != nil {
		ontsdk.SetPayer(tx, payer.Address)
		err = ontsdk.SignToTransaction(tx, payer)
		if err != nil {
			return ontology_common.UINT256_EMPTY, err
		}
	}
	return ontsdk.SendTransaction(tx)
}

func getLayer2OngBalance(layer2Sdk *ontology_go_sdk.Layer2Sdk, addr ontology_common.Address) uint64 {
	amount, err := layer2Sdk.Native.Ong.BalanceOf(addr)
	if err != nil {
		fmt.Printf("getLayer2OngBalance err: %s", err.Error())
	}
	return amount
}

func getLayer2OntBalance(layer2Sdk *ontology_go_sdk.Layer2Sdk, addr ontology_common.Address) uint64 {
	amount, err := layer2Sdk.Native.Ont.BalanceOf(addr)
	if err != nil {
		fmt.Printf("getLayer2OntBalance err: %s", err.Error())
	}
	return amount
}

func getOntologyOngBalance(ontSdk *ontology_go_sdk.OntologySdk, addr ontology_common.Address) uint64 {
	amount, err := ontSdk.Native.Ong.BalanceOf(addr)
	if err != nil {
		fmt.Printf("getOntologyOngBalance err: %s", err.Error())
	}
	return amount
}

func getOntologyOntBalance(ontSdk *ontology_go_sdk.OntologySdk, addr ontology_common.Address) uint64 {
	amount, err := ontSdk.Native.Ont.BalanceOf(addr)
	if err != nil {
		fmt.Printf("getOntologyOntBalance err: %s", err.Error())
	}
	return amount
}

func createLayer2Account() {
	// create alliance sdk
	layer2sdk := newLayer2Sdk()
	var wallet *ontology_go_sdk.Wallet
	var err error
	if !ontology_common.FileExisted("./wallet_layer2_new.dat") {
		wallet, err = layer2sdk.CreateWallet("./wallet_layer2_new.dat")
		if err != nil {
			return
		}
	} else {
		wallet, err = layer2sdk.OpenWallet("./wallet_layer2_new.dat")
		if err != nil {
			fmt.Errorf("createLayer2Account - wallet open error: %s", err.Error())
			return
		}
	}
	signer, err := wallet.GetDefaultAccount([]byte("1"))
	if err != nil || signer == nil {
		signer, err = wallet.NewDefaultSettingAccount([]byte("1"))
		if err != nil {
			fmt.Errorf("createLayer2Account - wallet password error")
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
