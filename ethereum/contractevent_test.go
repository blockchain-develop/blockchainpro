package ethereum

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/blockchainpro/ethereum/contractabi/eccm"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joeqian10/neo-gogogo/helper"
	mccommon "github.com/ontio/multi-chain/common"
	ontcommon "github.com/ontio/ontology/common"
	"math/big"
	"testing"
)

func ParserCrossChainRawData(name string, raw []byte) {
	source := mccommon.NewZeroCopySource(raw)
	txindex, _ := source.NextVarBytes()
	txhash, _ := source.NextVarBytes()
	sender, _ := source.NextVarBytes()
	tochainid, _ := source.NextUint64()
	tocontract, _ := source.NextVarBytes()
	method, _ := source.NextVarBytes()
	args, _ := source.NextVarBytes()
	fmt.Printf("index: %s, txhash: %s, sender: %s, tochainid: %d, tocontract: %s, method: %s, args: %s\n",
		hex.EncodeToString(txindex), hex.EncodeToString(txhash), hex.EncodeToString(sender), tochainid, hex.EncodeToString(tocontract),hex.EncodeToString(method),hex.EncodeToString(args))

	//
	ParserMethodArgs(name, tochainid, args)
}

func ParserMethodArgs(name string, toChainId uint64, args []byte) {
	source1 := mccommon.NewZeroCopySource(args)
	var assetaddress []byte
	var toaddress []byte
	var amount uint64
	if name == "btc" {
		toaddress, _ = source1.NextVarBytes()
		amount, _ = source1.NextUint64()
	} else {
		assetaddress, _ = source1.NextVarBytes()
		toaddress, _ = source1.NextVarBytes()
		amountBytesReverse, _ := source1.NextHash()
		amountBytes := mccommon.ToArrayReverse(amountBytesReverse[:])
		amountBigInt := big.NewInt(0)
		amountBigInt.SetBytes(amountBytes)
		amount = amountBigInt.Uint64()
	}
	var assetaddress2 string
	var toaddress2 string
	if toChainId == 1 {
		assetaddress2 = "0000000000000000000000000000000000000011"
		toaddress2 = hex.EncodeToString(toaddress)
		//toaddress2 = base58.Encode(toaddress)
	} else if toChainId == 3 {
		assetaddress1, _ := ontcommon.AddressParseFromBytes(assetaddress)
		toaddress1, _ := ontcommon.AddressParseFromBytes(toaddress)
		assetaddress2 = assetaddress1.ToHexString()
		toaddress2 = toaddress1.ToHexString()
	} else if toChainId == 4 {
		assetaddress1, _ := helper.UInt160FromBytes(assetaddress)
		toaddress1, _ := helper.UInt160FromBytes(toaddress)
		assetaddress2 = helper.ScriptHashToAddress(assetaddress1)
		toaddress2 = helper.ScriptHashToAddress(toaddress1)
	} else {
		assetaddress2 = ""
		toaddress2 = ""
	}
	fmt.Printf("to chain id: %d, to asset address: %s, to address: %s, amount: %d\n", toChainId, assetaddress2, toaddress2, amount)
}


// nok
func TestCrossChainEvent_ETH2ONT(t *testing.T) {
	url := "http://18.139.17.85:10331"
	ethclient, err := ethclient.Dial(url)
	if err != nil {
		fmt.Printf("getmocktokeninfo - cannot dial sync node, err: %s", err)
		return
	}

	addressString := "bA6F835ECAE18f5Fc5eBc074e5A0B94422a13126"
	eccmContractAddr := common.HexToAddress(addressString)
	eccmContract, err := eccm.NewEthCrossChainManager(eccmContractAddr, ethclient)
	if err != nil {
		fmt.Printf("NewEthCrossChainManager error: %v\n", err)
		return
	}

	//height := uint64(8022631)
	height := uint64(8029414)
	opt := &bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: context.Background(),
	}

	// get ethereum lock events from given block
	lockEvents, err := eccmContract.FilterCrossChainEvent(opt, nil)
	if err != nil {
		fmt.Printf("FilterCrossChainEvent error: %v\n", err)
		return
	}

	for lockEvents.Next() {
		evt := lockEvents.Event
		fmt.Printf("txid: %s, txhash: %s, send address: %s, tochainid: %d, tocontract: %s, height: %d, ProxyOrAssetContract: %s\n",
			hex.EncodeToString(evt.TxId), evt.Raw.TxHash.String(), evt.Sender.String(), evt.ToChainId,
			hex.EncodeToString(evt.ToContract), height, evt.ProxyOrAssetContract.String())
		raw := evt.Rawdata
		ParserCrossChainRawData("btc", raw)
	}
}


// nok
func TestCrossChainEvent_ETH2BTC(t *testing.T) {
	url := "http://18.139.17.85:10331"
	ethclient, err := ethclient.Dial(url)
	if err != nil {
		fmt.Printf("getmocktokeninfo - cannot dial sync node, err: %s", err)
		return
	}

	addressString := "bA6F835ECAE18f5Fc5eBc074e5A0B94422a13126"
	eccmContractAddr := common.HexToAddress(addressString)
	eccmContract, err := eccm.NewEthCrossChainManager(eccmContractAddr, ethclient)
	if err != nil {
		fmt.Printf("NewEthCrossChainManager error: %v\n", err)
		return
	}

	height := uint64(8029428)
	opt := &bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: context.Background(),
	}

	// get ethereum lock events from given block
	lockEvents, err := eccmContract.FilterCrossChainEvent(opt, nil)
	if err != nil {
		fmt.Printf("FilterCrossChainEvent error: %v\n", err)
		return
	}

	for lockEvents.Next() {
		evt := lockEvents.Event
		fmt.Printf("txid: %s, txhash: %s, send address: %s, tochainid: %d, tocontract: %s, height: %d, ProxyOrAssetContract: %s\n",
			hex.EncodeToString(evt.TxId), evt.Raw.TxHash.String(), evt.Sender.String(), evt.ToChainId,
			hex.EncodeToString(evt.ToContract), height, evt.ProxyOrAssetContract.String())
		raw := evt.Rawdata
		ParserCrossChainRawData("btc", raw)
	}
}
