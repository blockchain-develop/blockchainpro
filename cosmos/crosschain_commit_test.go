package cosmos

import (
	"encoding/hex"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/rootmulti"
	"github.com/tendermint/tendermint/crypto/merkle"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/types"
	"testing"
	"time"
)

type CosmosHeader struct {
	Header     *types.Header
	Commit     *types.Commit
	Valsets    []*types.Validator
}

type CosmosProofValue struct {
	Kp           string
	Value        []byte
}

func TestHeaderCDC(t *testing.T) {
	c := NewHTTPClient()
	cdc := NewCDC()
	height := int64(100)
	validatorSet := GetValidatorSet(c, height)
	signedHeader := GetSignedHeader(c, height)
	var commitHeader CosmosHeader
	commitHeader.Header = signedHeader.Header
	commitHeader.Valsets = validatorSet
	commitHeader.Commit = signedHeader.Commit

	headerBs, err := cdc.MarshalBinaryBare(commitHeader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("cosmos block header, height: %d,\n%s\n", height, hex.EncodeToString(headerBs))
}

func TestStoreProofCDC(t *testing.T) {
	c := NewHTTPClient()
	cdc := NewCDC()
	addr := AddressFromBech32("cosmos1pkrpxp6rdjxskvfagxdg3pfwwjvxkp2hc2g72e")
	key := append([]byte{0x01}, addr.Bytes()...)
	path := "/store/acc/key"
	_pres, err := c.ABCIQueryWithOptions(path, key, client.ABCIQueryOptions{Prove: true})
	if err != nil {
		panic(err)
	}
	data := _pres.Response.String()
	fmt.Printf("Query Height: %d, Result: %s\n", _pres.Response.Height, string(data))
	time.Sleep(time.Second * 5)
	block := GetBlock(c, _pres.Response.Height + 1)
	if _pres.Response.Value != nil {
		storeName, err := parseQueryStorePath(path)
		if err != nil {
			panic(err)
		}
		kp := merkle.KeyPath{}
		kp = kp.AppendKey([]byte(storeName), merkle.KeyEncodingURL)
		kp = kp.AppendKey(_pres.Response.Key, merkle.KeyEncodingURL)

		prt := rootmulti.DefaultProofRuntime()
		err = prt.VerifyValue(_pres.Response.Proof, block.AppHash, kp.String(), _pres.Response.Value)
		if err != nil {
			panic(err)
		}
		fmt.Printf("verify store proof successful.\n")

		var proofValue CosmosProofValue
		proofValue.Kp = kp.String()
		proofValue.Value = _pres.Response.Value
		proofvalueBs, err := cdc.MarshalBinaryBare(&proofValue)
		if err != nil {
			panic(err)
		}
		fmt.Printf("proof value is : %s\n", hex.EncodeToString(proofvalueBs))
	} else {
		prt := rootmulti.DefaultProofRuntime()
		err = prt.VerifyAbsence(_pres.Response.Proof, block.AppHash, string(_pres.Response.Key))
		if err != nil {
			panic(err)
		}
		fmt.Printf("verify store proof successful.")
	}
}
