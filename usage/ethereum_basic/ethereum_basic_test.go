package ethereum

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
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

func TestNonce(t *testing.T) {
	client := DefaultEthereumClient()
	addr := common.HexToAddress("0x733679CC9Ddc38f78Ad121E2326E61D20f416692")
	nonce := client.GetNonceAt(context.Background(), addr)
	code := client.GetCodeAt(context.Background(), addr)
	fmt.Printf("address: %s, nonce: %d, code: %s\nn", addr.String(), nonce, code)
}

func TestXXX(t *testing.T) {
	//client := DefaultEthereumClient()
	data, _ := hex.DecodeString("02f8d483aa36a70184773594008502540be4008309f5c694ee96c2d90f63249502b857ab97cf4f6e015b92c080b864abb8f4bd0000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000007a120000000000000000000000000000000000000000000000000000000000000001dc080a094c0d054aa46de6800a0e4e1e23446b19dcad82dbaff35c054ddf43e51381765a0130905a6bef57d19f850348d5c12f4e83ec958dc2536d8bf1e187313a50a6b84")
	var tx types.Transaction
	err := tx.UnmarshalBinary(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("tx: %v\n", tx)
	//
	chainId := big.NewInt(11155111)
	signer := types.NewLondonSigner(chainId)
	from, err := signer.Sender(&tx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("from: %s\n", from.String())
}
