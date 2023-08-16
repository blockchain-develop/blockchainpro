package wallet

import (
	"fmt"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"testing"
)

func TestHDWallet1(t *testing.T) {
	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		panic(err)
	}
	{
		path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
		account, err := wallet.Derive(path, false)
		if err != nil {
			panic(err)
		}
		fmt.Printf("address: %s\n", account.Address.Hex())
	}
	{
		path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
		account, err := wallet.Derive(path, false)
		if err != nil {
			panic(err)
		}
		fmt.Printf("address: %s\n", account.Address.Hex())
	}
}
