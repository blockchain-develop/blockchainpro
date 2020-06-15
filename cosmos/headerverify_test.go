package cosmos

import (
	"bytes"
	"fmt"
	"github.com/tendermint/tendermint/lite"
	"github.com/tendermint/tendermint/types"
	"testing"
)

func TestVerifyHeader(t *testing.T) {
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

func TestVerifyHeader1(t *testing.T) {
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

func TestVerifyHeader2(t *testing.T) {
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

func TestVerifyHeader3(t *testing.T) {
	chainId := "testing"
	height := int64(100)
	c := NewHTTPClient()
	preBlock := GetBlock(c, height - 1)
	validatorSet := GetValidatorSet(c, height)
	valset := types.NewValidatorSet(validatorSet)
	signedHeader := GetSignedHeader(c, height)
	if bytes.Equal(preBlock.NextValidatorsHash, valset.Hash()) != true {
		panic("next validator hash is not right!\n")
	}
	if bytes.Equal(signedHeader.Header.ValidatorsHash, valset.Hash()) != true {
		panic("validator hash is not right!\n")
	}
	if signedHeader.Commit.GetHeight() != signedHeader.Height {
		panic("commit height is not right!\n")
	}
	if bytes.Equal(signedHeader.Commit.BlockID.Hash, signedHeader.Header.Hash()) != true {
		panic("commit hash is not right!\n")
	}
	if err := signedHeader.Commit.ValidateBasic(); err != nil {
		panic(err)
	}
	if valset.Size() != len(signedHeader.Commit.Signatures) {
		panic("the size of precommits is not right!")
	}
	talliedVotingPower := int64(0)
	for idx, commitSig := range signedHeader.Commit.Signatures {
		if commitSig.Absent() {
			continue // OK, some precommits can be missing.
		}
		_, val := valset.GetByIndex(idx)
		// Validate signature.
		precommitSignBytes := signedHeader.Commit.VoteSignBytes(chainId, idx)
		if !val.PubKey.VerifyBytes(precommitSignBytes, commitSig.Signature) {
			panic("Invalid commit -- invalid signature!")
		}
		// Good precommit!
		if signedHeader.Commit.BlockID.Equals(commitSig.BlockID(signedHeader.Commit.BlockID)) {
			talliedVotingPower += val.VotingPower
		}
	}

	if talliedVotingPower <= valset.TotalVotingPower()*2/3 {
		panic("voteing power is not enough!")
	}
	fmt.Printf("verify header successful!\n")
}

