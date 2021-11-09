package bitcoin

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"testing"
)


func TestTransactionEncodeDecode(t *testing.T) {
	bitcoinClient, err := DefaultBitcoinClient()
	if err != nil {
		panic(err)
	}
	//
	var h chainhash.Hash
	err = chainhash.Decode(&h, "badd529cfa3f9bb297d56ea2588e87efbff5d3cb921ae3920a7f6c6b07364c4d")
	if err != nil {
		panic(err)
	}
	fmt.Printf("hash: %s\n", h.String())
	//
	tx, err := bitcoinClient.GetRawTransactionVerbose(&h)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", tx)
}
