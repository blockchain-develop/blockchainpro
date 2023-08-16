package solana

import (
	"context"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"testing"
)

func TestSaber_SaberSwap(t *testing.T) {
	saberProgramID := solana.MustPublicKeyFromBase58("SSwpkEEcbUqx4vtoEByFjSkhKdCT862DNVb52nZg1UZ")
	tokenProgramID := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	feePlayer := solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS")

	client := rpc.New(rpc.MainNetBetaSerum_RPC)
	ctx := context.Background()
	getRecentBlockHashResult, err := client.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}
	blockHash := getRecentBlockHashResult.Value.Blockhash

	builder := solana.NewTransactionBuilder()
	// accounts
	market := solana.MustPublicKeyFromBase58("YAkoNb6HKmSxQN9L8hiBE5tPJRsniSSMzND1boHmZxe")
	market_authority, _, _ := solana.FindProgramAddress([][]byte{market.Bytes()}, saberProgramID)
	user_authority := solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS")
	//user_authority := solana.MustPublicKeyFromBase58("Cici3VkBztsxuZucqBKuiT8pfUtbdNTH3sLSVfTyi7rM")
	user_source_account := solana.MustPublicKeyFromBase58("EDASG9z1jHYMxDqda9JeFM9E7UJAEg7j9LqJtqevwKL5")
	//user_source_account := solana.MustPublicKeyFromBase58("7UhPcms9wTAZU8PnekyCkt64U64fTjm6YGd5Tgeyrdij")
	token_source := solana.MustPublicKeyFromBase58("EnTrdMMpdhugeH6Ban6gYZWXughWxKtVGfCwFn78ZmY3")
	token_dest := solana.MustPublicKeyFromBase58("CfWX7o2TswwbxusJ4hCaPobu2jLCb1hfXuXJQjVq3jQF")
	user_dest_account := solana.MustPublicKeyFromBase58("C5dr2fk9zLvjuSjxe3tdvFurcZHKbfU1mHbpdv9Bfr7T")
	//user_dest_account := solana.MustPublicKeyFromBase58("9KaA7vEBUdRCcBWxfuMjxYwKfvu8Us3Cg5gkhVFt2LNk")
	feeAdmin := solana.MustPublicKeyFromBase58("GLztedC76MeBXjAmVXMezcHQzdmQaVLiXCZr9KEBSR6Y")
	sysClock := solana.MustPublicKeyFromBase58("SysvarC1ock11111111111111111111111111111111")
	//
	amountIn := uint64(0)
	{
		getAccountInfoResponse , err := client.GetAccountInfo(ctx, user_source_account)
		if err != nil {
			panic(err)
		}
		xx := decodeAccount(getAccountInfoResponse.Value)
		fmt.Printf("account: %s, amount: %d\n", user_source_account.String(), xx.Amount)
		amountIn = xx.Amount
	}

	{
		getAccountInfoResponse , err := client.GetAccountInfo(ctx, user_dest_account)
		if err != nil {
			panic(err)
		}
		xx := decodeAccount(getAccountInfoResponse.Value)
		fmt.Printf("account: %s, amount: %d\n", user_dest_account.String(), xx.Amount)
	}

	data := make([]byte, 17)
	data[0] = 1
	binary.LittleEndian.PutUint64(data[1:], amountIn)
	binary.LittleEndian.PutUint64(data[9:], 0)

	builder.AddInstruction(&orcaTransactionInstructions{
		accounts:  []*solana.AccountMeta{
			{PublicKey: market, IsSigner: false, IsWritable: true},
			{PublicKey: market_authority, IsSigner: false, IsWritable: false},
			{PublicKey: user_authority, IsSigner: true, IsWritable: false},
			{PublicKey: user_source_account, IsSigner: false, IsWritable: true},
			{PublicKey: token_source, IsSigner: false, IsWritable: true},
			{PublicKey: token_dest, IsSigner: false, IsWritable: true},
			{PublicKey: user_dest_account, IsSigner: false, IsWritable: true},
			{PublicKey: feeAdmin, IsSigner: false, IsWritable: true},
			{PublicKey: tokenProgramID, IsSigner: false, IsWritable: false},
			{PublicKey: sysClock, IsSigner: false, IsWritable: false},
		},
		data:      data,
		programID: saberProgramID,
	})

	builder.SetRecentBlockHash(blockHash)
	builder.SetFeePayer(feePlayer)
	trx, err := builder.Build()
	if err != nil {
		panic(err)
	}
	trx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		if key == solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS") {
			xx := solana.MustPrivateKeyFromBase58("")
			return &xx
		} else {
			return nil
		}
	})

	//
	trxJson, _ := json.MarshalIndent(trx, "", "    ")
	fmt.Printf("transaction: %s\n", trxJson)

	//
	txData, err := trx.MarshalBinary()
	if err != nil {
		panic(err)
	}
	base64Data := base64.StdEncoding.EncodeToString(txData)
	fmt.Printf("tx data: %s\n", base64Data)

	if false {
		simulateTransactionResponse, err := client.SimulateTransactionWithOpts(ctx, trx, &rpc.SimulateTransactionOpts{
			SigVerify:              true,
			Commitment:             rpc.CommitmentFinalized,
			Accounts: &rpc.SimulateTransactionAccountsOpts{
				Encoding:  solana.EncodingBase64,
				Addresses: []solana.PublicKey{user_source_account, user_dest_account},
			},
		})
		//simulateTransactionResponse, err := client.SimulateTransaction(ctx, trx)
		if err != nil {
			panic(err)
		}
		fmt.Printf("simulate logs: %s\n", simulateTransactionResponse.Value.Logs)
		{
			xx := decodeAccount(simulateTransactionResponse.Value.Accounts[0])
			fmt.Printf("account: %s, amount: %d\n", user_source_account.String(), xx.Amount)
		}

		{
			xx := decodeAccount(simulateTransactionResponse.Value.Accounts[1])
			fmt.Printf("account: %s, amount: %d\n", user_dest_account.String(), xx.Amount)
		}
	}
	if true {
		hash, err := client.SendTransaction(ctx, trx)
		//simulateTransactionResponse, err := client.SimulateTransaction(ctx, trx)
		if err != nil {
			panic(err)
		}
		fmt.Printf("hash: %s\n", hash.String())
	}
}
