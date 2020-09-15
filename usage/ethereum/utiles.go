package ethereum

import (
	"crypto/ecdsa"
	"fmt"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

type EthereumClient struct {
	rpcClient *rpc.Client
	Client    *ethclient.Client
}

func DefaultEthereumClient() (client *EthereumClient) {
	return NewEthereumClient("https://ropsten.infura.io/v3/dc891b662f354817983633c828b46eff")
}

func NewEthereumClient(url string) (client *EthereumClient) {
	rpcClient, err := rpc.Dial(url)
	if err != nil {
		fmt.Printf("can't connect to ethereum, %s", err)
		panic(err)
	}
	rawClient, err := ethclient.Dial(url)
	if err != nil {
		fmt.Printf("can't connect to ethereum, %s", err)
		panic(err)
	}
	return &EthereumClient{
		rpcClient: rpcClient,
		Client: rawClient,
	}
}


func DefaultPrivateKey() *ecdsa.PrivateKey {
	return NewPrivateKey("994D7BC4C1DE95D4C3069F3F64A032BC55482970F40141D074141D099CC88569")
}

func NewPrivateKey(key string) *ecdsa.PrivateKey {
	priKey, err := crypto.HexToECDSA(key)
	if err != nil {
		panic(err)
	}
	return priKey
}
// GetBlockByHash returns the given full block.
func (ec *EthereumClient) GetBlockByHash(ctx context.Context, hash common.Hash) (block *types.Block, _ error) {
	block, err := ec.Client.BlockByHash(ctx, hash)
	return block, err
}

// GetBlockByNumber returns the current canonical chain.
func (ec *EthereumClient) GetBlockByNumber(ctx context.Context, number int64) (block *types.Block, _ error) {
	if number < 0 {
		Block, err := ec.Client.BlockByNumber(ctx, nil)
		return Block, err
	} else {
		block, err := ec.Client.BlockByNumber(ctx, big.NewInt(number))
		return block, err
	}
}

// GetHeaderByNumber returns the given header
func (ec *EthereumClient) GetHeaderByNumber(ctx context.Context, number int64) (header *types.Header, err error) {
	if number < 0 {
		header, err = ec.Client.HeaderByNumber(ctx, nil)
	} else {
		header, err = ec.Client.HeaderByNumber(ctx, big.NewInt(number))
	}
	return header, err
}

// GetHeaderByHash returns the block header with the given hash.
func (ec *EthereumClient) GetHeaderByHash(ctx context.Context, hash common.Hash) (header *types.Header, _ error) {
	header, err := ec.Client.HeaderByHash(ctx, hash)
	return header, err
}

// GetCurrentBlockHeight returns current block height
func (ec *EthereumClient) GetCurrentBlockHeight(ctx context.Context) (height int64, _ error) {
	var result hexutil.Big
	err := ec.rpcClient.CallContext(ctx, &result, "eth_blockNumber")
	return (*big.Int)(&result).Int64(), err
}

func (ec *EthereumClient) GetTransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, error) {
	tx, _, err := ec.Client.TransactionByHash(ctx, hash)
	if err != nil {
		return nil, err
	} else {
		return tx, nil
	}
}

func (ec *EthereumClient) GetTransactionReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	receipt, err := ec.Client.TransactionReceipt(ctx, hash)
	if err != nil {
		return nil, err
	} else {
		return receipt, nil
	}
}

func (ec *EthereumClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return ec.Client.SuggestGasPrice(ctx)
}

func (ec *EthereumClient) GetNonceAt(ctx context.Context, address common.Address) uint64 {
	nonce, err := ec.Client.PendingNonceAt(ctx, address)
	if err != nil {
		panic(err)
	}
	return nonce
}

// Close client
func (ec *EthereumClient) Close() {
	ec.rpcClient.Close()
	ec.Client.Close()
}
