package solana

import (
	"context"
	json2 "encoding/json"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"testing"
)

func AccountInfo(addr solana.PublicKey) {
	client := rpc.New(URL)
	ctx := context.Background()
	result, err := client.GetAccountInfo(ctx, addr)
	if err != nil {
		panic(err)
	}
	json, _ := json2.Marshal(result.Value)
	fmt.Printf("token account: %s\n", json)
}

func TokenAccount(addr solana.PublicKey) {
	client := rpc.New("http://127.0.0.1:8899")
	ctx := context.Background()
	result, err := client.GetAccountInfo(ctx, addr)
	if err != nil {
		panic(err)
	}
	account := decodeAccount(result.Value)
	json, _ := json2.Marshal(account)
	fmt.Printf("token account: %s\n", json)
}


func TestGetAccountInfo(t *testing.T) {
	addr := solana.MustPublicKeyFromBase58("EpkBbKgCTrTrKomgMN9phNeueMyjWqKUE9aMCk9iRNWj")
	AccountInfo(addr)
}

func TestSPLAccountRent(t *testing.T) {
	client := rpc.New(rpc.MainNetBetaSerum_RPC)
	ctx := context.Background()
	slot, err := client.GetSlot(ctx, rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}
	block, err := client.GetBlock(ctx, slot - 10)
	if err != nil {
		panic(err)
	}
	blockJson, _ := json2.MarshalIndent(block, "", "    ")
	fmt.Printf("block: %s\n", string(blockJson))
	lamport, err := client.GetMinimumBalanceForRentExemption(ctx, 165, rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", lamport)
}
//144314326
//144311855
//144312977