package switcheo

import (
	"errors"
	"fmt"
	"github.com/tendermint/tendermint/rpc/client/http"
	tmtypes "github.com/tendermint/tendermint/types"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/polynetwork/cosmos-poly-module/btcx"
	"github.com/polynetwork/cosmos-poly-module/ccm"
	"github.com/polynetwork/cosmos-poly-module/ft"
	"github.com/polynetwork/cosmos-poly-module/headersync"
	"github.com/polynetwork/cosmos-poly-module/lockproxy"
)

func NewSwitcheoClient() *http.HTTP {
	config := types.GetConfig()
	config.SetBech32PrefixForAccount("swth", "swthpub")
	config.SetBech32PrefixForValidator("swthvaloper", "swthvaloperpub")
	config.SetBech32PrefixForConsensusNode("swthvalcons", "swthvalconspub")
	c, err := http.New("http://175.41.151.35:26657", "/websocket")
	if err != nil {
		fmt.Printf("new http failed, err: %s\n", err.Error())
	}
	return c
}

func NewCDC() *codec.Codec {
	cdc := codec.New()
	bank.RegisterCodec(cdc)
	types.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	auth.RegisterCodec(cdc)
	btcx.RegisterCodec(cdc)
	ccm.RegisterCodec(cdc)
	ft.RegisterCodec(cdc)
	headersync.RegisterCodec(cdc)
	lockproxy.RegisterCodec(cdc)
	return cdc
}

func GetBlock(c *http.HTTP, height int64) *tmtypes.Block {
	blockResult, err := c.Block(&height)
	if err != nil {
		panic(err)
	}
	return blockResult.Block
}

func GetValidatorSet(c *http.HTTP, height int64) []*tmtypes.Validator {
	validatorsResult, err := c.Validators(&height, 0, 0)
	if err != nil {
		panic(err)
	}
	return validatorsResult.Validators
}

func GetCommit(c *http.HTTP, height int64) *tmtypes.Commit {
	commitResult, err := c.Commit(&height)
	if err != nil {
		panic(err)
	}
	return commitResult.Commit
}

func GetSignedHeader(c *http.HTTP, height int64) tmtypes.SignedHeader{
	commitResult, err := c.Commit(&height)
	if err != nil {
		panic(err)
	}
	return commitResult.SignedHeader
}

func AddressFromBech32(address string) types.AccAddress {
	addr, err := types.AccAddressFromBech32(address)
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