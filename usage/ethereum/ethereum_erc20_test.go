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

func TestErc20Info(t *testing.T) {
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