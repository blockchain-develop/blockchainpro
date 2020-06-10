package bitcoin

import (
	"encoding/hex"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"golang.org/x/crypto/ripemd160"
)

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