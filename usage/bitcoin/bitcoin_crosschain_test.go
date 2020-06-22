package bitcoin

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/ontio/multi-chain/common"
	"testing"
)

func ParseCrossTransfer(rawTx *btcjson.TxRawResult) (xtype uint32, address []byte, tchainid uint64, fee uint64, extData []byte, err error) {
	extData, err = hex.DecodeString(rawTx.Vout[1].ScriptPubKey.Hex)
	if err != nil {
		return 0, nil, 0, 0, nil, err
	}
	source := common.NewZeroCopySource(extData[3:])
	tchainid, eof := source.NextUint64()
	if eof {
		return
	}
	fee, eof = source.NextUint64()
	if eof {
		return
	}
	address, eof = source.NextVarBytes()
	if eof {
		return
	}
	xtype = 1
	return
}

func TestCrossChainMsg(t *testing.T) {
	//btc := NewBtcTools("http://172.168.3.10:20336", "test", "test")
	btc := NewBtcTools("http://18.140.187.37:18332", "omnicorerpc", "EzriglUqnFC!")
	txhash := "a0ee3c9d499fcc834ee9b63ad44dd7aca4dac277dfb0c6919f33a6a485f3b230"
	btctx, err := btc.GetTx(txhash)
	if err != nil {
		fmt.Printf("get tx err: %v\n", err)
		return
	}
	xtype, address, tchainid, fee, _, err := ParseCrossTransfer(btctx)
	fmt.Printf("xtype: %d, address: %s, tchainid: %d, fee: %d\n", xtype, hex.EncodeToString(address), tchainid, fee)

	rk := GetUtxoKey1(&btctx.Vout[0].ScriptPubKey)
	fmt.Printf("redeemkey: %s\n", rk)
}
