package solana

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

var (
	SerumV11  = solana.MustPublicKeyFromBase58("4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn")
	OrcaV2    = solana.MustPublicKeyFromBase58("9W959DqEETiGZocYWCQPaJ6sBmUzgfxXfqGeTEdp3aQP")
	SerumV22  = solana.MustPublicKeyFromBase58("9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin")
	SerumV12  = solana.MustPublicKeyFromBase58("BJ3jrUzddfuSrZHXSCxMUUQsjKEyLmuuyZebkcaFp2fg")
	OrcaV1    = solana.MustPublicKeyFromBase58("DjVE6JNiYqPL2QXyCUUh8rNjHrbz9hXHNYt99MQ59qw1")
	SerumV21  = solana.MustPublicKeyFromBase58("EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o")
	Saber     = solana.MustPublicKeyFromBase58("SSwpkEEcbUqx4vtoEByFjSkhKdCT862DNVb52nZg1UZ")
	TokenSwap = solana.MustPublicKeyFromBase58("SwaPpA9LAaLfeLi3a68M4DjnLqgtticKg6CnyNwgAC8")
	System    = solana.MustPublicKeyFromBase58("11111111111111111111111111111111")
	SysClock  = solana.MustPublicKeyFromBase58("SysvarC1ock11111111111111111111111111111111")
	SysRent   = solana.MustPublicKeyFromBase58("SysvarRent111111111111111111111111111111111")
)

//1.45012972  1.44928972
//0.00114144  0.00114144
//2.29922904  2.29922904
var (
	//URL = "http://127.0.0.1:8899"
	URL = rpc.MainNetBetaSerum_RPC
	//URL = "https://autumn-empty-dawn.solana-mainnet.quiknode.pro/924b1527134b73309d1fd8b934a2f078ce31b189"
	//URL = "https://api.devnet.solana.com"
	//URL = "http://23.88.80.164:8899"
	WS = "http://127.0.0.1:8900"
)

var (
	//Player = solana.MustPublicKeyFromBase58("8ZqMeFqC2tAGPpywsyA9apwnKMJcZqFpgBbh6tN88jNq")
	//Key = "4kB7mbH6gRAQp55HwniRVsMn1WoZHZKxWD3eme5YU8GB3TcnwEEMTVaSG4txeDP8CpurjjUSugWBnJBcjmo4WyQu"
	//Player = solana.MustPublicKeyFromBase58("sjcCd2jMG9BfmwwTDiitmJp8vSTJEF4M2wuBm3hR1M9")
	Key = ""
	// mainnet
	Player = solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS")
	//Key = ""
)

var (
	Keys = map[solana.PublicKey]solana.PrivateKey {
		solana.MustPublicKeyFromBase58("3pfNpRNu31FBzx84TnefG6iBkSqQxGtuL5G5v9aaxyv8"):solana.MustPrivateKeyFromBase58("58WF8aF3FDDYLQoPjHiERk53gd2UgLtiJk1rr8saVsDHWz1ieJtEteNYhsDrYGPEYcbZ3SdqvG7TZbvwChaXMS5e"),
		solana.MustPublicKeyFromBase58("6JAPL6Xg3Ea9VViSKTSsyzxxFADA7Gq4E3ypaVHduDFB"):solana.MustPrivateKeyFromBase58("xciQtjmYxpUP39ZQPVm69H7C3gTzZyDC9HScmSNjeMNW6FxZ8ekYkLktVQSkoMCSQxapdJww96G9BYpzGnD6381"),
	}
)

type Instruction struct {
	IsAccounts  []*solana.AccountMeta
	IsData      []byte
	IsProgramID solana.PublicKey
}

func (i *Instruction) Accounts() []*solana.AccountMeta {
	return i.IsAccounts
}

func (i *Instruction) ProgramID() solana.PublicKey {
	return i.IsProgramID
}

func (i *Instruction) Data() ([]byte, error) {
	return i.IsData, nil
}

func SendTransaction(ins []solana.Instruction) {
	ctx := context.Background()
	client := rpc.New(URL)
	getRecentBlockHashResult, err := client.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}
	blockHash := getRecentBlockHashResult.Value.Blockhash

	builder := solana.NewTransactionBuilder()
	for _, in := range ins {
		builder.AddInstruction(in)
	}

	builder.SetRecentBlockHash(blockHash)
	builder.SetFeePayer(Player)
	trx, err := builder.Build()
	if err != nil {
		panic(err)
	}
	trx.Sign(getKey)

	trxJson, _ := json.MarshalIndent(trx, "", "    ")
	fmt.Printf("tx: %s\n", trxJson)
	/*
		txData, err := trx.MarshalBinary()
		if err != nil {
			return nil, err
		}
		base64Data := base64.StdEncoding.EncodeToString(txData)
		fmt.Printf("tx data: %s\n", base64Data)
	*/

	if true {
		response, err := client.SimulateTransactionWithOpts(ctx, trx, &rpc.SimulateTransactionOpts{
			SigVerify:              true,
			Commitment:             rpc.CommitmentFinalized,
		})
		if err != nil {
			panic(err)
		}
		simulateTransactionResponse := response.Value
		if simulateTransactionResponse.Logs == nil {
			panic(fmt.Errorf("log is nil, simulate failed before the transaction was able to executed, such as signature verification failure or invalid blockhash"))
		}
		logsJson, _ := json.MarshalIndent(simulateTransactionResponse.Logs, "", "    ")
		fmt.Printf("logs: %s\n", logsJson)
		if simulateTransactionResponse.Err != nil {
			panic(simulateTransactionResponse.Err)
		}
	}

	if false {
		response, err := client.SendTransactionWithOpts(ctx, trx, false, rpc.CommitmentFinalized)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		fmt.Printf("hash: %s\n",response.String())
	}
}



func getKey(key solana.PublicKey) *solana.PrivateKey {
	_, ok := Keys[key]
	fmt.Printf("use key: %s\n", key.String())
	if !ok {
		key = solana.MustPublicKeyFromBase58("3pfNpRNu31FBzx84TnefG6iBkSqQxGtuL5G5v9aaxyv8")
	}
	xx, ok := Keys[key]
	return &xx
}

