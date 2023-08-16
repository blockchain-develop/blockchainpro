package solana

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	confirm "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"testing"
	"time"
)

func TestFetchAccount(t *testing.T) {
	client := rpc.New(rpc.MainNetBetaSerum_RPC)
	ctx := context.Background()
	aa := []string{
		"5shiGfzZJgoYPADTD6gmru7v54MgMe2v8woBgufrGsuS",
		"6KCe8ptZFWhesmuoHJRZ91Kwviq8jHQW2vUkTRoGTZNr",
	}
	cc := make([]solana.PublicKey, 0)
	for _, bb := range aa {
		cc = append(cc, solana.MustPublicKeyFromBase58(bb))
	}
	r, err := client.GetMultipleAccounts(ctx, cc...)
	if err != nil {
		panic(err)
	}
	for i, dd := range r.Value {
		fmt.Printf("account: %s, owner: %s\n", cc[i], dd.Owner)
	}
	fmt.Printf("\n")
	Saber     := solana.MustPublicKeyFromBase58("SSwpkEEcbUqx4vtoEByFjSkhKdCT862DNVb52nZg1UZ")
	for i, dd := range r.Value {
		if dd.Owner == Saber {
			fmt.Printf("\"%s\":true,\n", cc[i])
		}
	}
}

func TestGetBlockHash(t *testing.T) {
	client := rpc.New(rpc.MainNetBetaSerum_RPC)
	ctx := context.Background()
	block, err := client.GetBlock(ctx, 117263263)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", block.Blockhash.String())
}

func TestTransfer(t *testing.T) {
	//
	client := rpc.New(rpc.DevNet_RPC)
	ctx := context.Background()
	var err error
	wsclient, err := ws.Connect(context.Background(), rpc.DevNet_WS)
	if err != nil {
		panic(err)
	}
	//
	var userA solana.PrivateKey
	if true {
		userA, err = solana.NewRandomPrivateKey()
		if err != nil {
			panic(err)
		}
		fmt.Printf("user a private key:%s\n", userA)
		fmt.Printf("user a public key:%s\n", userA.PublicKey())
	}
	if false {
		userA = solana.MustPrivateKeyFromBase58("ndh19KyeV2WuYHnJ6ZYCPcTDWBhC3mxoaMAP44L646WBEBn2c8uHxWdPBhfzN4xNRmQX3hLRjLwogP28eN4QCeH")
	}
	//
	amount := uint64(0)
	if true {
		// Airdrop 5 sol to the account so it will have something to transfer:
		out, err := client.RequestAirdrop(
			ctx,
			userA.PublicKey(),
			solana.LAMPORTS_PER_SOL*2,
			rpc.CommitmentFinalized,
		)
		if err != nil {
			panic(err)
		}
		fmt.Println("airdrop transaction signature:", out)
		time.Sleep(time.Second * 5)
	}
	for false {
		balanceResult, err := client.GetBalance(ctx, userA.PublicKey(), rpc.CommitmentFinalized)
		if err != nil {
			panic(err)
		}
		if balanceResult.Value >= 0 {
			amount = balanceResult.Value - 10000000
			break
		}
		time.Sleep(time.Second * 5)
	}
	//
	userB := solana.MustPublicKeyFromBase58("5Exy9MrF8tKPqzq6DsKt4Nkv7rohpRzBEr659euszkFH")
	//
	recent, err := client.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}
	//
	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			system.NewTransferInstruction(
				amount,
				userA.PublicKey(),
				userB,
			).Build(),
		},
		recent.Value.Blockhash,
		solana.TransactionPayer(userA.PublicKey()),
		)
	if err != nil {
		panic(err)
	}
	//
	_, err = tx.Sign(
		func(key solana.PublicKey) *solana.PrivateKey {
			if userA.PublicKey().Equals(key) {
				return &userA
			}
			return nil
		})
	if err != nil {
		panic(err)
	}
	//
	if false {
		sig, err := confirm.SendAndConfirmTransactionWithOpts(
			ctx,
			client,
			wsclient,
			tx,
			true,
			rpc.CommitmentFinalized,
		)
		if err != nil {
			panic(err)
		}
		fmt.Printf("sig: %s\n", sig)
	}
	if true {
		sig, err := client.SendTransactionWithOpts(
			ctx,
			tx,
			true,
			rpc.CommitmentFinalized)
		fmt.Printf("sig: %v\n", sig)
		fmt.Printf("err: %v\n", err)
	}
}



func TestTransfer_Local(t *testing.T) {
	//
	client := rpc.New(URL)
	ctx := context.Background()
	var err error
	/*
	wsclient, err := ws.Connect(context.Background(), WS)
	if err != nil {
		panic(err)
	}
	*/
	//
	var userA solana.PrivateKey
	if true {
		userA, err = solana.NewRandomPrivateKey()
		if err != nil {
			panic(err)
		}
		fmt.Printf("user a private key:%s\n", userA)
		fmt.Printf("user a public key:%s\n", userA.PublicKey())
	}
	if false {
		userA = solana.MustPrivateKeyFromBase58(Key)
	}
	//
	amount := uint64(0)
	if true {
		// Airdrop 5 sol to the account so it will have something to transfer:
		out, err := client.RequestAirdrop(
			ctx,
			userA.PublicKey(),
			solana.LAMPORTS_PER_SOL*2,
			rpc.CommitmentFinalized,
		)
		if err != nil {
			panic(err)
		}
		fmt.Println("airdrop transaction signature:", out)
		time.Sleep(time.Second * 30)
	}
	for true {
		balanceResult, err := client.GetBalance(ctx, userA.PublicKey(), rpc.CommitmentFinalized)
		if err != nil {
			panic(err)
		}
		if balanceResult.Value >= solana.LAMPORTS_PER_SOL*2 {
			amount = solana.LAMPORTS_PER_SOL*2
			break
		}
		time.Sleep(time.Second * 5)
	}
	//
	userB := solana.MustPublicKeyFromBase58("9quHCzxwykSsZf2JUawD6XrjZy7QZ93TvKCp6Zec3gQU")
	//
	recent, err := client.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}
	//
	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			system.NewTransferInstruction(
				amount,
				userA.PublicKey(),
				userB,
			).Build(),
		},
		recent.Value.Blockhash,
		solana.TransactionPayer(userA.PublicKey()),
	)
	if err != nil {
		panic(err)
	}
	//
	_, err = tx.Sign(
		func(key solana.PublicKey) *solana.PrivateKey {
			if userA.PublicKey().Equals(key) {
				return &userA
			}
			return nil
		})
	if err != nil {
		panic(err)
	}
	{
		txJson, _ := json.Marshal(tx)
		fmt.Printf("tx: %s", string(txJson))
	}
	//
	if false {
		sig, err := confirm.SendAndConfirmTransactionWithOpts(
			ctx,
			client,
			//wsclient,
			nil,
			tx,
			true,
			rpc.CommitmentFinalized,
		)
		if err != nil {
			panic(err)
		}
		fmt.Printf("sig: %s\n", sig)
	}
	if true {
		sig, err := client.SendTransactionWithOpts(
			ctx,
			tx,
			true,
			rpc.CommitmentFinalized)
		fmt.Printf("sig: %v\n", sig)
		fmt.Printf("err: %v\n", err)
	}
}
