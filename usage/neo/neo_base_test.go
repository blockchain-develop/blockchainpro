package neo

import (
	"encoding/hex"
	"fmt"
	"github.com/joeqian10/neo-gogogo/helper"
	"testing"
)


func TestNeoAddress(t *testing.T) {
	//hash_str := utils.HexStringReverse("bf9c0fd26055ff19245c8080df06d97ae32db3d7")
	hash_str := "6e43f9988f2771f1a2b140cb3faad424767d39fc"
	//hash_str = utils.HexStringReverse(hash_str)
	hash, err := hex.DecodeString(hash_str)
	if err != nil {
		panic(err)
	}
	addr, err := helper.UInt160FromBytes(hash)
	if err != nil {
		panic(err)
	}
	addr_str := helper.ScriptHashToAddress(addr)
	fmt.Printf("addr, %s, %s\n", hash_str, addr_str)
}

func TestCreateNeoAccount(t *testing.T) {
	CreateNeoAccount()
}