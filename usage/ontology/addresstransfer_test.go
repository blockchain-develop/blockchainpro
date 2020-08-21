package ontology

import (
	"fmt"
	"github.com/ontio/ontology/common"
	"testing"
)

func TestHex2Base58(t *testing.T) {
	//addrHexStr := "616f2a4a38396ff203ea01e6c070ae421bb8ce2d"
	addrHexStr := "f3b8a17f1f957f60c88f105e32ebff3f022e56a4"
	addr, err := common.AddressFromHexString(addrHexStr)
	if err != nil {
		fmt.Printf("AddressFromHexString err: %s", err.Error())
	}
	addrBase58 := addr.ToBase58()
	fmt.Printf("ontology address, hex: %s, base58: %s\n", addrHexStr, addrBase58)
}

func TestBase582Hex(t *testing.T) {
	addrBase58 := "AHe5NrFRBWYaJo9uB5iQkViXQ7naqQ8y6a"
	addr, err := common.AddressFromBase58(addrBase58)
	if err != nil {
		fmt.Printf("AddressFromBase58 err: %s", err.Error())
	}
	addrHexStr := addr.ToHexString()
	fmt.Printf("ontology address, hex: %s, base58: %s\n", addrHexStr, addrBase58)
}
