package bitcoin

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"testing"
)

func TestGetCurrentBlockHeight(t *testing.T) {
	//btc := NewBtcTools("http://172.168.3.10:20336", "test", "test")
	btc := NewDefaultBtcTools()
	height, err := btc.GetCurrentHeight()
	if err != nil {
		fmt.Printf("GetCurrentHeight err: %v\n", err)
		return
	}
	fmt.Printf("current height: %d\n", height)
}

func TestBTCAddress(t *testing.T) {
	add_bytes, _ := hex.DecodeString("c330431496364497d7257839737b5e4596f5ac06")
	address, _ := btcutil.NewAddressScriptHashFromHash(add_bytes, &chaincfg.TestNet3Params)
	fmt.Printf("address: %s\n", address.EncodeAddress())
}

func TestBTCAddress1(t *testing.T) {
	add_bytes, _ := hex.DecodeString("c330431496364497d7257839737b5e4596f5ac06")
	address, _ := btcutil.NewAddressScriptHashFromHash(add_bytes, &chaincfg.TestNet3Params)
	fmt.Printf("address: %s\n", address)
}

func TestBTCTransaction(t *testing.T) {
	btc := NewDefaultBtcTools()
	btctx, err := btc.GetTx("3d8ec4452116caf59a2cd16983bed8858dcbfdaaae9d9b881ff170f6263eb8ac")
	if err != nil {
		panic(err)
	}
	toAddr := GetAddress(&btctx.Vout[0].ScriptPubKey)
	fmt.Printf("address: %s\n", toAddr)
	fmt.Printf("addresses: %v\n", btctx.Vout[0].ScriptPubKey.Addresses)
}
