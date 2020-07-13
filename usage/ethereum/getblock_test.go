package ethereum

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetBlockByHash(t *testing.T) {
	ethClient := NewEthClient("https://ropsten.infura.io/v3/1ba5f3635395470e9a3f19ba7a852144")
	//
	hash := "0x37285a648b55482375bab8cd8fec721fb7f1604db2a5bc5aae2058d4791b4dd2"
	newHeader, err := ethClient.GetNodeHeaderByHash(hash)
	if err != nil {
		return
	}
	newheaderJson, _ := json.Marshal(newHeader)
	fmt.Printf("hash1: %s, hash: %s\n", hash, newHeader.Hash().String())
	fmt.Printf("the header: %s\n",hex.EncodeToString(newheaderJson))
	fmt.Printf("the header: %s\n", string(newheaderJson))
}

func TestGetBlockByHeight(t *testing.T) {
	ethClient := NewEthClient("https://mainnet.infura.io/v3/1ba5f3635395470e9a3f19ba7a852144")
	//ethClient := NewEthClient("https://ropsten.infura.io/v3/1ba5f3635395470e9a3f19ba7a852144")
	//
	for height := uint64(10418300);height < uint64(10418300 + 10);height ++ {
		header, err := ethClient.GetNodeHeader(height)
		if err != nil {
			fmt.Printf("GetNodeHeader on height :%d failed", height)
			return
		}
		headerJson, _ := json.Marshal(header)
		fmt.Printf("====================================================================================\n")
		fmt.Printf("height: %d, hash: %s\n", height, header.Hash().String())
		fmt.Printf("the header: %s\n", hex.EncodeToString(headerJson))
		fmt.Printf("the header: %s\n", string(headerJson))
	}
}
