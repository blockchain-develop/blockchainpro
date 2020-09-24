package ethereum

import (
	"context"
	"fmt"
	"github.com/blockchainpro/usage/ethereum/contractabi/erc20_abi"
	"github.com/blockchainpro/usage/ethereum/contractabi/usdt_abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)

func TestErc20BalanceOf(t *testing.T) {
	client := DefaultEthereumClient()
	erc20Addr_hex := "239100e629a9ca8e0bf45c7892b0fc72d78aa97a"
	erc20Address := ethcommon.HexToAddress(erc20Addr_hex)
	erc20Contract, err := erc20_abi.NewERC20(erc20Address, client.Client)
	if err != nil {
		fmt.Printf("GetSmartContractEventByBlock, error: %s", err.Error())
		return
	}

	userAddr_hex := "74197c375df1330ba9b2319daf98fee26ace7edb"
	userAddress := ethcommon.HexToAddress(userAddr_hex)
	balance, err := erc20Contract.BalanceOf(&bind.CallOpts{}, userAddress)
	if err != nil {
		panic(err)
	}
	fmt.Printf("user: %s balance of %s is : %s\n", userAddr_hex, erc20Addr_hex, balance.String())
}

func TestErc20Info_One(t *testing.T) {
	client := DefaultEthereumClient()
	erc20Addr_hex := "239100e629a9ca8e0bf45c7892b0fc72d78aa97a"
	erc20Address := ethcommon.HexToAddress(erc20Addr_hex)
	erc20Contract, err := usdt_abi.NewTetherToken(erc20Address, client.Client)
	if err != nil {
		fmt.Printf("GetSmartContractEventByBlock, error: %s", err.Error())
		return
	}
	name, err := erc20Contract.Name(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	totolSupply, err := erc20Contract.TotalSupply(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	decimal, err := erc20Contract.Decimals(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	symbol, err := erc20Contract.Symbol(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("erc20: %s, name: %s, totolSupply: %s, decimal: %s, symbol: %s\n",
		erc20Addr_hex, name, totolSupply.String(), decimal.String(), symbol)
}


func TestErc20Info(t *testing.T) {
	client := DefaultEthereumClient()
	tokens := []string{"c8757865920e0467f5d23b59845aa357a24ea38c","cb46c550539ac3db72dc7af7c89b11c306c727c2","dac17f958d2ee523a2206206994597c13d831ec7","2260fac5e5542a773aa44fbcfedf7c193bc2c599","6b175474e89094c44da98b954eedeac495271d0f","a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48","eb4c2781e4eba804ce9a9803c67d0893436bb27d","6a4c89eb9a26a2da34f13f8976daa9fd7526f35c","db0f18081b505a7de20b18ac41856bcb4ba86a1a","57ab1ec28d129707052df4df418d58a2d46d5f51","8e870d67f660d95d5be530380d0ec0bd388289e1","381225768DD2bd60D70482B51109D0DEFeE92503","e179198fd42f5de1a04ffd9a36d6dc428ceb13f7","2205d2f559ef91580090011aa4e0ef68ec33da44","bb44b36e588445d7da61a1e2e426664d03d40888","7757ffe3ac09bc6430f6896f720e77cf80ec1f74","2dd56dc238d1fc2f9aac3793a287f4e0af1b08b4","886f6F287Bb2eA7DE03830a5FD339EDc107c559f","7f0ad0525cb8c17d3f5c06ceb0aea20fa0d2ca0a","7245ded8459f59b0a680640535476c11b3cd0ef6"}
	for _, token := range tokens {
		erc20Addr_hex := token
		erc20Address := ethcommon.HexToAddress(erc20Addr_hex)
		erc20Contract, err := usdt_abi.NewTetherToken(erc20Address, client.Client)
		if err != nil {
			fmt.Printf("GetSmartContractEventByBlock, error: %s", err.Error())
			return
		}
		name, err := erc20Contract.Name(&bind.CallOpts{})
		if err != nil {
			panic(err)
		}
		totolSupply, err := erc20Contract.TotalSupply(&bind.CallOpts{})
		if err != nil {
			panic(err)
		}
		decimal, err := erc20Contract.Decimals(&bind.CallOpts{})
		if err != nil {
			panic(err)
		}
		symbol, err := erc20Contract.Symbol(&bind.CallOpts{})
		if err != nil {
			panic(err)
		}
		fmt.Printf("erc20: %s, name: %s, totolSupply: %s, decimal: %s, symbol: %s\n",
			erc20Addr_hex, name, totolSupply.String(), decimal.String(), symbol)
	}
}


func TestErc20Transfer(t *testing.T) {
	client := DefaultEthereumClient()
	erc20Addr_hex := "239100e629a9ca8e0bf45c7892b0fc72d78aa97a"
	erc20Address := ethcommon.HexToAddress(erc20Addr_hex)
	erc20Contract, err := usdt_abi.NewTetherToken(erc20Address, client.Client)
	if err != nil {
		fmt.Printf("GetSmartContractEventByBlock, error: %s", err.Error())
		return
	}

	privateKey := DefaultPrivateKey()
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	ctx := context.Background()
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(client.GetNonceAt(ctx, address)))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(5000000) // in units
	auth.GasPrice = big.NewInt(30000000000)
	toAddress := ethcommon.HexToAddress("0x5cD3143f91a13Fe971043E1e4605C1c23b46bF44")
	tx, err := erc20Contract.Transfer(auth, toAddress, big.NewInt(0))
	if err != nil {
		fmt.Printf("erc20 transfer - err: %s", err.Error())
	}
	fmt.Printf("erc20 transfer - tx: %s", tx.Hash().String())

	waitTransactionConfirm(client.Client, tx.Hash())
}