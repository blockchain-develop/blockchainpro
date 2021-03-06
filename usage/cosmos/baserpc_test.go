package cosmos

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/rootmulti"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/tendermint/tendermint/crypto/merkle"
	"github.com/tendermint/tendermint/rpc/client"
	"testing"
	"time"
)

func TestCurrentHeight(t *testing.T) {
	c := NewHTTPClient()
	s, err := c.Status()
	if err != nil {
		panic(err)
	}
	height := s.SyncInfo.LatestBlockHeight
	fmt.Printf("current block height: %d\n", height)
}

func TestGetBlock(t *testing.T) {
	c := NewHTTPClient()
	height := int64(10000)
	block, err := c.Block(&height)
	if err != nil {
		panic(err)
	}
	fmt.Printf("height: %d, block: %s\n", height, block.Block.String())
}

func TestGetValidatorSet(t *testing.T) {
	c := NewHTTPClient()
	height := int64(100)
	validators, err := c.Validators(&height, 0, 0)
	if err != nil {
		panic(err)
	}
	xx, _ := json.Marshal(validators)
	fmt.Printf("height: %d, validators: %s\n", height, string(xx))
}

func TestAddressFromBech32(t *testing.T) {
	address := "cosmos1rfu9elzahmpwz5vwrvwndy256r8909jqtvwhfq"
	addr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		panic(err)
	}
	fmt.Printf("New address: %s, %s, %s", address, addr.String(), hex.EncodeToString(addr.Bytes()))
}

func TestAddressFromHex(t *testing.T) {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("swth", "swthpub")
	config.SetBech32PrefixForValidator("swthvaloper", "swthvaloperpub")
	config.SetBech32PrefixForConsensusNode("swthvalcons", "swthvalconspub")

	address := "1a785cfc5dbec2e1518e1b1d369154d0ce579640"
	addr, err := sdk.AccAddressFromHex(address)
	if err != nil {
		panic(err)
	}
	fmt.Printf("New address: %s, %s\n", address, addr.String())
}

func TestGetAccount(t *testing.T) {
	addr := AddressFromBech32("cosmos1pkrpxp6rdjxskvfagxdg3pfwwjvxkp2hc2g72e")
	cdc := NewCDC()
	c := NewHTTPClient()
	raw, err := cdc.MarshalJSON(auth.NewQueryAccountParams(addr))
	if err != nil {
		panic(err)
	}
	_pres, err := c.ABCIQueryWithOptions("/custom/acc/account", raw, client.ABCIQueryOptions{})
	if err != nil {
		panic(err)
	}
	data := _pres.Response.String()
	fmt.Printf("GetBalance Height: %d, Result: %s", _pres.Response.Height, string(data))

	var account exported.Account
	err = cdc.UnmarshalJSON(_pres.Response.Value, &account)
	if err != nil {
		panic(err)
	}
	fmt.Printf("balances: %s\n", account.String())
}

func TestGetBalance1(t *testing.T) {
	addr := AddressFromBech32("cosmos1pkrpxp6rdjxskvfagxdg3pfwwjvxkp2hc2g72e")
	cdc := NewCDC()
	c := NewHTTPClient()
	param := bank.NewQueryBalanceParams(addr)
	raw, err := cdc.MarshalJSON(param)
	if err != nil {
		panic(err)
	}
	_pres, err := c.ABCIQueryWithOptions("/custom/bank/balances", raw, client.ABCIQueryOptions{Height:3090})
	if err != nil {
		panic(err)
	}
	data := _pres.Response.String()
	fmt.Printf("GetBalance Height: %d, Result: %s\n", _pres.Response.Height, string(data))

	var coins sdk.Coins
	cdc.UnmarshalJSON(_pres.Response.Value, &coins)
	fmt.Printf("balances: %s\n", coins.String())
}

func TestQueryStore(t *testing.T) {
	c := NewHTTPClient()
	addr := AddressFromBech32("cosmos1pkrpxp6rdjxskvfagxdg3pfwwjvxkp2hc2g72e")
	key := append([]byte{0x01}, addr.Bytes()...)
	path := "/store/acc/key"
	//_pres, err := c.ABCIQueryWithOptions(path, key, client.ABCIQueryOptions{Prove: true, Height: 2256})
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
		fmt.Printf("verify store proof successful.")
	} else {
		prt := rootmulti.DefaultProofRuntime()
		err = prt.VerifyAbsence(_pres.Response.Proof, block.AppHash, string(_pres.Response.Key))
		if err != nil {
			panic(err)
		}
		fmt.Printf("verify store proof successful.")
	}
}

func TestGetTx(t *testing.T) {
	c := NewHTTPClient()
	//hash,_ := hex.DecodeString(utils.HexStringReverse("27E655EB3C87F06FD231103AD50D1DF707E88F00D3193267E7CBCFB88A03BB70"))
	hash,_ := hex.DecodeString("27E655EB3C87F06FD231103AD50D1DF707E88F00D3193267E7CBCFB88A03BB70")
	tx, err := c.Tx(hash, false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("tx: %v\n", tx)
}
