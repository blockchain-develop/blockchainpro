package polynetworks

import (
	"fmt"
	"github.com/ontio/multi-chain/common"
	"github.com/ontio/multi-chain/native/service/header_sync/btc"
	crosschain_common "github.com/ontio/multi-chain/native/service/header_sync/common"
	crosschain_utils "github.com/ontio/multi-chain/native/service/utils"
	"testing"
)

func TestStoreQuery_BTC_GenesisHeader(t *testing.T) {
	sdk := newMultiChanSdk()
	//
	var sideChainId uint64 = 1
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
		sh := new(btc.StoredHeader)
		err := sh.Deserialization(common.NewZeroCopySource(result))
		if err != nil {
			fmt.Printf("Deserialization store err: %s\n", err)
		} else {
			fmt.Printf("The result is: %v\n", sh)
		}
	}
}
