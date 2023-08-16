package ethereum

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
	"time"
)

func EthTransfers() {
	client := DefaultEthereumClient()
	ctx := context.Background()

	counter := 10
	// sub account
	priKeys := make([]*ecdsa.PrivateKey, 0, counter)
	for i := 0; i< counter;i ++ {
		prikey := NewPrivateKey1()
		priKeys = append(priKeys, prikey)
	}
	// main account
	mainKey := NewPrivateKey("d2e10ad0c53aec44302b2fd5c1f656fe5ba3f6e7f3ba671d4bfb26ddda93114c")
	signer := types.NewLondonSigner(big.NewInt(3))
	//
	balance_before := client.GetBalance(ctx, crypto.PubkeyToAddress(mainKey.PublicKey).String())
	fmt.Printf("balance: %s\n", balance_before.String())

	// transfer mainkey -> 1000 user keys
	nonce := client.GetNonceAt(ctx, crypto.PubkeyToAddress(mainKey.PublicKey))
	value := big.NewInt(110000000000000000)
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		panic(err)
	}
	for i := 0;i < counter;i ++ {
		toAddr := crypto.PubkeyToAddress(priKeys[i].PublicKey)
		tx := types.NewTransaction(nonce, toAddr, value, uint64(300000), gasPrice, []byte{})
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
	//
	value1 := new(big.Int).Sub(value, new(big.Int).Mul(gasPrice, big.NewInt(300000)))
	for i := 0;i < counter;i ++ {
		fromAddr := crypto.PubkeyToAddress(priKeys[i].PublicKey)
		tx := types.NewTransaction(client.GetNonceAt(ctx, fromAddr), crypto.PubkeyToAddress(mainKey.PublicKey), value1, uint64(300000), gasPrice, []byte{})
		signed_tx, err := types.SignTx(tx, signer, priKeys[i])
		if err != nil {
			panic(err)
		}
		err = client.Client.SendTransaction(context.Background(), signed_tx)
		if err != nil {
			panic(err)
		}
		fmt.Printf("hash: %s\n", signed_tx.Hash().String())
	}

	time.Sleep(time.Second * 10)

	balance_after := client.GetBalance(ctx, crypto.PubkeyToAddress(mainKey.PublicKey).String())
	fmt.Printf("balance: %s\n", balance_after.String())
}

func TestEthTransfers(t *testing.T) {
	EthTransfers()
}
