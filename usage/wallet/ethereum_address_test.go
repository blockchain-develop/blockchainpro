package wallet

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

func TestEthereumAddress(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.PublicKey
	address := crypto.PubkeyToAddress(publicKey)
	fmt.Printf("%s\n", address)
}

func TestEthereumKeyStore(t *testing.T) {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "123456"
	account, err := ks.NewAccount(password)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", account.Address.String())
}
