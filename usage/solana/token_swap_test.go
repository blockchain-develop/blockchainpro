package solana

import (
	"bytes"
	"context"
	"encoding/binary"
	json2 "encoding/json"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"testing"
)

func TestTokenSwap_Market(t *testing.T) {
	client := rpc.New("http://127.0.0.1:8899")
	ctx := context.Background()
	lending := solana.MustPublicKeyFromBase58("9bHt39hwT8yXECSknLXbkGAAEDDdmLQodGtXd8NHD52u")
	//
	getProgramAccountsResult, err := client.GetProgramAccounts(ctx, lending)
	if err != nil {
		panic(err)
	}
	for _, account := range getProgramAccountsResult {
		data := account.Account.Data.GetBinary()
		if len(data) == int(SwapLayoutSize) {
			market := SwapLayout{}
			buf := bytes.NewReader(data)
			err = binary.Read(buf, binary.LittleEndian, &market)
			json, _ := json2.Marshal(market)
			fmt.Printf("market: %s\n", json)
			TokenAccount(market.TokenA)
			TokenAccount(market.TokenB)
		} else {
			fmt.Printf("account: %s, data size: %d\n", account.Pubkey.String(), len(account.Account.Data.GetBinary()))
		}
	}
}

func TestTokenSwap(t *testing.T) {
	client := rpc.New("http://127.0.0.1:8899")
	ctx := context.Background()
	//ALend7Ketfx5bxh6ghsCDXAoDrhvEmsXT3cynB6aPLgx
	lending := solana.MustPublicKeyFromBase58("9bHt39hwT8yXECSknLXbkGAAEDDdmLQodGtXd8NHD52u")
	//
	getProgramAccountsResult, err := client.GetProgramAccounts(ctx, lending)
	if err != nil {
		panic(err)
	}
	for _, account := range getProgramAccountsResult {
		fmt.Printf("account: %s, data len: %d\n", account.Pubkey.String(), len(account.Account.Data.GetBinary()))
	}
}
