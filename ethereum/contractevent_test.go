package ethereum

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/blockchainpro/ethereum/contractabi/eccm"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	mccommon "github.com/ontio/multi-chain/common"
	"math/big"
	"testing"
)

func ParserCrossChainRawData(raw []byte) {
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
	ParserMethodArgs(args)
}

func ParserMethodArgs(args []byte) {
	source := mccommon.NewZeroCopySource(args)
	assethash, _ := source.NextVarBytes()
	address1, _ := mccommon.AddressParseFromBytes(assethash)
	toaddress, _ := source.NextVarBytes()
	addrsss2, _ := mccommon.AddressParseFromBytes(toaddress)

	amountBytesReverse, _ := source.NextHash()
	amountBytes := mccommon.ToArrayReverse(amountBytesReverse[:])
	amountBigInt := big.NewInt(0)
	amountBigInt.SetBytes(amountBytes)
	amount := amountBigInt.Uint64()

	fmt.Printf("toassethash: %s, toaddress: %s, toamount: %d\n",
		address1.ToHexString(), addrsss2.ToHexString(), amount)
}

func TestContractEvent(t *testing.T) {
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

	height := uint64(8022631)
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
		ParserCrossChainRawData(raw)
	}
}
