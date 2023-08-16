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

func TestLending_Market(t *testing.T) {
	client := rpc.New(rpc.DevNet_RPC)
	ctx := context.Background()
	lending := solana.MustPublicKeyFromBase58("6TvznH3B2e3p2mbhufNBpgSrLx6UkgvxtVQvopEZ2kuH")
	//
	filters := make([]rpc.RPCFilter, 0)
	/*
	filters = append(filters, rpc.RPCFilter{
		DataSize: LendingMarketLen,
	})
	filters = append(filters, rpc.RPCFilter{
		DataSize: ReserveLen,
	})
	*/

	getProgramAccountsResult, err := client.GetProgramAccountsWithOpts(ctx, lending,
		&rpc.GetProgramAccountsOpts{
			Encoding: solana.EncodingBase64,
			Filters:  filters,
		})
	if err != nil {
		panic(err)
	}
	for _, account := range getProgramAccountsResult {
		data := account.Account.Data.GetBinary()
		if len(data) == int(LendingMarketLen) {
			market := LendingMarketLayout{}
			buf := bytes.NewReader(data)
			err = binary.Read(buf, binary.LittleEndian, &market)
			json, _ := json2.Marshal(market)
			fmt.Printf("market: %s\n", json)
		} else if len(data) == int(ReserveLen) {
			reserve := ReserveLayout{}
			buf := bytes.NewReader(data)
			err = binary.Read(buf, binary.LittleEndian, &reserve)
			json, _ := json2.Marshal(reserve)
			fmt.Printf("market: %s\n", json)
		} else {
			fmt.Printf("account: %s, data size: %d\n", account.Pubkey.String(), len(account.Account.Data.GetBinary()))
		}
	}
}

func TestLending_Market1(t *testing.T) {
	client := rpc.New(rpc.DevNet_RPC)
	ctx := context.Background()
	//ALend7Ketfx5bxh6ghsCDXAoDrhvEmsXT3cynB6aPLgx
	lending := solana.MustPublicKeyFromBase58("6TvznH3B2e3p2mbhufNBpgSrLx6UkgvxtVQvopEZ2kuH")
	//
	getProgramAccountsResult, err := client.GetProgramAccounts(ctx, lending)
	if err != nil {
		panic(err)
	}
	for _, account := range getProgramAccountsResult {
		fmt.Printf("account: %s, data len: %d\n", account.Pubkey.String(), len(account.Account.Data.GetBinary()))
	}
}
