package layer2

import (
	"encoding/hex"
	"fmt"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/tendermint/iavl"
	"testing"
	"time"
)

func TestLayer2TransferStable(t *testing.T) {
	// init
	ont_sdk := newOntologySdk()
	ont_account, _ := newOntologyUserAccount(ont_sdk)
	init_ont_account_balance := getOntologyOngBalance(ont_sdk, ont_account.Address)
	layer2_sdk := newLayer2Sdk()
	layer2_account, _ := newLayer2UserAccount(layer2_sdk)
	init_layer2_account_balance := getLayer2OngBalance(layer2_sdk, layer2_account.Address)
	layer2ContractAddress, _ := common.AddressFromHexString(LAYER2_CONTRACT)
	tokenAddress, _ := hex.DecodeString("0000000000000000000000000000000000000002")
	deposit_amount := 300000000
	withdraw_amount := 200000000
	for i := 0;i < 720;i ++ {
		{
			tx, err := ontologyDeposit(ont_sdk, ont_account, layer2ContractAddress, tokenAddress, uint64(deposit_amount))
			if err != nil {
				panic(err)
			}
			fmt.Printf("ontology deposit tx hash: %s\n", tx.ToHexString())
		}
		{
			txhash, err := layer2WithdrawTransferOng(layer2_sdk, layer2_account, layer2_account.Address, uint64(withdraw_amount))
			if err != nil {
				panic(err)
			}
			fmt.Printf("layer2 withdraw tx hash: %s\n", txhash.ToHexString())
		}
		time.Sleep(time.Second * 10)
	}
	time.Sleep(time.Second * 60)
	new_ont_account_balance := getOntologyOngBalance(ont_sdk, ont_account.Address)
	fmt.Printf("amount of ontology address %s is: %d %d\n", ont_account.Address.ToBase58(), init_ont_account_balance, new_ont_account_balance)
	new_layer2_account_balance := getLayer2OngBalance(layer2_sdk, layer2_account.Address)
	fmt.Printf("amount of layer2 address %s is: %d %d\n", layer2_account.Address.ToBase58(), init_layer2_account_balance, new_layer2_account_balance)
}

func TestGetProofStable(t *testing.T) {
	sdk := newLayer2Sdk()
	key, _ := sdk.GetLayer2StoreKey(STORE_CONTRACT, []byte("hello"))
	for i := 0;i < 720;i ++ {
		store, err := sdk.GetLayer2StoreProof(key)
		if err != nil {
			panic(err)
		}

		fmt.Printf("value: %s, proof: %s, height: %d\n", store.Value, store.Proof, store.Height)

		proof_byte, _ := hex.DecodeString(store.Proof)
		source := common.NewZeroCopySource(proof_byte)
		proof := new(utils.Layer2StoreProof)
		err = proof.Deserialization(source)
		if err != nil {
			panic(err)
		}
		proof_iavl := iavl.RangeProof(*proof)

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

		err = proof_iavl.Verify(block.Header.StateRoot.ToArray())
		if err != nil {
			panic(err)
		}
		fmt.Printf("verify successful!\n")
		time.Sleep(time.Second * 1)
	}
}