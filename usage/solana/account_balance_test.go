package solana

import (
	"context"
	"encoding/hex"
	json2 "encoding/json"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"os"
	"testing"
	"time"
)

func TestAccountBalance(t *testing.T) {
	client := rpc.New("https://autumn-empty-dawn.solana-mainnet.quiknode.pro/924b1527134b73309d1fd8b934a2f078ce31b189/")
	ctx := context.Background()
	addr := solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS")
	result, err := client.GetBalance(ctx, addr, rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", result.Value)
}

type NodeInfo struct {
	Name string `json:"name"`
	Key string `json:"key"`
	Rpc string `json:"rpc"`
	Usable bool `json:"usable"`
}

func tryNode(node *rpc.GetClusterNodesResult) *NodeInfo {
	nodeInfo := &NodeInfo{}
	nodeInfo.Key = node.Pubkey.String()
	if node.RPC == nil {
		return nodeInfo
	}
	// test
	{
		nodeInfo.Rpc = "http://" + *node.RPC
		client := rpc.New(nodeInfo.Rpc)
		addr := solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_, err := client.GetBalance(ctx, addr, rpc.CommitmentFinalized)
		if err != nil {
			return nodeInfo
		}
		nodeInfo.Usable = true
		return nodeInfo
	}
	/*
	{
		xx := *node.RPC
		xx = xx[:len(xx) - 4]
		xx = xx + "8900"
		wsx := "ws://" + xx
		fmt.Printf("ws : %s\n", wsx)
		client, err := ws.Connect(context.Background(), wsx)
		if err != nil {
			fmt.Printf("not unusable")
			return nodeInfo
		}
		su, err := client.SlotSubscribe()
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
			return nodeInfo
		}
		re, err := su.Recv()
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
			return nodeInfo
		}
		fmt.Printf("=================ok %v\n", re)
		su.Unsubscribe()
		return nodeInfo
	}

	 */
}

func TestValidators(t *testing.T) {
	client := rpc.New(URL)
	ctx := context.Background()
	response, err := client.GetClusterNodes(ctx)
	if err != nil {
		panic(err)
	}
	nodeInfos := make([]*NodeInfo, 0)
	for _, node := range response {
		nodeInfos = append(nodeInfos, tryNode(node))
	}
	json, _ := json2.MarshalIndent(nodeInfos, "", "  ")
	fmt.Printf(string(json))
	err = os.WriteFile("validator.json", json, 0644)
	if err != nil {
		panic(err)
	}
}

func TestGetBlock1(t *testing.T) {
	client := rpc.New(URL)
	ctx := context.Background()
	reward := false
	opts := &rpc.GetConfirmedBlockOpts{
		Encoding:           "jsonParsed",
		Rewards: &reward,
	}
	block, err := client.GetConfirmedBlockWithOpts(ctx, 128039160, opts)
	if err != nil {
		panic(err)
	}
	for _, tx := range block.Transactions {
		if tx.Transaction.Signatures[0] == solana.MustSignatureFromBase58("7aBM5xHMZ8ncpvp94mTsFGoqZomZTCmaJ6A11LUte1ejRe6iKKkd3xH4giQeFENhfEGpDkaCSaGY5LHHxDSMqdv") {
			fmt.Printf("%v", tx)
		}
	}
}

func TestPubkey(t *testing.T) {
	aa := solana.MustPublicKeyFromBase58("3pfNpRNu31FBzx84TnefG6iBkSqQxGtuL5G5v9aaxyv8")
	bb := aa.Bytes()
	fmt.Printf("bb: %s", hex.EncodeToString(bb))
}




