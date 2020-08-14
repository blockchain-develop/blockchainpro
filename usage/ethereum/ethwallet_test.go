package ethereum

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pborman/uuid"
	"io/ioutil"
	"testing"
)

func TestEthWallet(t *testing.T) {
	keyStoreFile := ""
	password := ""
	keyStore, err := ioutil.ReadFile(keyStoreFile)
	if err != nil {
		panic(err)
	}
	key, err := keystore.DecryptKey(keyStore, password)
	if err != nil {
		panic(err)
	}

	privKey := key.PrivateKey
	pubKey := &key.PrivateKey.PublicKey
	address := key.Address

	id := uuid.NewRandom()
	newKey := &keystore.Key{
		Id:         id,
		Address:    crypto.PubkeyToAddress(privKey.PublicKey),
		PrivateKey: privKey,
	}

	json, err := newKey.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Printf("key json: %s\n", string(json))

	privKeyBytes := crypto.FromECDSA(privKey)
	fmt.Printf("private key: %s\n", hex.EncodeToString(privKeyBytes))

	pubKeyBytes := crypto.FromECDSAPub(pubKey)
	fmt.Printf("public key: %s\n", hex.EncodeToString(pubKeyBytes))

	fmt.Printf("eth address: %s\n", address)
}
