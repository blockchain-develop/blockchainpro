package ethereum

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"io/ioutil"
	"math/big"
	"testing"
)

type EthereumClient struct {
	rpcClient *rpc.Client
	Client    *ethclient.Client
}

func DefaultEthereumClient() (client *EthereumClient) {
	//return NewEthereumClient("https://mainnet.infura.io/v3/dc891b662f354817983633c828b46eff")
	return NewEthereumClient("https://ropsten.infura.io/v3/19e799349b424211b5758903de1c47ea")
	//return NewEthereumClient("http://127.0.0.1:8085")
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
	// account 0xd8d50Be55FE241B3c026361a793aA950BceAE845
	return NewPrivateKey("d2e10ad0c53aec44302b2fd5c1f656fe5ba3f6e7f3ba671d4bfb26ddda93114c")
}

func NewPrivateKey(key string) *ecdsa.PrivateKey {
	priKey, err := crypto.HexToECDSA(key)
	if err != nil {
		panic(err)
	}
	return priKey
}

func NewPrivateKey1() *ecdsa.PrivateKey {
	key, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	return key
}

func (ec *EthereumClient) GetBalance(ctx context.Context, addr string) *big.Int {
	address := common.HexToAddress(addr)
	result, err := ec.Client.BalanceAt(ctx, address, nil)
	if err != nil {
		panic(err)
	}
	return result
}

func (ec *EthereumClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return ec.Client.SuggestGasPrice(ctx)
}

func (ec *EthereumClient) GetPendingNonceAt(ctx context.Context, address common.Address) uint64 {
	nonce, err := ec.Client.PendingNonceAt(ctx, address)
	if err != nil {
		panic(err)
	}
	return nonce
}

func (ec *EthereumClient) GetNonceAt(ctx context.Context, address common.Address) uint64 {
	nonce, err := ec.Client.NonceAt(ctx, address, nil)
	if err != nil {
		panic(err)
	}
	return nonce
}

func PrepareAccounts() {
	counter := 500
	// sub account
	for i := 0; i< counter;i ++ {
		prikey := NewPrivateKey1()
		crypto.SaveECDSA("./wallet/" + crypto.PubkeyToAddress(prikey.PublicKey).String(), prikey)
	}
}

func LoadAccounts() []*ecdsa.PrivateKey {
	accounts := make([]*ecdsa.PrivateKey, 0)
	files, _ := ioutil.ReadDir("./wallet/")
	for _, f := range files {
		account, err := crypto.LoadECDSA("./wallet/" + f.Name())
		if err != nil {
			panic(err)
		}
		accounts = append(accounts, account)
	}
	return accounts
}

func EthTransfers1() {
	client := DefaultEthereumClient()
	ctx := context.Background()

	prikeys := LoadAccounts()
	// main account
	mainKey := NewPrivateKey("d2e10ad0c53aec44302b2fd5c1f656fe5ba3f6e7f3ba671d4bfb26ddda93114c")
	signer := types.NewLondonSigner(big.NewInt(3))
	//
	balance_before := client.GetBalance(ctx, crypto.PubkeyToAddress(mainKey.PublicKey).String())
	fmt.Printf("balance: %s\n", balance_before.String())

	// transfer mainkey -> 1000 user keys
	nonce := client.GetNonceAt(ctx, crypto.PubkeyToAddress(mainKey.PublicKey))
	fmt.Printf("nonce: %d\n", nonce)
	value := big.NewInt(100000000000000)
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		panic(err)
	}
	gasPrice = new(big.Int).Add(gasPrice, big.NewInt(1000000))
	for _, prikey := range prikeys {
		toAddr := crypto.PubkeyToAddress(prikey.PublicKey)
		tx := types.NewTransaction(nonce, toAddr, value, uint64(27000), gasPrice, []byte{})
		signed_tx, err := types.SignTx(tx, signer, mainKey)
		if err != nil {
			panic(err)
		}
		err = client.Client.SendTransaction(context.Background(), signed_tx)
		if err != nil {
			panic(err)
		}
		fmt.Printf("hash: %s\n", signed_tx.Hash().String())
		nonce ++
	}
	balance_after := client.GetBalance(ctx, crypto.PubkeyToAddress(mainKey.PublicKey).String())
	fmt.Printf("balance: %s\n", balance_after.String())
}

func EthTransfers2() {
	client := DefaultEthereumClient()
	ctx := context.Background()

	prikeys := LoadAccounts()
	// main account
	mainKey := NewPrivateKey("d2e10ad0c53aec44302b2fd5c1f656fe5ba3f6e7f3ba671d4bfb26ddda93114c")
	signer := types.NewLondonSigner(big.NewInt(3))
	//
	balance_before := client.GetBalance(ctx, crypto.PubkeyToAddress(mainKey.PublicKey).String())
	fmt.Printf("balance: %s\n", balance_before.String())

	value := big.NewInt(100000000000000)
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		panic(err)
	}
	//
	value1 := new(big.Int).Sub(value, new(big.Int).Mul(gasPrice, big.NewInt(27000)))
	//value1 = new(big.Int).Sub(value1, big.NewInt(1000000000))
	for _, prikey := range prikeys {
		fromAddr := crypto.PubkeyToAddress(prikey.PublicKey)
		tx := types.NewTransaction(client.GetNonceAt(ctx, fromAddr), crypto.PubkeyToAddress(mainKey.PublicKey), value1, uint64(27000), gasPrice, []byte{})
		signed_tx, err := types.SignTx(tx, signer, prikey)
		if err != nil {
			panic(err)
		}
		err = client.Client.SendTransaction(context.Background(), signed_tx)
		if err != nil {
			panic(err)
		}
		fmt.Printf("hash: %s\n", signed_tx.Hash().String())
	}
	balance_after := client.GetBalance(ctx, crypto.PubkeyToAddress(mainKey.PublicKey).String())
	fmt.Printf("balance: %s\n", balance_after.String())
}

func EthTransfers3() {
	client := DefaultEthereumClient()
	ctx := context.Background()

	prikeys := LoadAccounts()
	// main account
	mainKey := NewPrivateKey("d2e10ad0c53aec44302b2fd5c1f656fe5ba3f6e7f3ba671d4bfb26ddda93114c")
	signer := types.NewLondonSigner(big.NewInt(3))
	//
	balance_before := client.GetBalance(ctx, crypto.PubkeyToAddress(mainKey.PublicKey).String())
	fmt.Printf("balance: %s\n", balance_before.String())

	// transfer mainkey -> 1000 user keys
	nonce := uint64(10000)
	value := big.NewInt(100000000000000)
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		panic(err)
	}
	for _, prikey := range prikeys {
		toAddr := crypto.PubkeyToAddress(prikey.PublicKey)
		tx := types.NewTransaction(nonce, toAddr, value, uint64(27000), gasPrice, []byte{})
		signed_tx, err := types.SignTx(tx, signer, mainKey)
		if err != nil {
			panic(err)
		}
		err = client.Client.SendTransaction(context.Background(), signed_tx)
		if err != nil {
			panic(err)
		}
		fmt.Printf("hash: %s\n", signed_tx.Hash().String())
		nonce ++
		break
	}
	balance_after := client.GetBalance(ctx, crypto.PubkeyToAddress(mainKey.PublicKey).String())
	fmt.Printf("balance: %s\n", balance_after.String())
}

func TestEthTransfers1(t *testing.T) {
	EthTransfers1()
}

func TestEthTransfers2(t *testing.T) {
	EthTransfers2()
}

func TestEthTransfers3(t *testing.T) {
	EthTransfers3()
}

func TestCreateAccount(t *testing.T)  {
	PrepareAccounts()
	LoadAccounts()
}
