package ethereum_london_fork

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"golang.org/x/crypto/sha3"
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


func TestEIP1559Transaction4(t *testing.T) {
	ethClient, _ := ethclient.Dial("https://ropsten.infura.io/v3/19e799349b424211b5758903de1c47ea")
	ctx := context.Background()
	privateKey,_ := crypto.HexToECDSA("994D7BC4C1DE95D4C3069F3F64A032BC55482970F40141D074141D099CC88569")
	fromAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce,_ := ethClient.PendingNonceAt(ctx, fromAddr)
	to := common.HexToAddress("0xd8d50Be55FE241B3c026361a793aA950BceAE845")
	amount := big.NewInt(1000000000)
	gasPrice, _ := ethClient.SuggestGasPrice(ctx)
	gasLimit := uint64(40000)
	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, nil)
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
	signer := types.MakeSigner(params.RopstenChainConfig, new(big.Int).SetInt64(10660000))
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

func TestEIP1559TransferErc20(t *testing.T) {
	ethClient, _ := ethclient.Dial("https://ropsten.infura.io/v3/19e799349b424211b5758903de1c47ea")
	ctx := context.Background()
	// user address
	privateKey,_ := crypto.HexToECDSA("d2e10ad0c53aec44302b2fd5c1f656fe5ba3f6e7f3ba671d4bfb26ddda93114c")
	fromAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
	// nonce
	nonce,_ := ethClient.PendingNonceAt(ctx, fromAddr)
	// eth amount
	amount := big.NewInt(0)
	// erc20 contract
	contract := common.HexToAddress("0x8682c69b0e6beae23821068edbe89caea498b3b4")
	// contract data
	funcSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(funcSignature)
	methodId := hash.Sum(nil)[:4]
	to := common.HexToAddress("0x5cD3143f91a13Fe971043E1e4605C1c23b46bF44")
	paddedTo := common.LeftPadBytes(to.Bytes(), 32)
	value := big.NewInt(10000000000)
	paddedValue := common.LeftPadBytes(value.Bytes(), 32)

	var data []byte
	data = append(data, methodId...)
	data = append(data, paddedTo...)
	data = append(data, paddedValue...)

	// tip price
	gasTipPrice, _ := ethClient.SuggestGasTipCap(ctx)
	//gasFeePrice := abi.MaxUint256
	gasFeePrice, _ := ethClient.SuggestGasPrice(ctx)
	// gas limit
	callMsg := ethereum.CallMsg{
		From: fromAddr, To: &contract, Gas: 0, GasFeeCap: abi.MaxUint256, GasTipCap: gasTipPrice,
		Value: amount, Data: data,
	}
	gasLimit, err := ethClient.EstimateGas(context.Background(), callMsg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("erc20 transfer: \n")
	fmt.Printf("    user: %s\n", fromAddr.String())
	fmt.Printf("    nonce: %d\n", nonce)
	fmt.Printf("    eth amount: %s\n", amount.String())
	fmt.Printf("    gas tip price: %s\n", gasTipPrice.String())
	fmt.Printf("    gas fee price: %s\n", gasFeePrice.String())
	fmt.Printf("    gas limit: %d\n", gasLimit)
	fmt.Printf("    erc20 contect: %s\n", contract.String())
	fmt.Printf("    transfer parameter:\n")
	fmt.Printf("        to: %s\n", to.String())
	fmt.Printf("        erc20 amount: %s\n", value.String())
	tx := types.NewTx(&types.DynamicFeeTx{
		Nonce:    nonce,
		To:       &contract,
		Value:    amount,
		Gas:      gasLimit,
		GasFeeCap: gasFeePrice,
		GasTipCap: gasTipPrice,
		Data:     data,
	})
	signer := types.MakeSigner(params.RopstenChainConfig, new(big.Int).SetInt64(10660000))
	signed_tx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		fmt.Printf("TestErc20Contract - err:" + err.Error())
		return
	}
	err = ethClient.SendTransaction(context.Background(), signed_tx)
	if err != nil {
		fmt.Printf("TestErc20Contract - send transaction error:%s\n", err.Error())
		return
	} else {
		fmt.Printf("tx hash : %s", signed_tx.Hash())
	}
}
