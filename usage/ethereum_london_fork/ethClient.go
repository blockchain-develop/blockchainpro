package ethereum_london_fork

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type EthereumClient struct {
	rpcClient *rpc.Client
	Client    *ethclient.Client
}

func DefaultEthereumClient() (client *EthereumClient) {
	return NewEthereumClient("https://mainnet.infura.io/v3/dc891b662f354817983633c828b46eff")
	//return NewEthereumClient("https://ropsten.infura.io/v3/19e799349b424211b5758903de1c47ea")
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

func (ec *EthereumClient) GetNonceAt(ctx context.Context, address common.Address) uint64 {
	nonce, err := ec.Client.PendingNonceAt(ctx, address)
	if err != nil {
		panic(err)
	}
	return nonce
}