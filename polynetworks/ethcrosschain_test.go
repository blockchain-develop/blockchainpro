package polynetworks

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/ontio/multi-chain/common"
	crosschain_common "github.com/ontio/multi-chain/native/service/header_sync/common"
	crosschain_utils "github.com/ontio/multi-chain/native/service/utils"
	"testing"
)

func TestStoreQuery_ETH_CurrentHeaderHeight(t *testing.T) {
	sdk := newMultiChanSdk()
	//
	var sideChainId uint64 = 2
	contractAddress := crosschain_utils.HeaderSyncContractAddress
	key := append([]byte(crosschain_common.CURRENT_HEADER_HEIGHT), crosschain_utils.GetUint64Bytes(sideChainId)...,)
	// try to get storage
	result, err := sdk.GetStorage(contractAddress.ToHexString(), key)
	if err != nil {
		fmt.Printf("Get Storage err: %s\n", err.Error())
	}
	if result == nil || len(result) == 0 {
		fmt.Printf("There is not this store\n")
	} else {
		height := binary.LittleEndian.Uint64(result)
		fmt.Printf("Current header height: %d\n", height)
	}
}

func TestStoreQuery_ETH_MainChainHash(t *testing.T) {
	sdk := newMultiChanSdk()
	//
	var sideChainId uint64 = 2
	var height uint64 = 7884063
	contractAddress := crosschain_utils.HeaderSyncContractAddress
	key := append([]byte(crosschain_common.MAIN_CHAIN), crosschain_utils.GetUint64Bytes(sideChainId)...,)
	key = append(key, crosschain_utils.GetUint64Bytes(height)...)
	// try to get storage
	result, err := sdk.GetStorage(contractAddress.ToHexString(), key)
	if err != nil {
		fmt.Printf("Get Storage err: %s\n", err.Error())
	}
	if result == nil || len(result) == 0 {
		fmt.Printf("There is not this store\n")
	} else {
		result = revertBytes(result)
		hash, err := common.Uint256ParseFromBytes(result)
		if err != nil {
			fmt.Printf("The store is not hash, err: %s\n", err.Error())
		} else {
			fmt.Printf("The hash is: %s\n", hash.ToHexString())
		}
	}
}

func TestStoreQuery_ETH_GenesisHeader(t *testing.T) {
	sdk := newMultiChanSdk()
	//
	var sideChainId uint64 = 2
	contractAddress := crosschain_utils.HeaderSyncContractAddress
	key := append([]byte(crosschain_common.GENESIS_HEADER), crosschain_utils.GetUint64Bytes(sideChainId)...)
	// try to get storage
	result, err := sdk.GetStorage(contractAddress.ToHexString(), key)
	if err != nil {
		fmt.Printf("Get Storage err: %s\n", err.Error())
	}
	if result == nil || len(result) == 0 {
		fmt.Printf("There is not this store\n")
	} else {
		var headerWithDifficultySum HeaderWithDifficultySum
		err := json.Unmarshal(result, &headerWithDifficultySum)
		if err != nil {
			fmt.Printf("Unmarshal store err: %s\n", err)
		} else {
			fmt.Printf("The hash is: %s\n", string(result))
		}
	}
}
