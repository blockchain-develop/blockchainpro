package ethereum

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"math/big"
	"testing"
)

func TestGetCurrentBlockHeight(t *testing.T) {
	ethClient := NewEthereumClient("http://40.115.136.96:8545")
	ctx := context.Background()
	height, err := ethClient.GetCurrentBlockHeight(ctx)
	if err != nil {
		t.Errorf("TestService_GetCurrentBlockHeight %v", err)
	}
	fmt.Println(height)
}

func TestGetHeaderByNumber(t *testing.T) {
	ethClient := NewEthereumClient("http://onto-eth.ont.io:10331")
	ctx := context.Background()
	header, err := ethClient.GetHeaderByNumber(ctx, 1)
	if err != nil {
		t.Errorf("TestService_GetHeaderByNumber %v", err)
	}
	fmt.Println(header.Hash().Hex())
}

func TestGetHeaderByHash(t *testing.T) {
	ethClient := NewEthereumClient("http://onto-eth.ont.io:10331")
	ctx := context.Background()
	hash := common.HexToHash("0x41800b5c3f1717687d85fc9018faac0a6e90b39deaa0b99e7fe4fe796ddeb26a")
	header, err := ethClient.GetHeaderByHash(ctx, hash)
	if err != nil {
		t.Errorf("TestService_GetHeaderByHash %v", err)
	}
	fmt.Println(header.Number)
}

func TestGetBlockByNumber(t *testing.T) {
	ethClient := NewEthereumClient("http://onto-eth.ont.io:10331")
	ctx := context.Background()
	block, err :=ethClient.GetBlockByNumber(ctx, 1)
	if err != nil {
		t.Errorf("TestService_GetHeaderByNumber %v", err)
	}
	fmt.Println(block.Hash().Hex())
}

func TestGetBlockByHash(t *testing.T) {
	ethClient := NewEthereumClient("http://onto-eth.ont.io:10331")
	ctx := context.Background()
	hash := common.HexToHash("0x41800b5c3f1717687d85fc9018faac0a6e90b39deaa0b99e7fe4fe796ddeb26a")
	block, err := ethClient.GetBlockByHash(ctx, hash)
	if err != nil {
		t.Errorf("TestService_GetHeaderByHash %v", err)
	}
	fmt.Println(block.Number())
}

func TestGetTransactionFee(t *testing.T) {
	ethClient := NewEthereumClient("http://onto-eth.ont.io:10331")
	ctx := context.Background()
	hash := common.HexToHash("0x911c32767fabb090813d9661803d508e05a4edef562704679cb351f65b81ada1")

	tx, err := ethClient.GetTransactionByHash(ctx, hash)
	if err != nil {
		panic(err)
	}
	receipt, err := ethClient.GetTransactionReceipt(ctx, hash)
	if err != nil {
		panic(err)
	}
	fee := new(big.Int).Mul(tx.GasPrice(), big.NewInt(int64(receipt.GasUsed)))
	precision := decimal.New(int64(1000000000000000000), 0)
	fee_new := decimal.New(int64(fee.Int64()), 0)
	fmt.Printf("transaction: %s, fee: %s\n", hash.String(), fee_new.Div(precision).String())
}


func TestSuggestGasPrice(t *testing.T) {
	ethClient := NewEthereumClient("https://ropsten.infura.io/v3/1ba5f3635395470e9a3f19ba7a852144")
	ctx := context.Background()
	gasprice, err := ethClient.SuggestGasPrice(ctx)
	if err != nil {
		t.Errorf("TestSuggestGasPrice %v", err)
	}
	fmt.Println(gasprice.String())
}