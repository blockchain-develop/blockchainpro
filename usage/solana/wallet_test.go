package solana

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"os"
	"testing"
)

func TestNewAccount(t *testing.T) {
	wallet := solana.NewWallet()
	account := wallet.PublicKey()
	private := wallet.PrivateKey
	fmt.Printf("account: %s, private: %s\n", account.String(), private.String())
}


func TestNewAccounts(t *testing.T) {
	for i := 0;i < 10;i ++ {
		wallet := solana.NewWallet()
		account := wallet.PublicKey()
		fmt.Printf("account: %s\n", account)
	}
}

func TestAirdrop(t *testing.T) {
	client := rpc.New(URL)
	client.RequestAirdrop(context.Background(), Player, 1000000000, rpc.CommitmentFinalized)
}


func TestGetRecentBlockHash(t *testing.T) {
	client := rpc.New("https://connect.runnode.com/?apikey=suthSNHk0SiKJyBlDR0q")
	ctx := context.Background()
	//
	getRecentBlockHashResult, err := client.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}
	blockHash := getRecentBlockHashResult.Value.Blockhash
	fmt.Printf("block hash: %s\n", blockHash.String())
}

func CreateArbitrageAccount() solana.PublicKey {
	client := rpc.New(URL)
	ctx := context.Background()
	//
	getRecentBlockHashResult, err := client.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}
	blockHash := getRecentBlockHashResult.Value.Blockhash
	//
	space := uint64(72)
	lamports, err := client.GetMinimumBalanceForRentExemption(ctx, space, rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}
	//
	keys := make(map[solana.PublicKey]solana.PrivateKey)
	keys[Player] = solana.MustPrivateKeyFromBase58(Key)
	//
	wallet := solana.NewWallet()
	splTokenAccount := wallet.PublicKey()
	fmt.Printf("spl token account: %s\n", splTokenAccount.String())
	fmt.Printf("spl token account pri: %s\n", wallet.PrivateKey.String())
	keys[splTokenAccount] = wallet.PrivateKey
	{
		file := fmt.Sprintf("%s", splTokenAccount.String())
		err = os.WriteFile(file, []byte(wallet.PrivateKey.String()), 0644)
		if err != nil {
			panic(err)
		}
	}
	//
	//
	builder := solana.NewTransactionBuilder()
	{
		data := make([]byte, 52)
		binary.LittleEndian.PutUint32(data[0:], 0)
		binary.LittleEndian.PutUint64(data[4:], lamports)
		binary.LittleEndian.PutUint64(data[12:], space)
		copy(data[20:], ARB.Bytes())
		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{
				{PublicKey: Player, IsSigner: true, IsWritable: true},
				{PublicKey: splTokenAccount, IsSigner: true, IsWritable: true},
			},
			IsData:      data,
			IsProgramID: System,
		}
		builder.AddInstruction(instruction)
	}

	builder.SetRecentBlockHash(blockHash)
	builder.SetFeePayer(Player)
	trx, err := builder.Build()
	if err != nil {
		panic(err)
	}
	trx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		p, ok := keys[key]
		if ok {
			return &p
		} else {
			return nil
		}
	})

	if false {
		res, err := client.SimulateTransactionWithOpts(ctx, trx, &rpc.SimulateTransactionOpts{
			SigVerify:              false,
			Commitment:             rpc.CommitmentFinalized,
			ReplaceRecentBlockhash: true,
		})
		trxJson, _ := json.MarshalIndent(trx, "", "    ")
		fmt.Printf("tx: %s\n", trxJson)

		resJson, _ := json.MarshalIndent(res, "", "    ")
		fmt.Printf("tx: %s\n", resJson)

		if err != nil {
			panic(err)
		}
	}
	if true {
		trxJson, _ := json.MarshalIndent(trx, "", "    ")
		fmt.Printf("tx: %s\n", trxJson)

		hash, err := client.SendTransactionWithOpts(ctx, trx, false, rpc.CommitmentFinalized)
		fmt.Printf("tx hash: %s\n", hash.String())
		fmt.Printf("err: %v\n", err)
		if err != nil {
			panic(err)
		}
	}
	return splTokenAccount
}

func TestCreateArbitrageAccount(t *testing.T) {
	CreateArbitrageAccount()
}


func TestProgramDeriveAddress(t *testing.T) {
	program := solana.MustPublicKeyFromBase58("6UeJYTLU1adaoHWeApWsoj1xNEDbWA2RhM2DLc8CrDDi")
	//sol_usdc := solana.MustPublicKeyFromBase58("APDFRM3HMr8CAGXwKHiu2f5ePSpaiEJhaURwhsRrUUt9")
	pricePda, _, err := solana.FindProgramAddress([][]byte{[]byte("PRICE")}, program)
	if err != nil {
		panic(err)
	}
	fmt.Printf("price pda: %s\n", pricePda.String())

	priceKey, err := solana.CreateWithSeed(pricePda, "POOL__ap", program)
	if err != nil {
		panic(err)
	}
	fmt.Printf("price key: %s\n", priceKey.String())
}

func TestProgramDeriveAddress1(t *testing.T) {
	program := solana.MustPublicKeyFromBase58("6UeJYTLU1adaoHWeApWsoj1xNEDbWA2RhM2DLc8CrDDi")
	//sol_usdc := solana.MustPublicKeyFromBase58("APDFRM3HMr8CAGXwKHiu2f5ePSpaiEJhaURwhsRrUUt9")
	basePda, _, err := solana.FindProgramAddress([][]byte{[]byte("2")}, program)
	if err != nil {
		panic(err)
	}
	fmt.Printf("base pda: %s\n", basePda.String())

	priceKey, err := solana.CreateWithSeed(basePda, "PriceSummaries", program)
	if err != nil {
		panic(err)
	}
	fmt.Printf("price key: %s\n", priceKey.String())
}

func TestGetBlock(t *testing.T) {
	client := rpc.New(URL)
	ctx := context.Background()
	client.GetBlock(ctx, 1)
}
