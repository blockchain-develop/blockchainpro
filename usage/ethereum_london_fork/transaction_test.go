package ethereum_london_fork

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"testing"
)

func TestEthTransfer(t *testing.T) {
	client := DefaultEthereumClient()
	ctx := context.Background()
	contractAddr := common.HexToAddress("239100e629a9ca8e0bf45c7892b0fc72d78aa97a")
	privateKey := DefaultPrivateKey()
	fromAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
	client.Client.SuggestGasPrice(ctx)
	tx := types.NewTransaction(client.GetNonceAt(ctx, fromAddr), contractAddr, big.NewInt(0), uint64(8000000), big.NewInt(30000000000), nil)
	signer := types.MakeSigner(params.RopstenChainConfig, new(big.Int).SetInt64(1))
	signed_tx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		fmt.Printf("TestInvokeContract - err:" + err.Error())
		return
	}

	err = client.Client.SendTransaction(context.Background(), signed_tx)
	if err != nil {
		fmt.Printf("TestInvokeContract - send transaction error:%s\n", err.Error())
		return
	}
}
