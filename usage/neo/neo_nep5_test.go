package neo

import (
	"fmt"
	"testing"
)

func TestNep5(t *testing.T) {
	nep5Hash := "74f2dc36a68fdc4682034178eb2220729231db76"
	nep5 := NewNep5(nep5Hash)
	decimal, err := nep5.Decimals()
	if err != nil {
		panic(err)
	}
	name, err := nep5.Name()
	if err != nil {
		panic(err)
	}
	fmt.Printf("nep5: %s, decimal: %d, name: %s\n", nep5Hash, decimal, name)
}
