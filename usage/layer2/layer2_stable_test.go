package layer2

import (
	"encoding/hex"
	"fmt"
	"testing"
	"github.com/ontio/ontology/common"
	"time"
)

func TestStable(t *testing.T) {
	// init
	ont_sdk := newOntologySdk()
	ont_account, _ := newOntologyUserAccount(ont_sdk)
	init_ont_account_balance := getOntologyBalance(ont_sdk, ont_account.Address)
	layer2_sdk := newLayer2Sdk()
	layer2_account, _ := newLayer2UserAccount(layer2_sdk)
	init_layer2_account_balance := getLayer2Balance(layer2_sdk, layer2_account.Address)
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
			txhash, err := layer2WithdrawTransfer(layer2_sdk, layer2_account, layer2_account.Address, uint64(withdraw_amount))
			if err != nil {
				panic(err)
			}
			fmt.Printf("layer2 withdraw tx hash: %s\n", txhash.ToHexString())
		}
		time.Sleep(time.Second * 10)
	}
	time.Sleep(time.Second * 60)
	new_ont_account_balance := getOntologyBalance(ont_sdk, ont_account.Address)
	fmt.Printf("amount of ontology address %s is: %d %d\n", ont_account.Address.ToBase58(), init_ont_account_balance, new_ont_account_balance)
	new_layer2_account_balance := getLayer2Balance(layer2_sdk, layer2_account.Address)
	fmt.Printf("amount of layer2 address %s is: %d %d\n", layer2_account.Address.ToBase58(), init_layer2_account_balance, new_layer2_account_balance)
}

func TestGetProofStable(t *testing.T) {
	sdk := newLayer2Sdk()

	for i := 0;i < 720;i ++ {
		_, err := sdk.GetStoreProof(STORE_CONTRACT, []byte("hello"))
		if err != nil {
			panic(err)
		}
	}
}