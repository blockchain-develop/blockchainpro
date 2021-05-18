package ethereum

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/blockchainpro/usage/ethereum/contractabi/lock_proxy_abi"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
	"strings"
	"testing"
)

func TestKey2Address(t *testing.T) {
	privateKey := DefaultPrivateKey()
	publicKey := privateKey.PublicKey
	address := crypto.PubkeyToAddress(publicKey)
	fmt.Printf("private key: %s, public key: %s, address: %s\n",
		hex.EncodeToString(crypto.FromECDSA(privateKey)),
		hex.EncodeToString(crypto.FromECDSAPub(&privateKey.PublicKey)),
		address.String())
}

func TestTransactionEncode(t *testing.T) {
	client := DefaultEthereumClient()
	ctx := context.Background()
	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	if err != nil {
		fmt.Printf("TestTransactionEncode - err:" + err.Error())
		return
	}
	toAddr := common.HexToAddress("0x5cD3143f91a13Fe971043E1e4605C1c23b46bF44")
	txData, err := contractabi.Pack("lock", uint64(1), toAddr.Bytes(), big.NewInt(int64(1000000)))
	if err != nil {
		fmt.Printf("TestTransactionEncode - err:" + err.Error())
		return
	}
	fmt.Printf("TestTransactionEncode - txdata:%s\n", hex.EncodeToString(txData))

	contractAddr := common.HexToAddress("239100e629a9ca8e0bf45c7892b0fc72d78aa97a")
	fromAddr := common.HexToAddress("0x5cD3143f91a13Fe971043E1e4605C1c23b46bF44")
	tx := types.NewTransaction(client.GetNonceAt(ctx, fromAddr), contractAddr, big.NewInt(0), uint64(8000000), big.NewInt(30000000000), txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)
	fmt.Printf("TestTransactionEncode - tx is %s\n", hex.EncodeToString(bf.Bytes()))
}

func TestAddressEncode(t *testing.T) {
}

