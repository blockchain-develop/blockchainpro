package bitcoin

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/golangcrypto/ripemd160"
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

func GetUtxoKey(scriptPk []byte) string {
	switch txscript.GetScriptClass(scriptPk) {
	case txscript.MultiSigTy:
		return hex.EncodeToString(btcutil.Hash160(scriptPk))
	case txscript.ScriptHashTy:
		return hex.EncodeToString(scriptPk[2:22])
	case txscript.WitnessV0ScriptHashTy:
		hasher := ripemd160.New()
		hasher.Write(scriptPk[2:34])
		return hex.EncodeToString(hasher.Sum(nil))
	default:
		return ""
	}
}

func GetUtxoKey1(scriptPk *btcjson.ScriptPubKeyResult) string {
	scriptPkBytes, _ := hex.DecodeString(scriptPk.Hex)
	switch scriptPk.Type {
	case "multisig":
		return hex.EncodeToString(btcutil.Hash160(scriptPkBytes))
	case "scripthash":
		return hex.EncodeToString(scriptPkBytes[2:22])
	case "witness_v0_scripthash":
		hasher := ripemd160.New()
		hasher.Write(scriptPkBytes[2:34])
		return hex.EncodeToString(hasher.Sum(nil))
	default:
		return ""
	}
}

func TestCrossChainMsg(t *testing.T) {
	btc := NewBtcTools("http://172.168.3.10:20336", "test", "test")
	txhash := "296da9b23d03108845d4cfacf920b42c88fd48be99e2d6d2b4ce5b77c55825b6"
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
