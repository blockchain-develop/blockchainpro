package ethereum

import (
	"context"
	"fmt"
	"github.com/blockchainpro/usage/ethereum/contractabi/erc20_abi"
	"github.com/blockchainpro/usage/ethereum/contractabi/usdt_abi"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"strings"
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
	erc20Addr_hex := "Bb2b8038a1640196FbE3e38816F3e67Cba72D940"
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
	user1 := ethcommon.HexToAddress("c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	balance1, err := erc20Contract.BalanceOf(&bind.CallOpts{}, user1)
	if err != nil {
		panic(err)
	}

	user2 := ethcommon.HexToAddress("2260fac5e5542a773aa44fbcfedf7c193bc2c599")
	balance2, err := erc20Contract.BalanceOf(&bind.CallOpts{}, user2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("erc20: %s, name: %s, totolSupply: %s, decimal: %s, symbol: %s\n",
		erc20Addr_hex, name, totolSupply.String(), decimal.String(), symbol)
	fmt.Printf("user1 balance: %s, user2 balance: %s\n", balance1.String(), balance2.String())
}

func TestErc20Info_Two(t *testing.T) {
	client := DefaultEthereumClient()
	erc20Addr_hex := "ee9801669c6138e84bd50deb500827b776777d28"
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
	user1 := ethcommon.HexToAddress("c23E44fF96EE6070f446826FB4EEcFE917F3C6fE")
	balance1, err := erc20Contract.BalanceOf(&bind.CallOpts{}, user1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("erc20: %s, name: %s, totolSupply: %s, decimal: %s, symbol: %s， balance:%s\n",
		erc20Addr_hex, name, totolSupply.String(), decimal.String(), symbol, balance1.String())
}

func TestErc20Info_Three(t *testing.T) {
	client := DefaultEthereumClient()
	erc20Addr_hex := "c7283b66Eb1EB5FB86327f08e1B5816b0720212B"
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
	user1 := ethcommon.HexToAddress("Bb2b8038a1640196FbE3e38816F3e67Cba72D940")
	balance1, err := erc20Contract.BalanceOf(&bind.CallOpts{}, user1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("erc20: %s, name: %s, totolSupply: %s, decimal: %s, symbol: %s, balance: %s\n",
		erc20Addr_hex, name, totolSupply.String(), decimal.String(), symbol, balance1.String())
}

func TestErc20Info(t *testing.T) {
	client := DefaultEthereumClient()
	tokens := []string{"aa00FBcFAce6FF4913CC8322B12F63e2AD45D448"}
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
		owner := ethcommon.HexToAddress("D8aE73e06552E270340b63A8bcAbf9277a1aac99")
		balance, err := erc20Contract.BalanceOf(&bind.CallOpts{}, owner)
		if err != nil {
			panic(err)
		}
		fmt.Printf("blance of address: %s is %d\n", owner.String(), balance.Uint64())
		fmt.Printf("erc20: %s, name: %s, totolSupply: %s, decimal: %s, symbol: %s\n",
			erc20Addr_hex, name, totolSupply.String(), decimal.String(), symbol)
	}
}

func TestErc20Transfer(t *testing.T) {
	client := DefaultEthereumClient()
	erc20Addr_hex := "0x92d6c1e31e14520e676a687f0a93788b716beff5"
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

func TestErc20Transfer2(t *testing.T) {
	client := DefaultEthereumClient()
	//ctx := context.Background()
	contractabi, err := abi.JSON(strings.NewReader(usdt_abi.ERC20ABI))
	if err != nil {
		fmt.Printf("TestInvokeContract - err:" + err.Error())
		return
	}

	toAddr := ethcommon.HexToAddress("0x92d6c1e31e14520e676a687f0a93788b716beff5")
	txData, err := contractabi.Pack("transfer", toAddr, new(big.Int).SetInt64(1000000000000000000000))
	if err != nil {
		fmt.Printf("TestInvokeContract - err:" + err.Error())
		return
	}

	fromAddr := ethcommon.HexToAddress("")

	callMsg := ethereum.CallMsg{
		From: fromAddr, To: &toAddr, Gas: 0, GasPrice: big.NewInt(30000000000),
		Value: big.NewInt(0), Data: txData,
	}
	gasLimit, err := client.Client.EstimateGas(context.Background(), callMsg)
	if err != nil {
		fmt.Printf("commitDepositEventsWithHeader - estimate gas limit error: %s", err.Error())
		panic(err)
	}
	fmt.Printf("gas limit: %d", gasLimit)
}
