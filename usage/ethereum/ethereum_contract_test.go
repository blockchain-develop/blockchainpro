package ethereum

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/blockchainpro/usage/ethereum/contractabi/erc20_abi"
	"github.com/blockchainpro/usage/ethereum/contractabi/lock_proxy_abi"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"strings"
	"testing"
)

func TestDeployContract(t *testing.T) {
	client := DefaultEthereumClient()
	privateKey := DefaultPrivateKey()
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	ctx := context.Background()

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(client.GetNonceAt(ctx, address)))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(5000000) // in units
	auth.GasPrice = big.NewInt(30000000000)
	addr, tx, contract, err := erc20_abi.DeployERC20(auth, client.Client)
	if err != nil {
		fmt.Printf("deploy contract - err: %s", err.Error())
	}
	fmt.Printf("deploy contract - the user address of the new contract is: %s", addr.String())
	fmt.Printf("deploy contract - tx: %s", tx.Hash().String())
	fmt.Printf("deploy contract - contract: %v", contract)
}

func TestInvokeContract(t *testing.T) {
	client := DefaultEthereumClient()
	ctx := context.Background()
	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	if err != nil {
		fmt.Printf("TestInvokeContract - err:" + err.Error())
		return
	}

	toAddr := common.HexToAddress("0x5cD3143f91a13Fe971043E1e4605C1c23b46bF44")
	txData, err := contractabi.Pack("lock", uint64(1), toAddr.Bytes(), big.NewInt(int64(1000000)))
	if err != nil {
		fmt.Printf("TestInvokeContract - err:" + err.Error())
		return
	}

	fmt.Printf("TestInvokeContract - txdata:%s\n", hex.EncodeToString(txData))

	contractAddr := common.HexToAddress("239100e629a9ca8e0bf45c7892b0fc72d78aa97a")
	privateKey := DefaultPrivateKey()
	fromAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
	tx := types.NewTransaction(client.GetNonceAt(ctx, fromAddr), contractAddr, big.NewInt(0), uint64(8000000), big.NewInt(30000000000), txData)
	signed_tx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		fmt.Printf("TestInvokeContract - err:" + err.Error())
		return	
	}

	err = client.Client.SendTransaction(context.Background(), signed_tx)
	if err != nil {
		fmt.Printf("TestInvokeContract - send transaction error:%s\n", err.Error())
		return
	}
	waitTransactionConfirm(client.Client, signed_tx.Hash())
}
