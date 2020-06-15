package cosmos

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/crypto/multisig"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

func NewHTTPClient() *client.HTTP {
	c, err := client.NewHTTP("http://172.168.3.93:26657", "/websocket")
	//c := rpcclient.NewHTTP("https://lcd.nylira.net:26657", "/websocket")
	if err != nil {
		fmt.Printf("new http failed, err: %s\n", err.Error())
	}
	return c
}

func NewCDC() *codec.Codec {
	cdc := codec.New()

	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	cdc.RegisterConcrete(ed25519.PubKeyEd25519{}, ed25519.PubKeyAminoName, nil)
	cdc.RegisterConcrete(secp256k1.PubKeySecp256k1{}, secp256k1.PubKeyAminoName, nil)
	cdc.RegisterConcrete(multisig.PubKeyMultisigThreshold{}, multisig.PubKeyMultisigThresholdAminoRoute, nil)

	cdc.RegisterInterface((*crypto.PrivKey)(nil), nil)
	cdc.RegisterConcrete(ed25519.PrivKeyEd25519{}, ed25519.PrivKeyAminoName, nil)
	cdc.RegisterConcrete(secp256k1.PrivKeySecp256k1{}, secp256k1.PrivKeyAminoName, nil)
	return cdc
}

func GetBlock(c *client.HTTP, height int64) *types.Block {
	blockResult, err := c.Block(&height)
	if err != nil {
		panic(err)
	}
	return blockResult.Block
}

func GetValidatorSet(c *client.HTTP, height int64) []*types.Validator {
	validatorsResult, err := c.Validators(&height, 0, 0)
	if err != nil {
		panic(err)
	}
	return validatorsResult.Validators
}

func GetCommit(c *client.HTTP, height int64) *types.Commit {
	commitResult, err := c.Commit(&height)
	if err != nil {
		panic(err)
	}
	return commitResult.Commit
}

func GetSignedHeader(c *client.HTTP, height int64) types.SignedHeader{
	commitResult, err := c.Commit(&height)
	if err != nil {
		panic(err)
	}
	return commitResult.SignedHeader
}

func AddressFromBech32(address string) sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		panic(err)
	}
	fmt.Printf("New address: %s, %s\n", address, addr.String())
	return addr
}


func parseQueryStorePath(path string) (storeName string, err error) {
	if !strings.HasPrefix(path, "/") {
		return "", fmt.Errorf("expected path to start with /")
	}

	paths := strings.SplitN(path[1:], "/", 3)
	switch {
	case len(paths) != 3:
		return "", errors.New("expected format like /store/<storeName>/key")
	case paths[0] != "store":
		return "", errors.New("expected format like /store/<storeName>/key")
	case paths[2] != "key":
		return "", errors.New("expected format like /store/<storeName>/key")
	}

	return paths[1], nil
}