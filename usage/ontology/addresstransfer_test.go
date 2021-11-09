package ontology

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/itchyny/base58-go"
	"github.com/ontio/ontology-crypto/keypair"
	"github.com/ontio/ontology-crypto/signature"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/core/types"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/ripemd160"
	"math/big"
	"testing"
)

func TestHex2Base58(t *testing.T) {
	//addrHexStr := "616f2a4a38396ff203ea01e6c070ae421bb8ce2d"
	addrHexStr := "f3b8a17f1f957f60c88f105e32ebff3f022e56a4"
	addr, err := common.AddressFromHexString(addrHexStr)
	if err != nil {
		fmt.Printf("AddressFromHexString err: %s", err.Error())
	}
	addrBase58 := addr.ToBase58()
	fmt.Printf("ontology address, hex: %s, base58: %s\n", addrHexStr, addrBase58)
}

func TestBase582Hex(t *testing.T) {
	addrBase58 := "AHe5NrFRBWYaJo9uB5iQkViXQ7naqQ8y6a"
	addr, err := common.AddressFromBase58(addrBase58)
	if err != nil {
		fmt.Printf("AddressFromBase58 err: %s", err.Error())
	}
	addrHexStr := addr.ToHexString()
	fmt.Printf("ontology address, hex: %s, base58: %s\n", addrHexStr, addrBase58)
}

func TestPubkey2Address(t *testing.T) {
	hexpubkey := "317c350d8ea8c8e25360cd455da2366d58186c61390c01ec8d372971a48fba4e"
	data, err := hex.DecodeString(hexpubkey)
	if err != nil {
		panic(err)
	}
	if len(data) != ed25519.PublicKeySize {
		panic("deserializing public key failed: invalid length for Ed25519 public key")
	}
	pubkey := ed25519.PublicKey(data)
	addr := types.AddressFromPubKey(pubkey)
	fmt.Printf("pubkey hex: %s, address: %s\n", hexpubkey, addr.ToBase58())
}

func TestPubkey2Address2(t *testing.T) {
	hexpubkey := "317c350d8ea8c8e25360cd455da2366d58186c61390c01ec8d372971a48fba4e"
	data, err := hex.DecodeString(hexpubkey)
	if err != nil {
		panic(err)
	}

	if len(data) != ed25519.PublicKeySize {
		panic("deserializing public key failed: invalid length for Ed25519 public key")
	}
	//
	{
		temp := make([]byte, 36)
		temp[0] = 34
		temp[1] = byte(keypair.PK_EDDSA)
		temp[2] = keypair.ED25519
		copy(temp[3:32+3], data)
		temp[35] = 172

		var addr1 common.Address
		temp1 := sha256.Sum256(temp)
		md := ripemd160.New()
		md.Write(temp1[:])
		md.Sum(addr1[:0])

		fmt.Printf("adress: %s\n", addr1.ToBase58())
	}

	{
		temp := make([]byte, 36)
		{
			temp[0] = 34
			temp[1] = byte(keypair.PK_EDDSA)
			temp[2] = keypair.ED25519
			copy(temp[3:32+3], data)
			temp[35] = 172
		}

		var addr1 common.Address
		{
			temp1 := sha256.Sum256(temp)
			md := ripemd160.New()
			md.Write(temp1[:])
			md.Sum(addr1[:0])
		}

		var addr string
		{
			temp2 := append([]byte{23}, addr1[:]...)
			temp3 := sha256.Sum256(temp2)
			temp4 := sha256.Sum256(temp3[:])
			temp2 = append(temp2, temp4[0:4]...)

			bi := new(big.Int).SetBytes(temp2).String()
			encoded, _ := base58.BitcoinEncoding.Encode([]byte(bi))
			addr = string(encoded)
		}

		fmt.Printf("adress: %s\n", addr)
		{
			if addr == "" || len(addr) > 2048 {
				panic(err)
			}
			decoded, err := base58.BitcoinEncoding.Decode([]byte(addr))
			if err != nil {
				panic(err)
			}

			x, ok := new(big.Int).SetString(string(decoded), 10)
			if !ok {
				panic(err)
			}

			buf := x.Bytes()
			if len(buf) != 25 || buf[0] != byte(23) {
				panic(err)
			}

			addx := buf[1:21]

			temp2 := append([]byte{23}, addx[:]...)
			temp3 := sha256.Sum256(temp2)
			temp4 := sha256.Sum256(temp3[:])
			temp2 = append(temp2, temp4[0:4]...)

			bi := new(big.Int).SetBytes(temp2).String()
			encoded, _ := base58.BitcoinEncoding.Encode([]byte(bi))
			addr2 := string(encoded)

			if addr != addr2 {
				panic(err)
			}
		}
	}

	{
		temp := make([]byte, 34)
		temp[0] = 33
		copy(temp[1:32+1], data)
		temp[33] = 172

		var addr1 common.Address
		temp1 := sha256.Sum256(temp)
		md := ripemd160.New()
		md.Write(temp1[:])
		md.Sum(addr1[:0])

		fmt.Printf("adress: %s\n", addr1.ToBase58())
	}

	{
		pubkey := ed25519.PublicKey(data)
		addr := types.AddressFromPubKey(pubkey)
		fmt.Printf("pubkey hex: %s, address: %s\n", hexpubkey, addr.ToBase58())
	}
}

func Test25519Address(t *testing.T) {
	account := ontology_go_sdk.NewAccount(signature.SHA512withEDDSA)
	hexpubkey := hex.EncodeToString(account.PublicKey.(ed25519.PublicKey))
	data, err := hex.DecodeString(hexpubkey)
	if err != nil {
		panic(err)
	}
	if len(data) != ed25519.PublicKeySize {
		panic("deserializing public key failed: invalid length for Ed25519 public key")
	}
	pubkey := ed25519.PublicKey(data)
	addr := types.AddressFromPubKey(pubkey)
	fmt.Printf("ontology address: %s, pubkey hex: %s, address: %s\n", account.Address.ToBase58(), hexpubkey, addr.ToBase58())
}

func Test256Address(t *testing.T) {
	account := ontology_go_sdk.NewAccount()
	addr := types.AddressFromPubKey(account.PublicKey)
	fmt.Printf("ontology address: %s, address: %s\n", account.Address.ToBase58(), addr.ToBase58())
}
