package neo

import (
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/nep5"
	"github.com/joeqian10/neo-gogogo/rpc"
)

func NewNeoClient() *rpc.RpcClient {
	rawClient := rpc.NewClient("http://seed10.ngd.network:11332")
	return rawClient
}

func NewNep5(hash string) *nep5.Nep5Helper {
	scriptHash, err := helper.UInt160FromString(hash)
	if err != nil {
		panic(err)
	}
	nep5 := nep5.NewNep5Helper(scriptHash, "http://seed1.ngd.network:20332")
	return nep5
}
