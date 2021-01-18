package neo

import (
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/nep5"
	"github.com/joeqian10/neo-gogogo/rpc"
	"github.com/joeqian10/neo-gogogo/wallet"
)

func NewNeoClient() *rpc.RpcClient {
	rawClient := rpc.NewClient("http://seed8.ngd.network:10332")
	return rawClient
}

func NewNep5(hash string) *nep5.Nep5Helper {
	scriptHash, err := helper.UInt160FromString(hash)
	if err != nil {
		panic(err)
	}
	nep5 := nep5.NewNep5Helper(scriptHash, "http://seed1.ngd.network:10332")
	return nep5
}

func CreateNeoAccount() {
	testWallet := wallet.NewWallet()
	err := testWallet.AddNewAccount()
	if err != nil {
		panic(err)
	}
	err = testWallet.Save("neo.wallet")
	if err != nil {
		panic(err)
	}
}
