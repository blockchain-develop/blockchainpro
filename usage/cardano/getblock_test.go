package cardano

import (
	"context"
	"fmt"
	"github.com/coinbase/rosetta-sdk-go/types"
	"strconv"
	"strings"
	"testing"
)

func TestNetworkList(t *testing.T) {
	client := NewClient()
	ctx := context.Background()
	request := &types.MetadataRequest{}
	networkList, rosettaErr, err := client.NetworkAPI.NetworkList(ctx, request)
	if err != nil {
		panic(err)
	}
	if rosettaErr != nil {
		panic(err)
	}
	fmt.Printf("%v", networkList)
}

func TestNetworkStatus(t *testing.T) {
	client := NewClient()
	ctx := context.Background()
	var primaryNetwork *types.NetworkIdentifier
	{
		request := &types.MetadataRequest{}
		networkList, rosettaErr, err := client.NetworkAPI.NetworkList(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(err)
		}
		primaryNetwork = networkList.NetworkIdentifiers[0]
	}
	request := &types.NetworkRequest{NetworkIdentifier: primaryNetwork}
	networkStatus, rosettaErr, err := client.NetworkAPI.NetworkStatus(ctx, request)
	if err != nil {
		panic(err)
	}
	if rosettaErr != nil {
		panic(err)
	}
	fmt.Printf("current height: %d", networkStatus.CurrentBlockIdentifier.Index)
}

func TestGetBlock(t *testing.T) {
	client := NewClient()
	ctx := context.Background()
	var primaryNetwork *types.NetworkIdentifier
	{
		request := &types.MetadataRequest{}
		networkList, rosettaErr, err := client.NetworkAPI.NetworkList(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(err)
		}
		primaryNetwork = networkList.NetworkIdentifiers[0]
	}
	blockIdentifier := &types.BlockIdentifier{
		Index: 3042923,
		Hash:  "",
	}
	identifier := types.ConstructPartialBlockIdentifier(blockIdentifier)
	request := &types.BlockRequest{NetworkIdentifier: primaryNetwork, BlockIdentifier: identifier}
	block, rosettaErr, err := client.BlockAPI.Block(ctx, request)
	if err != nil {
		panic(err)
	}
	if rosettaErr != nil {
		panic(err)
	}
	fmt.Printf("block height: %d, tarnsaction count: %d", block.Block.BlockIdentifier.Index, len(block.Block.Transactions))
}

func TestGetBlocks(t *testing.T) {
	client := NewClient()
	ctx := context.Background()
	var primaryNetwork *types.NetworkIdentifier
	{
		request := &types.MetadataRequest{}
		networkList, rosettaErr, err := client.NetworkAPI.NetworkList(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(err)
		}
		primaryNetwork = networkList.NetworkIdentifiers[0]
	}
	for i := 0; i < 597700;i ++ {
		blockIdentifier := &types.BlockIdentifier{
			Index: int64(i),
			Hash:  "",
		}
		identifier := types.ConstructPartialBlockIdentifier(blockIdentifier)
		request := &types.BlockRequest{NetworkIdentifier: primaryNetwork, BlockIdentifier: identifier}
		block, rosettaErr, err := client.BlockAPI.Block(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(err)
		}
		if len(block.Block.Transactions) > 0 {
			fmt.Printf("block height: %d, tarnsaction count: %d\n", block.Block.BlockIdentifier.Index, len(block.Block.Transactions))
		}
		fmt.Printf("block height: %d, tarnsaction count: %d\n", block.Block.BlockIdentifier.Index, len(block.Block.Transactions))
	}
}


func TestBalance(t *testing.T) {
	client := NewClient()
	ctx := context.Background()
	var primaryNetwork *types.NetworkIdentifier
	{
		request := &types.MetadataRequest{}
		networkList, rosettaErr, err := client.NetworkAPI.NetworkList(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(err)
		}
		primaryNetwork = networkList.NetworkIdentifiers[0]
	}
	accountIdentifier := &types.AccountIdentifier{
		Address:    "addr_test1qzgd37t7vwufes5xmfltc2ccy8eu5nrrtfg33434f8qp0kdzql220qjjr7klj3z0zxh3dufhg43ne29znkcjxqtrarmqkzazsk",
	}
	currencies := make([]*types.Currency, 0)
	currencies = append(currencies, &types.Currency{
		Symbol:   "ADA",
		//Decimals: 6,
		Metadata: nil,
	})

	request := &types.AccountBalanceRequest{NetworkIdentifier: primaryNetwork, AccountIdentifier: accountIdentifier, Currencies: currencies}
	networkStatus, rosettaErr, err := client.AccountAPI.AccountBalance(ctx, request)
	if err != nil {
		panic(err)
	}
	if rosettaErr != nil {
		panic(err)
	}
	fmt.Printf("balance: %s", networkStatus.Balances[0].Value)
}

func TestTransfer(t *testing.T) {
	client := NewClient()
	ctx := context.Background()
	var primaryNetwork *types.NetworkIdentifier
	{
		request := &types.MetadataRequest{}
		networkList, rosettaErr, err := client.NetworkAPI.NetworkList(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(err)
		}
		primaryNetwork = networkList.NetworkIdentifiers[0]
	}
		blockIdentifier := &types.BlockIdentifier{
			Index: int64(3036948),
			Hash:  "",
		}
		identifier := types.ConstructPartialBlockIdentifier(blockIdentifier)
		request := &types.BlockRequest{NetworkIdentifier: primaryNetwork, BlockIdentifier: identifier}
		block, rosettaErr, err := client.BlockAPI.Block(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(err)
		}
		for _, transaction := range block.Block.Transactions {
			inputs := make([]*types.Operation, 0)
			outputs := make([]*types.Operation, 0)
			for _, operation := range transaction.Operations {
				if operation.Type == "input" {
					inputs = append(inputs, operation)
				} else if operation.Type == "output" {
					outputs = append(outputs, operation)
				} else {
					panic("not support")
				}
			}
			fmt.Printf("transaction hash: %s\n", transaction.TransactionIdentifier.Hash)
			for _, input := range inputs {
				txid, vout := getCoinIdentifier(input.CoinChange.CoinIdentifier)
				fmt.Printf("tx %s, %d, address: %s, coin: %s amount: %s\n", txid, vout,
					input.Account.Address, input.Amount.Currency.Symbol, input.Amount.Value)
			}
			for _, output := range outputs {
				fmt.Printf("address: %s, coin: %s, amount: %s", output.Account.Address, output.Amount.Currency.Symbol, output.Amount.Value)
			}
		}
}

func getCoinIdentifier(coinIdentifier *types.CoinIdentifier) (string, uint32) {
	values := strings.Split(coinIdentifier.Identifier, ":")
	txid := values[0]
	vout, _ := strconv.ParseInt(values[1], 10, 32)
	return txid, uint32(vout)
}
