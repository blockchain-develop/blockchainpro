package ethereum

import (
	"context"
	"fmt"
	"github.com/blockchainpro/usage/ethereum/contractabi/usdt_abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

func Erc20Transfer() {
	client := DefaultEthereumClient()
	erc20AddressHex := "0x8682c69b0e6beae23821068edbe89caea498b3b4"
	erc20Address := ethcommon.HexToAddress(erc20AddressHex)
	erc20Contract, err := usdt_abi.NewTetherToken(erc20Address, client.Client)
	if err != nil {
		fmt.Printf("new token, error: %s", err.Error())
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


func EthTransfer() {
	client := DefaultEthereumClient()
	erc20AddressHex := "0x8682c69b0e6beae23821068edbe89caea498b3b4"
	erc20Address := ethcommon.HexToAddress(erc20AddressHex)
	erc20Contract, err := usdt_abi.NewTetherToken(erc20Address, client.Client)
	if err != nil {
		fmt.Printf("new token, error: %s", err.Error())
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
