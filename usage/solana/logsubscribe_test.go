package solana

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"testing"
)

func TestLogSubscribe(t *testing.T) {
	url := "wss://199.127.60.141:443"
	client, err := ws.Connect(context.Background(), url)
	if err != nil {
		panic(err)
	}
	program := solana.MustPublicKeyFromBase58("SwaPpA9LAaLfeLi3a68M4DjnLqgtticKg6CnyNwgAC8")
	{
		sub, err := client.LogsSubscribeMentions(
			program,
			rpc.CommitmentRecent)
		if err != nil {
			panic(err)
		}
		defer sub.Unsubscribe()
		for {
			got, err := sub.Recv()
			if err != nil {
				panic(err)
			}
			spew.Dump(got)
		}
	}
}


