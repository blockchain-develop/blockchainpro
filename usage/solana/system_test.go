package solana

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"testing"
)

func TestSystem_AccountAssign(t *testing.T) {
	ins := make([]solana.Instruction, 0)
	{
		myaccount := solana.MustPublicKeyFromBase58("6JAPL6Xg3Ea9VViSKTSsyzxxFADA7Gq4E3ypaVHduDFB")
		ownerId := solana.MustPublicKeyFromBase58("Atv4JkoVwY2hVEpH1wHhpZx4Q534ZLNm64VxVRPRJQzq")
		//
		data := make([]byte, 36)
		binary.LittleEndian.PutUint32(data[0:], 1)
		copy(data[4:], ownerId.Bytes())

		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{
				{PublicKey: myaccount, IsSigner: true, IsWritable: true},
			},
			IsData:      data,
			IsProgramID: solana.MustPublicKeyFromBase58("11111111111111111111111111111111"),
		}
		ins = append(ins, instruction)
	}

	if true {
		myaccount := solana.MustPublicKeyFromBase58("6JAPL6Xg3Ea9VViSKTSsyzxxFADA7Gq4E3ypaVHduDFB")
		ownerId := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
		//
		data := make([]byte, 36)
		binary.LittleEndian.PutUint32(data[0:], 1)
		copy(data[4:], ownerId.Bytes())

		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{
				{PublicKey: myaccount, IsSigner: true, IsWritable: true},
			},
			IsData:      data,
			IsProgramID: solana.MustPublicKeyFromBase58("11111111111111111111111111111111"),
		}
		ins = append(ins, instruction)
	}
	//
	ctx := context.Background()
	rpcClient := rpc.New(rpc.MainNetBetaSerum_RPC)
	blockHashRes, err := rpcClient.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}
	//
	builder := solana.NewTransactionBuilder()
	for _, i := range ins {
		builder.AddInstruction(i)
	}
	builder.SetRecentBlockHash(blockHashRes.Value.Blockhash)
	builder.SetFeePayer(Player)
	trx, err := builder.Build()
	if err != nil {
		panic(err)
	}
	trx.Sign(getKey)
	//
	response, err := rpcClient.SimulateTransactionWithOpts(ctx, trx, &rpc.SimulateTransactionOpts{
		SigVerify:              true,
		Commitment:             rpc.CommitmentFinalized,
		ReplaceRecentBlockhash: false,
	})
	if err != nil {
		fmt.Printf("SimulateTransactionWithOpts err: %s\n", err.Error())
		return
	}
	simulateTransactionResponse := response.Value
	if simulateTransactionResponse.Logs == nil {
		fmt.Printf("log is nil, simulate failed before the transaction was able to executed, such as signature verification failure or invalid blockhash\n")
		return
	}
	logsJson, _ := json.MarshalIndent(simulateTransactionResponse.Logs, "", "    ")
	fmt.Printf("logs: %s\n", string(logsJson))
	if simulateTransactionResponse.Err != nil {
		fmt.Printf("SimulateTransactionWithOpts err: %s\n", simulateTransactionResponse.Err)
		return
	}
	fmt.Printf("SimulateTransactionWithOpts successful!\n")
	return
}
