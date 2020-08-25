package neo

import "github.com/joeqian10/neo-gogogo/rpc"

func NewNeoClient() *rpc.RpcClient {
	rawClient := rpc.NewClient("http://seed10.ngd.network:11332")
	return rawClient
}
