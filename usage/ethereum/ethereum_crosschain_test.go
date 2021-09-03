package ethereum

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/polynetwork/eth-contracts/go_abi/eccm_abi"
	"math/big"
	"strings"
	"testing"
)

func TestInvokeCrossChain(t *testing.T) {
	client := DefaultEthereumClient()
	ctx := context.Background()
	contractabi, err := abi.JSON(strings.NewReader(eccm_abi.EthCrossChainManagerABI))
	if err != nil {
		fmt.Printf("TestInvokeContract - err:" + err.Error())
		return
	}
	toAddr := common.HexToAddress("0x5cD3143f91a13Fe971043E1e4605C1c23b46bF44")
	txData, err := contractabi.Pack("crossChain", uint64(3), toAddr.Bytes(), toAddr.Bytes(), toAddr.Bytes())
	if err != nil {
		fmt.Printf("TestInvokeContract - err:" + err.Error())
		return
	}
	fmt.Printf("TestInvokeContract - txdata:%s\n", hex.EncodeToString(txData))

	contractAddr := common.HexToAddress("726532586C50ec9f4080B71f906a3d9779bbd64F")
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

func TestInvokeCrossChain1(t *testing.T) {
	client := DefaultEthereumClient()
	privateKey := DefaultPrivateKey()
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	ctx := context.Background()
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(client.GetNonceAt(ctx, address)))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(5000000) // in units
	auth.GasPrice = big.NewInt(30000000000)

	contractAddr := "726532586C50ec9f4080B71f906a3d9779bbd64F"
	lockAddress := common.HexToAddress(contractAddr)
	eccm, err := eccm_abi.NewEthCrossChainManager(lockAddress, client.Client)

	tx, err := eccm.CrossChain(auth, 3, []byte{}, []byte{}, []byte{})
	if err != nil {
		fmt.Printf("deploy contract - err: %s", err.Error())
	}
	fmt.Printf("deploy contract - tx: %s", tx.Hash().String())
}
