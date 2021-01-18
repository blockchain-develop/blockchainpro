package poly

import (
	"fmt"
	"github.com/ontio/ontology-crypto/keypair"
	poly_go_sdk "github.com/polynetwork/poly-go-sdk"
	"github.com/polynetwork/poly/common"
)

func NewSdk(url string) *poly_go_sdk.PolySdk {
	sdk := poly_go_sdk.NewPolySdk()
	sdk.NewRpcClient().SetAddress(url)
	return sdk
}

func createPolyAccount() {
	ontsdk := poly_go_sdk.NewPolySdk()
	var wallet *poly_go_sdk.Wallet
	var err error
	if !common.FileExisted("./wallet_ontology_new.dat") {
		wallet, err = ontsdk.CreateWallet("./wallet_ontology_new.dat")
		if err != nil {
			return
		}
	} else {
		wallet, err = ontsdk.OpenWallet("./wallet_ontology_new.dat")
		if err != nil {
			fmt.Errorf("createOntologyAccount - wallet open error: %s", err.Error())
			return
		}
	}
	signer, err := wallet.GetDefaultAccount([]byte("1"))
	if err != nil || signer == nil {
		signer, err = wallet.NewDefaultSettingAccount([]byte("1"))
		if err != nil {
			fmt.Errorf("createOntologyAccount - wallet password error")
			return
		}
		err = wallet.Save()
		if err != nil {
			return
		}
	}
	pri_key, _ := keypair.Key2WIF(signer.PrivateKey)
	addr := signer.Address.ToBase58()
	fmt.Printf("private key: %s, address: %s %s\n", string(pri_key), addr, signer.Address.ToHexString())
}
