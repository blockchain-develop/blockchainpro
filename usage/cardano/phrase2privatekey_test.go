package cardano

import (
	"encoding/hex"
	"fmt"
	"github.com/coinbase/rosetta-sdk-go/keys"
	"github.com/coinbase/rosetta-sdk-go/types"
	"testing"
)

func TestPhrase2PrivateKey(t *testing.T) {
	keypair, _ := keys.GenerateKeypair(types.Edwards25519)
	prikey := hex.EncodeToString(keypair.PrivateKey)
	pubkey := hex.EncodeToString(keypair.PublicKey.Bytes)
	fmt.Printf("private key: %s, public key: %s\n", prikey, pubkey)
	// private key: d0a6e05eefdc19d965821c6540200448e5ee5bc01eb872db80ffe5ab497e3e31,
	// public key: d112db6be0349f95950279f3d6e0b8012e5c54d81269a126e6ce742eaeaf9a49
	// addr_test1vryt3x0dhuya0uw5l7r0c0lqjs4zvdg4ludcc5q4pw3cfwqx6f5yr


	// private key: 3b179d1d8804a4b8750a7ae3edd9ba19ef6567332ca8cd746abd81f7c1e5eebc,
	// public key: aa2f420ba77275c41983ce78fea70b70778e1721e25ae5b3212ac4d04fcae703
	// addr_test1vpesqj2kchqvprwg7xaj3rtr62xkqm74qhf0h7c47pkr98qgnmh2x

	// private key: d266cbf18f345bb21e362b319a19c185a590170bdf4b5d2e5173f3f817196d9e,
	// public key: 0858d5d6bb5ec1d25e4f719ca44190db2e9d18cbe8a803d06231cb3628ef56b8
	// addr_test1vqm0wdwt5rut2h6jxg7rure0smsgr6lf05x8hwacvd2amscanamxg
}


func TestWif2PrivateKey(t *testing.T) {
	privateKey, err := keys.ImportPrivateKey("45A04AC83247399C337A26A69567E022E20D3A4BF1A0908411354EEBE56F532F", types.Edwards25519)
	//privateKey, err := keys.ImportPrivateKey("41d9523b87b9bd89a4d07c9b957ae68a7472d8145d7956a692df1a8ad91957a2c117d9dd874447f47306f50a650f1e08bf4bec2cfcb2af91660f23f2db912977", types.Edwards25519)
	if err != nil {
		panic(err)
	}
	fmt.Printf("private key: %s, public key: %s\n",
		hex.EncodeToString(privateKey.PrivateKey), hex.EncodeToString(privateKey.PublicKey.Bytes))
}
