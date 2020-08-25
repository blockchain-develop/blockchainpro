package ethereum

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestGetCurrentBlockHeight(t *testing.T) {
	ethClient := NewEthereumClient("http://onto-eth.ont.io:10331")
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
