package layer2

import (
	"encoding/hex"
	"fmt"
	"github.com/ontio/ontology-crypto/keypair"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	ontology_common "github.com/ontio/ontology/common"
	ontology_types "github.com/ontio/ontology/core/types"
	"testing"
	"time"
)

func newLayer2Sdk() *ontology_go_sdk.OntologySdk {
	// create alliance sdk
	layer2_sdk := ontology_go_sdk.NewOntologySdk(ontology_go_sdk.LAYER2_SDK)
	layer2_sdk.NewRpcClient().SetAddress("http://127.0.0.1:40336")
	//layer2_sdk.NewWebSocketClient().Connect("ws://localhost:40335")
	return layer2_sdk
}

func newOntologySdk() *ontology_go_sdk.OntologySdk {
	ontSdk := ontology_go_sdk.NewOntologySdk(ontology_go_sdk.ONTOLOFY_SDK)
	ontSdk.NewRpcClient().SetAddress("http://polaris5.ont.io:20336")
	return ontSdk
}

func newLayer2OperatorAccount( layer2Sdk *ontology_go_sdk.OntologySdk) (*ontology_go_sdk.Account, error) {
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

func newLayer2UserAccount(ontsdk *ontology_go_sdk.OntologySdk) (*ontology_go_sdk.Account, error) {
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

func layer2DepositTransfer(ontsdk *ontology_go_sdk.OntologySdk, gasPrice, gasLimit uint64, payer *ontology_go_sdk.Account, to ontology_common.Address, amount uint64) (ontology_common.Uint256, error) {
	tx, err := ontsdk.Native.Ong.NewTransferTransaction(gasPrice, gasLimit, ontology_common.ADDRESS_EMPTY, to, amount)
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

func layer2WithdrawTransfer(ontsdk *ontology_go_sdk.OntologySdk, gasPrice, gasLimit uint64, payer *ontology_go_sdk.Account, from ontology_common.Address, amount uint64) (ontology_common.Uint256, error) {
	tx, err := ontsdk.Native.Ong.NewTransferTransaction(gasPrice, gasLimit, from, ontology_common.ADDRESS_EMPTY, amount)
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

func layer2Transfer(ontsdk *ontology_go_sdk.OntologySdk, gasPrice, gasLimit uint64, payer *ontology_go_sdk.Account, from ontology_common.Address, to ontology_common.Address, amount uint64) (ontology_common.Uint256, error) {
	tx, err := ontsdk.Native.Ong.NewTransferTransaction(gasPrice, gasLimit, from, to, amount)
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

func getLayer2Balance(layer2Sdk *ontology_go_sdk.OntologySdk, addr ontology_common.Address) uint64 {
	amount, err := layer2Sdk.Native.Ong.BalanceOf(addr)
	if err != nil {
		fmt.Printf("getLayer2Balance err: %s", err.Error())
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

func TestLayer2Deposit(t *testing.T) {
	// create alliance sdk
	layer2_sdk := newLayer2Sdk()
	account_operator, _ := newLayer2OperatorAccount(layer2_sdk)
	account_user, _ := newLayer2UserAccount(layer2_sdk)
	//
	balance := getLayer2Balance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
	//
	var gasPrice uint64 = 0
	var gasLimit uint64 = 20000
	var amount uint64 = 10000000
	{
		txhash, err := layer2DepositTransfer(layer2_sdk, gasPrice, gasLimit, account_operator, account_user.Address, amount)
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
		} else {
			fmt.Printf("tx hash: %s\n", txhash.ToHexString())
		}
	}
	//
	balance = getLayer2Balance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
}

func TestLayer2Withdraw(t *testing.T) {
	// create alliance sdk
	layer2_sdk := newLayer2Sdk()
	account_user, _ := newLayer2UserAccount(layer2_sdk)
	//
	balance := getLayer2Balance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
	//
	var gasPrice uint64 = 0
	var gasLimit uint64 = 20000
	var amount uint64 = 2000000
	{
		txhash, err := layer2WithdrawTransfer(layer2_sdk, gasPrice, gasLimit, account_user, account_user.Address, amount)
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
		} else {
			fmt.Printf("tx hash: %s\n", txhash.ToHexString())
		}
	}
	//
	balance = getLayer2Balance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
}

func TestLayer2Transfer(t *testing.T) {
	// create alliance sdk
	layer2_sdk := newLayer2Sdk()
	account_operator, _ := newLayer2OperatorAccount(layer2_sdk)
	account_user, _ := newLayer2UserAccount(layer2_sdk)
	//
	balance := getLayer2Balance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
	//
	var gasPrice uint64 = 0
	var gasLimit uint64 = 20000
	var amount uint64 = 100
	{
		txhash, err := layer2Transfer(layer2_sdk, gasPrice, gasLimit, account_user, account_user.Address, account_operator.Address, amount)
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
		} else {
			fmt.Printf("tx hash: %s\n", txhash.ToHexString())
		}
	}
	//
	balance = getLayer2Balance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
}

func TestLayer2TransferPerformance(t *testing.T) {
	// create alliance sdk
	layer2_sdk := newLayer2Sdk()
	account_operator, _ := newLayer2OperatorAccount(layer2_sdk)
	account_user, _ := newLayer2UserAccount(layer2_sdk)
	//
	{
		balance := getLayer2Balance(layer2_sdk, account_user.Address)
		fmt.Printf("amount of address1 %s is: %d\n", account_user.Address.ToBase58(), balance)
		balance1 := getLayer2Balance(layer2_sdk, account_operator.Address)
		fmt.Printf("amount of address2 %s is: %d\n", account_operator.Address.ToBase58(), balance1)
	}
	//
	var gasPrice uint64 = 0
	var gasLimit uint64 = 20000
	var amount uint64 = 100
	var txCounter int = 1000000;

	start := time.Now().Unix()
	for i := 0;i < txCounter;i ++ {
		txhash, err := layer2Transfer(layer2_sdk, gasPrice, gasLimit, account_user, account_user.Address, account_operator.Address, amount)
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
		} else {
			fmt.Printf("tx hash: %s\n", txhash.ToHexString())
		}
	}
	end := time.Now().Unix()
	//
	time.Sleep(time.Second * 5)

	//
	{
		balance := getLayer2Balance(layer2_sdk, account_user.Address)
		fmt.Printf("amount of address1 %s is: %d\n", account_user.Address.ToBase58(), balance)
		balance1 := getLayer2Balance(layer2_sdk, account_operator.Address)
		fmt.Printf("amount of address2 %s is: %d\n", account_operator.Address.ToBase58(), balance1)
	}
	fmt.Printf("tx counter: %d, time: %d\n", txCounter, end - start)
}

func TestLayer2Balance(t *testing.T) {
	// create alliance sdk
	layer2_sdk := newLayer2Sdk()
	account_user, _ := newLayer2UserAccount(layer2_sdk)
	//
	balance := getLayer2Balance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
}

func TestGetCurrentHeight(t *testing.T) {
	layer2_sdk := newLayer2Sdk()
	height, err := layer2_sdk.GetCurrentBlockHeight()
	if err != nil {
		fmt.Printf("get current block height err: %s", err.Error())
	}
	fmt.Printf("current block height: %d\n", height)
}

func TestGetBlock(t *testing.T) {
	layer2_sdk := newLayer2Sdk()
	_, err := layer2_sdk.GetBlockByHeight(20)
	if err != nil {
		fmt.Printf("get current block height err: %s", err.Error())
	}
	fmt.Printf("get current block successful!\n")
}

func TestCommitLayer2State2Ontology(t *testing.T) {
	//
	ontSdk := newOntologySdk()
	contractAddress, _ := ontology_common.AddressFromHexString("4229a92d90d446d1598e12e35698b681ae4d4642")
	depositids := make([]int, 0)
	for i := 0;i < 2;i ++ {
		depositids = append(depositids, 3 + i)
	}
	withdrawAmounts := make([]uint64, 0)
	toAddresses := make([]ontology_common.Address, 0)
	assetAddress := make([][]byte, 0)
	for i := 0;i < 1;i ++ {
		withdrawAmounts = append(withdrawAmounts, 3)
		toAddress, _ := ontology_common.AddressFromBase58("AMUGPqbVJ3TG6pe7xRpxxaeh4ai4fu9ahc")
		toAddresses = append(toAddresses,toAddress)
		tokenAddress, _ := hex.DecodeString("0000000000000000000000000000000000000002")
		assetAddress = append(assetAddress, tokenAddress)
	}
	tx, err := ontSdk.NeoVM.NewNeoVMInvokeTransaction(500, 40000, contractAddress, []interface{}{"updateState", []interface{}{
		"0000000000000000000000000000000000000000000000000000000000000000", 6, "1.0.0",
		depositids, []interface{}{},[]interface{}{},[]interface{}{}}})
	if err != nil {
		fmt.Printf("new transaction failed!")
	}
	account_operator, err := newOntologyOperatorAccount(ontSdk)
	if err != nil {
		fmt.Printf("newOntologyAccount failed!")
	}
	ontSdk.SetPayer(tx, account_operator.Address)
	err = ontSdk.SignToTransaction(tx, account_operator)
	if err != nil {
		fmt.Printf("SignToTransaction failed!")
	}
	txHash, err := ontSdk.SendTransaction(tx)
	if err != nil {
		fmt.Printf("SignToTransaction failed! err: %s", err.Error())
	}
	fmt.Printf("hash: %s", txHash.ToHexString())
}

func TestOntologyDeposit(t *testing.T) {
	ontSdk := newOntologySdk()
	contractAddress, _ := ontology_common.AddressFromHexString("d1b62355d7c88a76fefee8f4ce14efb477992a3c")
	account_user, err := newOntologyUserAccount(ontSdk)
	if err != nil {
		fmt.Printf("ontology account err: %s", err.Error())
		return
	}
	tokenAddress, _ := hex.DecodeString("0000000000000000000000000000000000000002")
	tx, err := ontSdk.NeoVM.NewNeoVMInvokeTransaction(500, 40000, contractAddress, []interface{}{"deposit", []interface{}{
		account_user.Address, 3000000000, tokenAddress}})
	if err != nil {
		fmt.Printf("new transaction failed!")
	}
	ontSdk.SetPayer(tx, account_user.Address)
	err = ontSdk.SignToTransaction(tx, account_user)
	if err != nil {
		fmt.Printf("SignToTransaction failed!")
	}
	txHash, err := ontSdk.SendTransaction(tx)
	if err != nil {
		fmt.Printf("SignToTransaction failed! err: %s", err.Error())
	}
	fmt.Printf("hash: %s", txHash.ToHexString())
}

func TestGetLayer2StateByHeight(t *testing.T) {
	//
	ontSdk := newOntologySdk()
	contractAddress, _ := ontology_common.AddressFromHexString("1c6f220296de4f9c4666a258bd236519d41860d5")
	tx, err := ontSdk.NeoVM.NewNeoVMInvokeTransaction(0, 0, contractAddress, []interface{}{"getStateRootByHeight", []interface{}{1}})
	if err != nil {
		fmt.Printf("new transaction failed!")
	}
	result, err := ontSdk.PreExecTransaction(tx)
	if err != nil {
		fmt.Printf("PreExecTransaction failed! err: %s", err.Error())
		return
	}
	if result == nil {
		fmt.Printf("can not find the result")
		return
	}
	tt, _ := result.Result.ToArray()
	if len(tt) != 3 {
		fmt.Printf("result is not right")
		return
	}
	item0,_ := tt[0].ToString()
	item1,_ := tt[1].ToInteger()
	item2,_ := tt[2].ToInteger()
	fmt.Printf("item0: %s, item1: %d, item2: %d\n", item0, item1, item2)
}

func getOntologyBalance(ontSdk *ontology_go_sdk.OntologySdk, addr ontology_common.Address) uint64 {
	amount, err := ontSdk.Native.Ong.BalanceOf(addr)
	if err != nil {
		fmt.Printf("getOntologyBalance err: %s", err.Error())
	}
	return amount
}

func TestOntologyBalance(t *testing.T) {
	// create alliance sdk
	ontSdk := newOntologySdk()
	account_user, _ := newOntologyUserAccount(ontSdk)
	//
	balance := getOntologyBalance(ontSdk, account_user.Address)
	fmt.Printf("amount of address %s is: %d\n", account_user.Address.ToBase58(), balance)
}


func TestCreateLayer2Account(t *testing.T) {
	createLayer2Account()
}

func TestCreateOntologyAccount(t *testing.T) {
	createOntologyAccount()
}
