package ontology

import (
	"fmt"
	"testing"
)

func TestOep4(t *testing.T) {
	sdk := newOntologySdk()
	contract := "11c60400f54c17df0d8e9c4a38c333b66c1f1c54"
	name, err := OEP4Name(sdk, contract)
	if err != nil {
		panic(err)
	}
	symbol, err := OEP4Symbol(sdk, contract)
	if err != nil {
		panic(err)
	}
	decimal, err := OEP4Decimals(sdk, contract)
	if err != nil {
		panic(err)
	}
	supply, err := OEP4Supply(sdk, contract)
	if err != nil {
		panic(err)
	}
	fmt.Printf("oep4: %s, name: %s, symbol: %s, decimal: %d, supply: %s\n", contract, name, symbol, decimal, supply)
}
