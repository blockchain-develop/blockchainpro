package ethereum

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/blockchainpro/usage/ethereum/contractabi/eccm_abi"
	"github.com/blockchainpro/usage/ethereum/contractabi/lock_proxy_abi"
	"github.com/blockchainpro/usage/ethereum/contractabi/usdt_abi"
	"github.com/blockchainpro/usage/utiles/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
	"testing"
)

func TestECCMNotify(t *testing.T) {
	client := DefaultEthereumClient()
	contractAddr := "838bf9e95cb12dd76a54c9f9d2e3082eaf928270"
	height := uint64(12431286)
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
		fmt.Printf("raw: %s\n", hex.EncodeToString(evt.Rawdata))
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
		fmt.Printf("raw: %s\n", hex.EncodeToString(evt.Raw.Data))
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

func TestUsdtNotify(t *testing.T) {
	client := DefaultEthereumClient()
	contractAddr := "0xdac17f958d2ee523a2206206994597c13d831ec7"
	height := uint64(10097862)
	usdtAddress := ethcommon.HexToAddress(contractAddr)
	usdtContract, err := usdt_abi.NewERC20(usdtAddress, client.Client)
	if err != nil {
		panic(err)
	}
	opt := &bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: context.Background(),
	}

	// get ethereum lock events from given block
	transferEvents, err := usdtContract.FilterTransfer(opt, nil, nil)
	if err != nil {
		panic(err)
	}

	for transferEvents.Next() {
		evt := transferEvents.Event
		fmt.Printf("evt: \n    from: %s\n    to: %s\n    value: %s\n", evt.From.String(), evt.To.String(), evt.Value.String())
	}
}

func TestUsdtNotify1(t *testing.T) {
	client := DefaultEthereumClient()
	ctx := context.Background()
	hash := ethcommon.HexToHash("0xeef1725fd660767404e94dad7f5280a7eeca838210457d8026e9b73645de338f")
	receipt, err := client.GetTransactionReceipt(ctx, hash)
	if err != nil {
		panic(err)
	}

	logs := receipt.Logs
	for _, log := range logs {
		txHash := log.TxHash.String()
		blockHash := log.BlockHash.String()
		from := ethcommon.BytesToAddress(log.Topics[1].Bytes())
		to := ethcommon.BytesToAddress(log.Topics[2].Bytes())
		value := new(big.Int).SetBytes(log.Data)

		fmt.Printf("evt: \n    block hash: %s\n    tx hash: %s\n    from: %s\n    to: %s\n    value: %s\n", blockHash, txHash, from.String(), to.String(), value.String())
	}
}