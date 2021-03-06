package ethereum

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/blockchainpro/usage/ethereum/contractabi/eccm_abi"
	"github.com/blockchainpro/usage/ethereum/contractabi/lock_proxy_abi"
	"github.com/blockchainpro/usage/utiles/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"strings"
	"testing"
)

func TestECCMNotify(t *testing.T) {
	client := DefaultEthereumClient()
	contractAddr := "726532586c50ec9f4080b71f906a3d9779bbd64f"
	height := uint64(8711998)
	lockAddress := ethcommon.HexToAddress(contractAddr)
	lockContract, err := eccm_abi.NewEthCrossChainManager(lockAddress, client.Client)
	if err != nil {
		fmt.Printf("GetSmartContractEventByBlock, error: %s", err.Error())
		return
	}
	opt := &bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: context.Background(),
	}

	// get ethereum lock events from given block
	lockEvents, err := lockContract.FilterCrossChainEvent(opt, nil)
	if err != nil {
		fmt.Printf("GetSmartContractEventByBlock, filter lock events :%s", err.Error())
		return
	}

	for lockEvents.Next() {
		evt := lockEvents.Event
		fmt.Printf("cross id: %s, tx hash: %s, user: %s, tchain: %d, contract: %s, height: %d\n", 
			hex.EncodeToString(evt.TxId), evt.Raw.TxHash.String()[2:], strings.ToLower(evt.Sender.String()[2:]),uint32(evt.ToChainId),
			strings.ToLower(evt.ProxyOrAssetContract.String()[2:]), height)
	}

	// ethereum unlock events from given block
	unlockEvents, err := lockContract.FilterVerifyHeaderAndExecuteTxEvent(opt)
	if err != nil {
		fmt.Printf("GetSmartContractEventByBlock, filter unlock events :%s", err.Error())
		return
	}

	for unlockEvents.Next() {
		evt := unlockEvents.Event
		fmt.Printf("unlock event: tx hash: %s, rtx hash: %s, contract: %s, fchain id: %d, height: %d\n", evt.Raw.TxHash.String()[2:], 
			common.HexStringReverse(hex.EncodeToString(evt.CrossChainTxHash)), hex.EncodeToString(evt.ToContract), uint32(evt.FromChainID), height)
	}
	fmt.Printf("successful\n")
}

func TestLockNotify(t *testing.T) {
	client := DefaultEthereumClient()
	contractAddr := "250e76987d838a75310c34bf422ea9f1ac4cc906"
	height := uint64(10927126)
	proxyAddress := ethcommon.HexToAddress(contractAddr)
	lockContract, err := lock_proxy_abi.NewLockProxy(proxyAddress, client.Client)
	if err != nil {
		fmt.Printf("GetSmartContractEventByBlock, error: %s", err.Error())
		return
	}
	opt := &bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: context.Background(),
	}

	// get ethereum lock events from given block
	lockEvents, err := lockContract.FilterLockEvent(opt)
	if err != nil {
		fmt.Printf("GetSmartContractEventByBlock, filter lock events :%s", err.Error())
		return
	}

	for lockEvents.Next() {
		evt := lockEvents.Event
		fmt.Printf("lock event: txhash: %s, from address: %s, from asset hash: %s, to chaind id: %d, to asset hash: %s, to address: %s, amount: %s\n",
			evt.Raw.TxHash.String()[2:], evt.FromAddress.String()[2:], evt.FromAssetHash.String()[2:], uint32(evt.ToChainId),
			hex.EncodeToString(evt.ToAssetHash), hex.EncodeToString(evt.ToAddress), evt.Amount.String())
		// amountNew := new(big.Int).SetBytes(common.RevertBytes(evt.Amount.Bytes()))
		//amountNew, _ := strconv.ParseUint(string(evt.Amount.Bytes()), 16, 64)
	}

	// ethereum unlock events from given block
	unlockEvents, err := lockContract.FilterUnlockEvent(opt)
	if err != nil {
		fmt.Printf("GetSmartContractEventByBlock, filter unlock events :%s", err.Error())
		return
	}

	for unlockEvents.Next() {
		evt := unlockEvents.Event
		fmt.Printf("unlock event: txhash: %s, to asset hash: %s, to address: %s, amount: %s\n",
			evt.Raw.TxHash.String()[2:], evt.ToAssetHash.String()[2:], evt.ToAddress.String()[2:], evt.Amount.String())
	}
	fmt.Printf("successful\n")
}
