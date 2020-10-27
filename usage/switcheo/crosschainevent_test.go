package switcheo

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"testing"
)

func TestCrossChainEvent_from_cosmos(t *testing.T) {
	client := NewSwitcheoClient()
	status, err := client.Status()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current height: %d\n", status.SyncInfo.LatestBlockHeight)

	query := fmt.Sprintf("tx.height=%d AND make_from_cosmos_proof.status='1'", 2880742)
	//query := fmt.Sprintf("tx.height=%d AND tm.event='make_from_cosmos_proof'", 1761)
	res, err := client.TxSearch(query, false, 1, 100, "asc")
	if err != nil {
		panic(err)
	}
	if res.TotalCount == 0 {
		fmt.Printf("There is no event!\n")
		return
	}

	for _, tx := range res.Txs {
		for _, e := range tx.TxResult.Events {
			fmt.Printf("event type: %s\n", e.Type)
			if e.Type == "make_from_cosmos_proof" {
				fmt.Printf("event: %s, status: %s, txhash: %s, txid: %s, key: %s, fromaddress: %s, from contract: %s, to chainid: %s, txparam: %s\n",
					e.Type, string(e.Attributes[0].Value), tx.Hash.String(), string(e.Attributes[1].Value), string(e.Attributes[2].Value),
					string(e.Attributes[3].Value), string(e.Attributes[4].Value), string(e.Attributes[5].Value),
					string(e.Attributes[6].Value))
			} else if e.Type == "lock" {
				fmt.Printf("event: %s, txhash: %s, from asset hash: %s, from address: %s, to chainid: %s, to assset hash: %s, to address: %s, amount: %s\n",
					e.Type, tx.Hash.String(), string(e.Attributes[0].Value), string(e.Attributes[3].Value), string(e.Attributes[1].Value),
					string(e.Attributes[2].Value), string(e.Attributes[4].Value), string(e.Attributes[5].Value))
			}
		}
	}
}

func TestCrossChainEvent_to_cosmos(t *testing.T) {
	client := NewSwitcheoClient()
	status, err := client.Status()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current height: %d\n", status.SyncInfo.LatestBlockHeight)

	//query := fmt.Sprintf("tx.height=%d AND verify_to_cosmos_proof.status='1'", 226293)
	query := fmt.Sprintf("tx.height=%d", 226293)
	res, err := client.TxSearch(query, false, 1, 100, "asc")
	if err != nil {
		panic(err)
	}
	if res.TotalCount == 0 {
		fmt.Printf("There is no event!\n")
		return
	}
	for _, tx := range res.Txs {
		for _, e := range tx.TxResult.Events {
			if e.Type == "make_from_cosmos_proof" {
				fmt.Printf("event: %s, status: %s, txhash: %s, txid: %s, key: %s, fromaddress: %s, from contract: %s, to chainid: %s, txparam: %s\n",
					e.Type, string(e.Attributes[0].Value), tx.Hash.String(), string(e.Attributes[1].Value), string(e.Attributes[2].Value),
					string(e.Attributes[3].Value), string(e.Attributes[4].Value), string(e.Attributes[5].Value),
					string(e.Attributes[6].Value))
			} else if e.Type == "lock" {
				fmt.Printf("event: %s, txhash: %s, from asset hash: %s, from address: %s, to chainid: %s, to assset hash: %s, to address: %s, amount: %s\n",
					e.Type, tx.Hash.String(), string(e.Attributes[0].Value), string(e.Attributes[3].Value), string(e.Attributes[1].Value),
					string(e.Attributes[2].Value), string(e.Attributes[4].Value), string(e.Attributes[5].Value))
			}
		}
	}
}

func TestCrossChainEventSearch(t *testing.T) {
	client := NewSwitcheoClient()

	status, err := client.Status()
	if err != nil {
		panic(err)
	}
	fmt.Printf("current height: %d\n", status.SyncInfo.LatestBlockHeight)

	for i := 1500;i < 1800;i ++ {
		//query := fmt.Sprintf("tx.height=%d AND make_from_cosmos_proof.status='1'", i)
		query := fmt.Sprintf("tx.height=%d", i)
		res, err := client.TxSearch(query, false, 1, 100, "asc")
		if err != nil {
			panic(err)
		}
		if res.TotalCount == 0 {
			//fmt.Printf("There is no event!\n")
			continue
		}

		for _, tx := range res.Txs {
			for _, e := range tx.TxResult.Events {
					fmt.Printf("height: %d event: %s\n", i, e.Type)
			}
		}
	}
}

func TestCrossChainEventSearch1(t *testing.T) {
	client := NewSwitcheoClient()
	query := fmt.Sprintf("tx.height=%d", 2880742)
	res, err := client.TxSearch(query, false, 1, 100, "asc")
	if err != nil {
		panic(err)
	}
	if res.TotalCount == 0 {
		fmt.Printf("There is no event!\n")
	}
	for _, tx := range res.Txs {
		fmt.Printf("tx hash: %s\n", tx.Hash.String())
		fmt.Printf("tx gas usded: %d\n", tx.TxResult.GasUsed)
		for _, e := range tx.TxResult.Events {
			fmt.Printf("event: %s\n", e.Type)
		}
		//
		decoder := auth.DefaultTxDecoder(NewCDC())
		stdTx, err := decoder(tx.Tx)
		if err != nil {
			panic(err)
		}
		/*
		aa := unsafe.Pointer(&stdTx)
		bb := (*authtypes.StdTx)(aa)
		*/
		aa := stdTx
		bb, ok := aa.(authtypes.StdTx)
		if !ok {
			fmt.Printf("xxx")
		}

		//xx := authtypes.StdTx(stdTx)
		fmt.Printf("xxxx: %v", bb)
	}
}
