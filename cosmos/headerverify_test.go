package cosmos

import (
	"bytes"
	"fmt"
	"github.com/tendermint/tendermint/lite"
	"github.com/tendermint/tendermint/types"
	"testing"
)

func TestVrifyHeder(t *testing.T) {
	chainId := "testing"
	height := int64(100)
	c := NewHTTPClient()
	preBlock := GetBlock(c, height - 1)
	curBlock := GetBlock(c, height)
	nextBlock := GetBlock(c, height + 1)
	validatorSet := GetValidatorSet(c, height)
	valset := types.NewValidatorSet(validatorSet)
	verifier := lite.NewBaseVerifier(chainId, height, valset)
	signedHeader := types.SignedHeader{
		Header: &curBlock.Header,
		Commit: nextBlock.LastCommit,
	}
	if bytes.Equal(preBlock.NextValidatorsHash, valset.Hash()) != true {
		panic("block validator is not right!")
	}
	err := verifier.Verify(signedHeader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("verify header successful!")
}

func TestVrifyHeder1(t *testing.T) {
	chainId := "testing"
	height := int64(100)
	c := NewHTTPClient()
	preBlock := GetBlock(c, height - 1)
	curBlock := GetBlock(c, height)
	validatorSet := GetValidatorSet(c, height)
	valset := types.NewValidatorSet(validatorSet)
	verifier := lite.NewBaseVerifier(chainId, height, valset)
	signedHeader := types.SignedHeader{
		Header: &curBlock.Header,
		Commit: GetCommit(c, height),
	}
	if bytes.Equal(preBlock.NextValidatorsHash, valset.Hash()) != true {
		panic("block validator is not right!")
	}
	err := verifier.Verify(signedHeader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("verify header successful!")
}

func TestVrifyHeder2(t *testing.T) {
	chainId := "testing"
	height := int64(100)
	c := NewHTTPClient()
	preBlock := GetBlock(c, height - 1)
	validatorSet := GetValidatorSet(c, height)
	valset := types.NewValidatorSet(validatorSet)
	if bytes.Equal(preBlock.NextValidatorsHash, valset.Hash()) != true {
		panic("block validator is not right!")
	}
	verifier := lite.NewBaseVerifier(chainId, height, valset)
	signedHeader := GetSignedHeader(c, height)
	err := verifier.Verify(signedHeader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("verify header successful!")
}

