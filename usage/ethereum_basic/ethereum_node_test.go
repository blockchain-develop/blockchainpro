package ethereum

import (
	"context"
	json2 "encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"math/big"
	"testing"
	"time"
)

func TestGetCurrentBlockHeight(t *testing.T) {
	ethClient := DefaultEthereumClient()
	ctx := context.Background()
	height, err := ethClient.GetCurrentBlockHeight(ctx)
	if err != nil {
		t.Errorf("TestService_GetCurrentBlockHeight %v", err)
	}
	fmt.Println(height)
}

func TestGetHeaderByNumber(t *testing.T) {
	ethClient := DefaultEthereumClient()
	ctx := context.Background()
	header, err := ethClient.GetHeaderByNumber(ctx, 1)
	if err != nil {
		t.Errorf("TestService_GetHeaderByNumber %v", err)
	}
	fmt.Println(header.Hash().Hex())
}

func TestGetBalance(t *testing.T) {
	ethClient := DefaultEthereumClient()
	ctx := context.Background()
	addr := common.HexToAddress("0xd8d50Be55FE241B3c026361a793aA950BceAE845")
	result, err := ethClient.Client.BalanceAt(ctx, addr, nil)
	if err != nil {
		t.Errorf("TestGetBalance %v", err)
	}
	fmt.Println(result.String())
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
	block, err := ethClient.GetBlockByNumber(ctx, 1)
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

func TestGetTransaction(t *testing.T) {
	ethClient := NewEthereumClient("https://mainnet.infura.io/v3/dc891b662f354817983633c828b46eff")
	ctx := context.Background()
	hash := common.HexToHash("0xa4ad61391bcbff0c9a2465879f5b2cda0a0ccf324ad75c41a338a10bfe7fc071")
	transaction, err := ethClient.GetTransactionByHash(ctx, hash)
	if err != nil {
		t.Errorf("TestSuggestGasPrice %v", err)
	}
	json, _ := json2.Marshal(transaction)
	fmt.Printf("%s", json)
}

func TestGetTransactionReceipt(t *testing.T) {
	ethClient := NewEthereumClient("https://data-seed-prebsc-1-s1.binance.org:8545")
	ctx := context.Background()
	hash := common.HexToHash("0x4b9e1ed2a7d9fbddd9394429cf3a6432438a1e5a8b161b9206cbfddecbe109a5")
	transaction, err := ethClient.GetTransactionReceipt(ctx, hash)
	if err != nil {
		t.Errorf("TestSuggestGasPrice %v", err)
	}
	json, _ := json2.Marshal(transaction)
	fmt.Printf("%s", json)
}

func TestGetTransactionReceipt1(t *testing.T) {
	ethClient := NewEthereumClient("https://palpable-newest-county.bsc.quiknode.pro/568b217050b55a0d4e7cc2f8e048e3de32a3fbaf")
	ctx := context.Background()
	var height int64
	var err error
	var block *types.Block
	var receipt *types.Receipt
	{ // get latest block height
		height, err = ethClient.GetCurrentBlockHeight(ctx)
		if err != nil {
			panic(err)
		}
	}
	for true {
		t0 := time.Now().UnixNano()
		fmt.Printf("try block: %d, time: %d\n", height, t0)
		for true {
			block, err = ethClient.GetBlockByNumber(ctx, height)
			if err != nil {
				fmt.Printf("GetBlockByNumber error: %s\n", err.Error())
				continue
			}
			break
		}
		t1 := time.Now().UnixNano()
		fmt.Printf("get block: %s, time: %d\n", block.Hash().String(), t1)
		transactions := block.Transactions()
		if len(transactions) == 0 {
			continue
		}
		// only for test
		transaction := transactions[0]
		for true {
			receipt, err = ethClient.GetTransactionReceipt(ctx, transaction.Hash())
			if err != nil {
				fmt.Printf("GetTransactionReceipt error: %s\n", err.Error())
				continue
			}
			break
		}
		t2 := time.Now().UnixNano()
		fmt.Printf("get receipts: %d, block hash: %s, time: %d, diff: %d\n", len(receipt.Logs), receipt.BlockHash.String(), t2, t2-t1)
		height += 1
	}
}

func TestGetTransactionReceipt2(t *testing.T) {
	ethClient := NewEthereumClient("https://palpable-newest-county.bsc.quiknode.pro/568b217050b55a0d4e7cc2f8e048e3de32a3fbaf")
	ctx := context.Background()
	var height int64
	var err error
	var block *types.Block
	var logs []types.Log
	{ // get latest block height
		height, err = ethClient.GetCurrentBlockHeight(ctx)
		if err != nil {
			panic(err)
		}
	}
	for true {
		t0 := time.Now().UnixNano()
		fmt.Printf("try block: %d, time: %d\n", height, t0)
		for true {
			block, err = ethClient.GetBlockByNumber(ctx, height)
			if err != nil {
				fmt.Printf("GetBlockByNumber error: %s\n", err.Error())
				continue
			}
			break
		}
		t1 := time.Now().UnixNano()
		fmt.Printf("get block: %s, time: %d\n", block.Hash().String(), t1)
		// only for test
		for true {
			logs, err = ethClient.GetLogs(ctx, height, height)
			if err != nil {
				fmt.Printf("GetTransactionReceipt error: %s\n", err.Error())
				continue
			}
			break
		}
		t2 := time.Now().UnixNano()
		fmt.Printf("get logs: %d, time: %d, diff: %d\n", len(logs), t2, t2-t1)
		height += 1
	}
}
