package ethereum_london_fork

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"testing"
)

func TestEIP1559Transaction1(t *testing.T) {
	ethClient, _ := ethclient.Dial("https://ropsten.infura.io/v3/19e799349b424211b5758903de1c47ea")
	ctx := context.Background()
	privateKey,_ := crypto.HexToECDSA("994D7BC4C1DE95D4C3069F3F64A032BC55482970F40141D074141D099CC88569")
	fromAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce,_ := ethClient.PendingNonceAt(ctx, fromAddr)
	contractAddr := common.HexToAddress("239100e629a9ca8e0bf45c7892b0fc72d78aa97a")
	amount := big.NewInt(0)
	gasPrice, _ := ethClient.SuggestGasPrice(ctx)
	gasLimit := uint64(8000000)
	tx := types.NewTransaction(nonce, contractAddr, amount, gasLimit, gasPrice, nil)
	/*
	tx = types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &contractAddr,
		Value:    amount,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     nil,
	})
	 */
	signer := types.MakeSigner(params.RopstenChainConfig, new(big.Int).SetInt64(12900000))
	signed_tx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		fmt.Printf("TestInvokeContract - err:" + err.Error())
		return
	}
	err = ethClient.SendTransaction(context.Background(), signed_tx)
	if err != nil {
		fmt.Printf("TestInvokeContract - send transaction error:%s\n", err.Error())
		return
	}
}

func TestEIP1559Transaction2(t *testing.T) {
	ethClient, _ := ethclient.Dial("https://ropsten.infura.io/v3/19e799349b424211b5758903de1c47ea")
	ctx := context.Background()
	privateKey,_ := crypto.HexToECDSA("994D7BC4C1DE95D4C3069F3F64A032BC55482970F40141D074141D099CC88569")
	fromAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce,_ := ethClient.PendingNonceAt(ctx, fromAddr)
	contractAddr := common.HexToAddress("239100e629a9ca8e0bf45c7892b0fc72d78aa97a")
	amount := big.NewInt(0)
	gasPrice, _ := ethClient.SuggestGasPrice(ctx)
	gasTip, _ := ethClient.SuggestGasTipCap(ctx)
	gasLimit := uint64(8000000)
	tx := types.NewTransaction(nonce, contractAddr, amount, gasLimit, gasPrice.Sub(gasPrice, gasTip), nil)
	/*
	tx = types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &contractAddr,
		Value:    amount,
		Gas:      gasLimit,
		GasPrice: gasPrice.Sub(gasPrice, gasTip),
		Data:     nil,
	})
	*/
	signer := types.MakeSigner(params.RopstenChainConfig, new(big.Int).SetInt64(12900000))
	signed_tx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		fmt.Printf("TestInvokeContract - err:" + err.Error())
		return
	}
	err = ethClient.SendTransaction(context.Background(), signed_tx)
	if err != nil {
		fmt.Printf("TestInvokeContract - send transaction error:%s\n", err.Error())
		return
	}
}

func TestEIP1559Transaction3(t *testing.T) {
	ethClient, _ := ethclient.Dial("https://ropsten.infura.io/v3/19e799349b424211b5758903de1c47ea")
	ctx := context.Background()
	privateKey,_ := crypto.HexToECDSA("994D7BC4C1DE95D4C3069F3F64A032BC55482970F40141D074141D099CC88569")
	fromAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce,_ := ethClient.PendingNonceAt(ctx, fromAddr)
	contractAddr := common.HexToAddress("239100e629a9ca8e0bf45c7892b0fc72d78aa97a")
	amount := big.NewInt(0)
	//gasPrice, _ := ethClient.SuggestGasPrice(ctx)
	gasTip, _ := ethClient.SuggestGasTipCap(ctx)
	gasLimit := uint64(8000000)
	tx := types.NewTx(&types.DynamicFeeTx{
		Nonce:    nonce,
		To:       &contractAddr,
		Value:    amount,
		Gas:      gasLimit,
		GasFeeCap: abi.MaxUint256,
		GasTipCap: gasTip,
		Data:     nil,
	})
	signer := types.MakeSigner(params.RopstenChainConfig, new(big.Int).SetInt64(12900000))
	signed_tx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		fmt.Printf("TestInvokeContract - err:" + err.Error())
		return
	}
	err = ethClient.SendTransaction(context.Background(), signed_tx)
	if err != nil {
		fmt.Printf("TestInvokeContract - send transaction error:%s\n", err.Error())
		return
	}
}
