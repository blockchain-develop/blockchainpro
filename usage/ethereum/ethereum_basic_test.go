package ethereum

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
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
