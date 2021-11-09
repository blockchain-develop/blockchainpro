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

func getKey(key solana.PublicKey) *solana.PrivateKey {
	wallet := solana.NewWallet()
	return &wallet.PrivateKey
}

type orcaTransactionInstructions struct {
	accounts []*solana.AccountMeta
	data []byte
	programID solana.PublicKey
}

func (t *orcaTransactionInstructions) Accounts() []*solana.AccountMeta {
	return t.accounts
}

func (t *orcaTransactionInstructions) ProgramID() solana.PublicKey {
	return t.programID
}

func (t *orcaTransactionInstructions) Data() ([]byte, error) {
	return t.data, nil
}

/*
   let account_info_iter = &mut accounts.iter();
   let swap_info = next_account_info(account_info_iter)?;
   let authority_info = next_account_info(account_info_iter)?;
   let user_transfer_authority_info = next_account_info(account_info_iter)?;
   let source_info = next_account_info(account_info_iter)?;
   let swap_source_info = next_account_info(account_info_iter)?;
   let swap_destination_info = next_account_info(account_info_iter)?;
   let destination_info = next_account_info(account_info_iter)?;
   let pool_mint_info = next_account_info(account_info_iter)?;
   let pool_fee_account_info = next_account_info(account_info_iter)?;
   let token_program_info = next_account_info(account_info_iter)?;
 */
func TestSimulateTransaction(t *testing.T) {
	orcaProgramID := solana.MustPublicKeyFromBase58("DjVE6JNiYqPL2QXyCUUh8rNjHrbz9hXHNYt99MQ59qw1")
	tokenProgramID := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	//
	tokenswap1 := solana.MustPublicKeyFromBase58("2Qqh2DS448qZMwhA8o5jBnSQF54uPPFHUJZULErA1or1")
	tokenswap1_authority, _, _ := solana.FindProgramAddress([][]byte{tokenswap1.Bytes()}, orcaProgramID)
	tokenswap1_user_authority := solana.MustPublicKeyFromBase58("F8Vyqk3unwxkXukZFQeYyGmFfTG3CAX4v24iyrjEYBJV")
	tokenswap1_user_account := solana.MustPublicKeyFromBase58("36c6YqAwyGKQG66XEp2dJc5JqjaBNv7sVghEtJv4c7u6")
	tokenswap1_tokenA := solana.MustPublicKeyFromBase58("Bm2Sj68iAWhmxmQfSLfPzbvDQsd9MAAiCrbDwA7KQ5Va")
	tokenswap1_tokenB := solana.MustPublicKeyFromBase58("DdcQxiv9d8JuQMYmerWDgxsciASqHxkEpymqVSkizZpc")
	tokneswap1_user2_account := solana.MustPublicKeyFromBase58("648kz2PXcnUwDWHqhu1k38gFLbnRMV88J9ByLtcesE2a")
	tokenswap1_pool := solana.MustPublicKeyFromBase58("E4cthfUFaDd4x5t1vbeBNBHm7isqhM8kapthPzPJz1M2")
	tokenswap1_fee_account := solana.MustPublicKeyFromBase58("GS6F9UV9TLcEW74LFCt3Fj5Lpvnfexngvj1VbAF5qUNv")
	data := make([]byte, 16)
	binary.LittleEndian.PutUint64(data, 100000)
	binary.LittleEndian.PutUint64(data[8:], 0)

	//tokenswap2 := solana.MustPublicKeyFromBase58("")
	//tokenswap3 := solana.MustPublicKeyFromBase58("")
	instructions := []solana.Instruction {
		&orcaTransactionInstructions{
			accounts:  []*solana.AccountMeta{
				{PublicKey: tokenswap1, IsSigner: false, IsWritable: false},
				{PublicKey: tokenswap1_authority, IsSigner: false, IsWritable: false},
				{PublicKey: tokenswap1_user_authority, IsSigner: false, IsWritable: false},
				{PublicKey: tokenswap1_user_account, IsSigner: false, IsWritable: true},
				{PublicKey: tokenswap1_tokenA, IsSigner: false, IsWritable: true},
				{PublicKey: tokenswap1_tokenB, IsSigner: false, IsWritable: true},
				{PublicKey: tokneswap1_user2_account, IsSigner: false, IsWritable: true},
				{PublicKey: tokenswap1_pool, IsSigner: false, IsWritable: true},
				{PublicKey: tokenswap1_fee_account, IsSigner: false, IsWritable: true},
				{PublicKey: tokenProgramID, IsSigner: false, IsWritable: false},
				{},
			},
			data:      data,
			programID: orcaProgramID,
		},
		/*
		&orcaTransactionInstructions{
			accounts:  []*solana.AccountMeta{
				{PublicKey: tokenswap2, IsSigner: true, IsWritable: false},
			},
			data:      []byte{},
			programID: solana.MustPublicKeyFromBase58(""),
		},
		&orcaTransactionInstructions{
			accounts:  []*solana.AccountMeta{
				{PublicKey: tokenswap3, IsSigner: true, IsWritable: false},
			},
			data:      []byte{},
			programID: solana.MustPublicKeyFromBase58(""),
		},
		*/
	}
	blockHash, err := solana.HashFromBase58("5Kh7VPdCjJmjsTTBHTdUU5KoMo3Y8vvU5zGzzXac2fg3")
	if err != nil {
		panic(err)
	}
	trx, err := solana.NewTransaction(instructions, blockHash)
	if err != nil {
		panic(err)
	}
	client := rpc.New(rpc.MainNetBetaSerum_RPC)
	ctx := context.Background()
	simulateTransactionResponse, err := client.SimulateTransactionWithOpts(ctx, trx, &rpc.SimulateTransactionOpts{
		//SigVerify:              false,
		//Commitment:             "",
		//ReplaceRecentBlockhash: false,
		Accounts:               &rpc.SimulateTransactionAccountsOpts{
			Encoding: solana.EncodingBase64,
			Addresses: []solana.PublicKey{tokneswap1_user2_account},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("SimulateTransactionResponse: %v\n", simulateTransactionResponse)
}

func TestSimulateTransaction2(t *testing.T) {
	orcaProgramID := solana.MustPublicKeyFromBase58("DjVE6JNiYqPL2QXyCUUh8rNjHrbz9hXHNYt99MQ59qw1")
	tokenProgramID := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	feePlayer := solana.MustPublicKeyFromBase58("Y2akr3bXHRsqyP1QJtbm9G9N88ZV4t1KfaFeDzKRTfr")

	client := rpc.New(rpc.MainNetBetaSerum_RPC)
	ctx := context.Background()
	getRecentBlockHashResult, err := client.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}
	blockHash := getRecentBlockHashResult.Value.Blockhash

	builder := solana.NewTransactionBuilder()
	//
	/*
	tokenswap1 := solana.MustPublicKeyFromBase58("2Qqh2DS448qZMwhA8o5jBnSQF54uPPFHUJZULErA1or1")
	tokenswap1_authority, _, _ := solana.FindProgramAddress([][]byte{tokenswap1.Bytes()}, orcaProgramID)
	tokenswap1_user_authority := solana.MustPublicKeyFromBase58("F8Vyqk3unwxkXukZFQeYyGmFfTG3CAX4v24iyrjEYBJV")
	tokenswap1_user_account := solana.MustPublicKeyFromBase58("36c6YqAwyGKQG66XEp2dJc5JqjaBNv7sVghEtJv4c7u6")
	tokenswap1_tokenA := solana.MustPublicKeyFromBase58("Bm2Sj68iAWhmxmQfSLfPzbvDQsd9MAAiCrbDwA7KQ5Va")
	tokenswap1_tokenB := solana.MustPublicKeyFromBase58("DdcQxiv9d8JuQMYmerWDgxsciASqHxkEpymqVSkizZpc")
	tokneswap1_user2_account := solana.MustPublicKeyFromBase58("648kz2PXcnUwDWHqhu1k38gFLbnRMV88J9ByLtcesE2a")
	tokenswap1_pool := solana.MustPublicKeyFromBase58("E4cthfUFaDd4x5t1vbeBNBHm7isqhM8kapthPzPJz1M2")
	tokenswap1_fee_account := solana.MustPublicKeyFromBase58("GS6F9UV9TLcEW74LFCt3Fj5Lpvnfexngvj1VbAF5qUNv")
	data := make([]byte, 17)
	data[0] = 1
	binary.LittleEndian.PutUint64(data[1:], 100000)
	binary.LittleEndian.PutUint64(data[9:], 0)
	builder.AddInstruction(&orcaTransactionInstructions{
		accounts:  []*solana.AccountMeta{
			{PublicKey: tokenswap1, IsSigner: false, IsWritable: false},
			{PublicKey: tokenswap1_authority, IsSigner: true, IsWritable: false},
			{PublicKey: tokenswap1_user_authority, IsSigner: true, IsWritable: false},
			{PublicKey: tokenswap1_user_account, IsSigner: false, IsWritable: true},
			{PublicKey: tokenswap1_tokenA, IsSigner: false, IsWritable: true},
			{PublicKey: tokenswap1_tokenB, IsSigner: false, IsWritable: true},
			{PublicKey: tokneswap1_user2_account, IsSigner: false, IsWritable: true},
			{PublicKey: tokenswap1_pool, IsSigner: false, IsWritable: true},
			{PublicKey: tokenswap1_fee_account, IsSigner: false, IsWritable: true},
			{PublicKey: tokenProgramID, IsSigner: false, IsWritable: false},
			{},
		},
		data:      data,
		programID: orcaProgramID,
	})
	{
		getAccountInfoResponse , err := client.GetAccountInfo(ctx, tokneswap1_user2_account)
		if err != nil {
			panic(err)
		}
		rJson, _ := json.MarshalIndent(getAccountInfoResponse, "", "    ")
		fmt.Printf("getAccountInfoResponse: %s\n", string(rJson))
	}
	*/
	/*
		//3oRPcFaRHvv9pPR6nRasigVDkm3k9kTjdfjxUpgLV5Pq
		tokenswap2 := solana.MustPublicKeyFromBase58("3oRPcFaRHvv9pPR6nRasigVDkm3k9kTjdfjxUpgLV5Pq")
		tokenswap2_authority, _, _ := solana.FindProgramAddress([][]byte{tokenswap2.Bytes()}, orcaProgramID)
		tokenswap2_user_authority := solana.MustPublicKeyFromBase58("DkJ76rxxV4HnCm4TZAtPTykBjuthXZHjsk9NHyAUmUCT")
		tokenswap2_user_account := solana.MustPublicKeyFromBase58("648kz2PXcnUwDWHqhu1k38gFLbnRMV88J9ByLtcesE2a")
		tokenswap2_tokenA := solana.MustPublicKeyFromBase58("EvbbFpEwh142yj8GZbp2FAjiSNhr32eaVJ6d1ia8r9jr")
		tokenswap2_tokenB := solana.MustPublicKeyFromBase58("5d8G5r5xqTpQWVWbEeM93jE52ZXXrgQe8MkznX5BVRsV")
		tokneswap2_user2_account := solana.MustPublicKeyFromBase58("btetJ8RQZ448nKtc7mNWGvZQAUXTJpWuZDo5h9qLr3Z")
		tokenswap2_pool := solana.MustPublicKeyFromBase58("8qNqTaKKbdZuzQPWWXy5wNVkJh54ex8zvvnEnTFkrKMP")
		tokenswap2_fee_account := solana.MustPublicKeyFromBase58("9ZjnwbeTQwc4XQZsG2saqfyfyMxL2ypUsNteSRqtQkAg")
		data2 := make([]byte, 17)
		data2[0] = 1
		binary.LittleEndian.PutUint64(data2[1:], 15381)
		binary.LittleEndian.PutUint64(data2[9:], 0)
	builder.AddInstruction(&orcaTransactionInstructions{
		accounts:  []*solana.AccountMeta{
			{PublicKey: tokenswap2, IsSigner: false, IsWritable: false},
			{PublicKey: tokenswap2_authority, IsSigner: true, IsWritable: false},
			{PublicKey: tokenswap2_user_authority, IsSigner: true, IsWritable: false},
			{PublicKey: tokenswap2_user_account, IsSigner: false, IsWritable: true},
			{PublicKey: tokenswap2_tokenA, IsSigner: false, IsWritable: true},
			{PublicKey: tokenswap2_tokenB, IsSigner: false, IsWritable: true},
			{PublicKey: tokneswap2_user2_account, IsSigner: false, IsWritable: true},
			{PublicKey: tokenswap2_pool, IsSigner: false, IsWritable: true},
			{PublicKey: tokenswap2_fee_account, IsSigner: false, IsWritable: true},
			{PublicKey: tokenProgramID, IsSigner: false, IsWritable: false},
			{},
		},
		data:      data2,
		programID: orcaProgramID,
	})
		{
			getAccountInfoResponse , err := client.GetAccountInfo(ctx, tokneswap2_user2_account)
			if err != nil {
				panic(err)
			}
			rJson, _ := json.MarshalIndent(getAccountInfoResponse, "", "    ")
			fmt.Printf("getAccountInfoResponse: %s\n", string(rJson))
		}
	*/

		//tokenswap3 := solana.MustPublicKeyFromBase58("")
		tokenswap3 := solana.MustPublicKeyFromBase58("6fTRDD7sYxCN7oyoSQaN1AWC3P2m8A6gVZzGrpej9DvL")
		tokenswap3_authority, _, _ := solana.FindProgramAddress([][]byte{tokenswap3.Bytes()}, orcaProgramID)
		tokenswap3_user_authority := solana.MustPublicKeyFromBase58("HJEXUEyvgYuwm8rBoLPc7ZdVty8aPqo7MUw9cK4v6afi")
		tokenswap3_user_account := solana.MustPublicKeyFromBase58("btetJ8RQZ448nKtc7mNWGvZQAUXTJpWuZDo5h9qLr3Z")
		tokenswap3_tokenA := solana.MustPublicKeyFromBase58("7VcwKUtdKnvcgNhZt5BQHsbPrXLxhdVomsgrr7k2N5P5")
		tokenswap3_tokenB := solana.MustPublicKeyFromBase58("FdiTt7XQ94fGkgorywN1GuXqQzmURHCDgYtUutWRcy4q")
		tokneswap3_user2_account := solana.MustPublicKeyFromBase58("36c6YqAwyGKQG66XEp2dJc5JqjaBNv7sVghEtJv4c7u6")
		tokenswap3_pool := solana.MustPublicKeyFromBase58("ECFcUGwHHMaZynAQpqRHkYeTBnS5GnPWZywM8aggcs3A")
		tokenswap3_fee_account := solana.MustPublicKeyFromBase58("4pdzKqAGd1WbXn1L4UpY4r58irTfjFYMYNudBrqbQaYJ")
		data3 := make([]byte, 17)
		data3[0] = 1
		binary.LittleEndian.PutUint64(data3[1:], 15999)
		binary.LittleEndian.PutUint64(data3[9:], 0)
	builder.AddInstruction(&orcaTransactionInstructions{
		accounts:  []*solana.AccountMeta{
			{PublicKey: tokenswap3, IsSigner: false, IsWritable: false},
			{PublicKey: tokenswap3_authority, IsSigner: true, IsWritable: false},
			{PublicKey: tokenswap3_user_authority, IsSigner: true, IsWritable: false},
			{PublicKey: tokenswap3_user_account, IsSigner: false, IsWritable: true},
			{PublicKey: tokenswap3_tokenA, IsSigner: false, IsWritable: true},
			{PublicKey: tokenswap3_tokenB, IsSigner: false, IsWritable: true},
			{PublicKey: tokneswap3_user2_account, IsSigner: false, IsWritable: true},
			{PublicKey: tokenswap3_pool, IsSigner: false, IsWritable: true},
			{PublicKey: tokenswap3_fee_account, IsSigner: false, IsWritable: true},
			{PublicKey: tokenProgramID, IsSigner: false, IsWritable: false},
			{},
		},
		data:      data3,
		programID: orcaProgramID,
	})
		{
			getAccountInfoResponse , err := client.GetAccountInfo(ctx, tokneswap3_user2_account)
			if err != nil {
				panic(err)
			}
			rJson, _ := json.MarshalIndent(getAccountInfoResponse, "", "    ")
			fmt.Printf("getAccountInfoResponse: %s\n", string(rJson))
		}

	builder.SetRecentBlockHash(blockHash)
	builder.SetFeePayer(feePlayer)
	trx, err := builder.Build()
	if err != nil {
		panic(err)
	}
	trx.Sign(getKey)

	//
	txData, err := trx.MarshalBinary()
	if err != nil {
		panic(err)
	}
	base64Data := base64.StdEncoding.EncodeToString(txData)
	fmt.Printf("tx data: %s\n", base64Data)

	simulateTransactionResponse, err := client.SimulateTransactionWithOpts(ctx, trx, &rpc.SimulateTransactionOpts{
		SigVerify:              false,
		Commitment:             rpc.CommitmentFinalized,
		ReplaceRecentBlockhash: false,
		Accounts:               &rpc.SimulateTransactionAccountsOpts{
			Encoding: solana.EncodingBase64,
			Addresses: []solana.PublicKey{tokneswap3_user2_account},
		},
	})
	//simulateTransactionResponse, err := client.SimulateTransaction(ctx, trx)
	if err != nil {
		panic(err)
	}
	{
		rJson, _ := json.MarshalIndent(simulateTransactionResponse, "", "    ")
		fmt.Printf("SimulateTransactionResponse: %s\n", string(rJson))
	}
}
