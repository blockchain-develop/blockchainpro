package solana

import (
	"context"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"math"
	"testing"
)
/*
func getKey(key solana.PublicKey) *solana.PrivateKey {
	wallet := solana.NewWallet()
	return &wallet.PrivateKey
}
*/

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
	//feePlayer := solana.MustPublicKeyFromBase58("Y2akr3bXHRsqyP1QJtbm9G9N88ZV4t1KfaFeDzKRTfr")

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
	//builder.SetFeePayer(feePlayer)
	trx, err := builder.Build()
	if err != nil {
		panic(err)
	}
	trx.Sign(getKey)

	//
	trxJson, _ := json.MarshalIndent(trx, "", "    ")
	fmt.Printf("transaction: %s", trxJson)

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


func TestSimulateTransaction3(t *testing.T) {
	orcaProgramID := solana.MustPublicKeyFromBase58("9W959DqEETiGZocYWCQPaJ6sBmUzgfxXfqGeTEdp3aQP")
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
	tokenswap3 := solana.MustPublicKeyFromBase58("EGZ7tiLeH62TPV1gL8WwbXGzEPa9zmcpVnnkPKKnrE2U")
	tokenswap3_authority, _, _ := solana.FindProgramAddress([][]byte{tokenswap3.Bytes()}, orcaProgramID)
	tokenswap3_user_authority := solana.MustPublicKeyFromBase58("CzZAjoEqA6sjqtaiZiPqDkmxG6UuZWxwRWCenbBMc8Xz")
	tokenswap3_user_account := solana.MustPublicKeyFromBase58("29cTsXahEoEBwbHwVc59jToybFpagbBMV6Lh45pWEmiK")
	tokenswap3_tokenA := solana.MustPublicKeyFromBase58("ANP74VNsHwSrq9uUSjiSNyNWvf6ZPrKTmE4gHoNd13Lg")
	tokenswap3_tokenB := solana.MustPublicKeyFromBase58("75HgnSvXbWKZBpZHveX68ZzAhDqMzNDS29X6BGLtxMo1")
	tokneswap3_user2_account := solana.MustPublicKeyFromBase58("4F6F9dNRMvsGr1kGPg2N1YUgR1zH9VUie4ETGKmZAiXN")
	tokenswap3_pool := solana.MustPublicKeyFromBase58("APDFRM3HMr8CAGXwKHiu2f5ePSpaiEJhaURwhsRrUUt9")
	tokenswap3_fee_account := solana.MustPublicKeyFromBase58("8JnSiuvQq3BVuCU3n4DrSTw9chBSPvEMswrhtifVkr1o")
	data3 := make([]byte, 17)
	data3[0] = 1
	binary.LittleEndian.PutUint64(data3[1:], 1000000000)
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
	trxJson, _ := json.MarshalIndent(trx, "", "    ")
	fmt.Printf("transaction: %s", trxJson)

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



func TestSimulateTransaction_Serum(t *testing.T) {
	serumProgram := solana.MustPublicKeyFromBase58("9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin")
	market := solana.MustPublicKeyFromBase58("HxFLKUAmAMLz1jtT3hbvCMELwH5H9tpM2QugP8sKyfhW")
	tokenProgramID := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	user := solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS")
	usdc_account := solana.MustPublicKeyFromBase58("C5dr2fk9zLvjuSjxe3tdvFurcZHKbfU1mHbpdv9Bfr7T")
	atlas_account := solana.MustPublicKeyFromBase58("2WfUhA3UtiJDTxWsvUjbk45qc61w3FUpZ38hKRMmMkwL")


	client := rpc.New(rpc.MainNetBetaSerum_RPC)
	ctx := context.Background()
	getRecentBlockHashResult, err := client.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}
	blockHash := getRecentBlockHashResult.Value.Blockhash

	//
	{
		getAccountInfoResponse , err := client.GetAccountInfo(ctx, atlas_account)
		if err != nil {
			panic(err)
		}
		xx := decodeAccount(getAccountInfoResponse.Value)
		fmt.Printf("account: %s, amount: %d\n", atlas_account.String(), xx.Amount)
	}

	{
		getAccountInfoResponse , err := client.GetAccountInfo(ctx, usdc_account)
		if err != nil {
			panic(err)
		}
		xx := decodeAccount(getAccountInfoResponse.Value)
		fmt.Printf("account: %s, amount: %d\n", usdc_account.String(), xx.Amount)
	}

	builder := solana.NewTransactionBuilder()
	{
		//
		orderType := uint32(1)         // ImmediateOrCancel
		selfTradeBehavior := uint32(0) // decrementTake
		side := uint32(0)
		limitPrice := uint64(math.MaxUint64)
		maxBaseQuantity := uint64(math.MaxUint64)
		maxQuoteQuantity := uint64(10000000000)
		data := make([]byte, 51)
		data[0] = 0
		binary.LittleEndian.PutUint32(data[1:], uint32(10))
		binary.LittleEndian.PutUint32(data[5:], side)
		binary.LittleEndian.PutUint64(data[9:], limitPrice)
		binary.LittleEndian.PutUint64(data[17:], maxBaseQuantity)
		binary.LittleEndian.PutUint64(data[25:], maxQuoteQuantity)
		binary.LittleEndian.PutUint32(data[33:], selfTradeBehavior)
		binary.LittleEndian.PutUint32(data[37:], orderType)
		binary.LittleEndian.PutUint64(data[41:], 0)
		binary.LittleEndian.PutUint16(data[49:], 65535)

		openOrdersKey := solana.MustPublicKeyFromBase58("AvigzQtp1YpSS3Wd3Kpz7Z1QpR651bnKkpvoutXK3tX3")
		RequestQueue := solana.MustPublicKeyFromBase58("88HGNyFaHEcNLfSbGQE4knRQ8gnjC3jo6YvPRAv477Uw")
		EventQueue := solana.MustPublicKeyFromBase58("qeQC4u5vpo5QMC17V5UMkQfK67vu3DHtBYVT1hFSGCK")
		Bids := solana.MustPublicKeyFromBase58("Bc5wovapX1tRjZfyZVpsGH73Gq5LGN4ANsj8kaEhfY7c")
		Asks := solana.MustPublicKeyFromBase58("4EHg2ANFFEKLFkpLxgiyinJ1UDWsG2p8rVoAjFfjMDKc")
		BaseVault := solana.MustPublicKeyFromBase58("5XQ7xYE3ujVA21HGbvFGVG4pLgqVHSfR9anz2EfmZ3nA")
		QuoteVault := solana.MustPublicKeyFromBase58("ArUDWPwzGQFfa7t7nSdkp1Dj6tYA3icXEq8K7goz9WoG")
		builder.AddInstruction(&Instruction{
			IsAccounts: []*solana.AccountMeta{
				{PublicKey: market, IsSigner: false, IsWritable: true},
				{PublicKey: openOrdersKey, IsSigner: false, IsWritable: true},
				{PublicKey: RequestQueue, IsSigner: false, IsWritable: true},
				{PublicKey: EventQueue, IsSigner: false, IsWritable: true},
				{PublicKey: Bids, IsSigner: false, IsWritable: true},
				{PublicKey: Asks, IsSigner: false, IsWritable: true},
				{PublicKey: usdc_account, IsSigner: false, IsWritable: true},
				{PublicKey: user, IsSigner: true, IsWritable: false},
				{PublicKey: BaseVault, IsSigner: false, IsWritable: true},
				{PublicKey: QuoteVault, IsSigner: false, IsWritable: true},
				{PublicKey: tokenProgramID, IsSigner: false, IsWritable: false},
				{PublicKey: SysRent, IsSigner: false, IsWritable: false},
			},
			IsData:      data,
			IsProgramID: serumProgram,
		})
	}
	{
		//
		data := make([]byte, 5)
		data[0] = 0
		binary.LittleEndian.PutUint32(data[1:], uint32(5))

		nonce := make([]byte, 8)
		binary.LittleEndian.PutUint64(nonce, 0)
		vaultSigner, err := solana.CreateProgramAddress([][]byte{market.Bytes(), nonce}, serumProgram)
		if err != nil {
			panic(err)
		}
		openOrdersKey := solana.MustPublicKeyFromBase58("AvigzQtp1YpSS3Wd3Kpz7Z1QpR651bnKkpvoutXK3tX3")
		BaseVault := solana.MustPublicKeyFromBase58("5XQ7xYE3ujVA21HGbvFGVG4pLgqVHSfR9anz2EfmZ3nA")
		QuoteVault := solana.MustPublicKeyFromBase58("ArUDWPwzGQFfa7t7nSdkp1Dj6tYA3icXEq8K7goz9WoG")
		builder.AddInstruction(&Instruction{
			IsAccounts: []*solana.AccountMeta{
				{PublicKey: market, IsSigner: false, IsWritable: true},
				{PublicKey: openOrdersKey, IsSigner: false, IsWritable: true},
				{PublicKey: user, IsSigner: true, IsWritable: false},
				{PublicKey: BaseVault, IsSigner: false, IsWritable: true},
				{PublicKey: QuoteVault, IsSigner: false, IsWritable: true},
				{PublicKey: atlas_account, IsSigner: false, IsWritable: true},
				{PublicKey: usdc_account, IsSigner: false, IsWritable: true},
				{PublicKey: vaultSigner, IsSigner: false, IsWritable: false},
				{PublicKey: tokenProgramID, IsSigner: false, IsWritable: false},
			},
			IsData:      data,
			IsProgramID: serumProgram,
		})
	}

	builder.SetRecentBlockHash(blockHash)
	builder.SetFeePayer(user)
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

	if true {
		simulateTransactionResponse, err := client.SimulateTransactionWithOpts(ctx, trx, &rpc.SimulateTransactionOpts{
			SigVerify:              true,
			Commitment:             rpc.CommitmentFinalized,
			Accounts: &rpc.SimulateTransactionAccountsOpts{
				Encoding:  solana.EncodingBase64,
				Addresses: []solana.PublicKey{atlas_account, usdc_account},
			},
		})
		//simulateTransactionResponse, err := client.SimulateTransaction(ctx, trx)
		if err != nil {
			panic(err)
		}
		fmt.Printf("simulate logs: %s\n", simulateTransactionResponse.Value.Logs)
		{
			xx := decodeAccount(simulateTransactionResponse.Value.Accounts[0])
			fmt.Printf("account: %s, amount: %d\n", atlas_account.String(), xx.Amount)
		}

		{
			xx := decodeAccount(simulateTransactionResponse.Value.Accounts[1])
			fmt.Printf("account: %s, amount: %d\n", usdc_account.String(), xx.Amount)
		}
	}
	if false {
		hash, err := client.SendTransaction(ctx, trx)
		//simulateTransactionResponse, err := client.SimulateTransaction(ctx, trx)
		if err != nil {
			panic(err)
		}
		fmt.Printf("hash: %s\n", hash.String())
	}
}



func TestSimulateTransaction_Serum_InArb(t *testing.T) {
	client := rpc.New(rpc.MainNetBetaSerum_RPC)
	ctx := context.Background()
	getRecentBlockHashResult, err := client.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}
	blockHash := getRecentBlockHashResult.Value.Blockhash

	//
	arbitrage := solana.MustPublicKeyFromBase58("7H4ShpibmzrKS8yPJX9wi1ZyrRYzw5tLym7RjWvAxcHA")
	exchange := solana.MustPublicKeyFromBase58("HhUVfHYvGby6k7zHrAcmA52YQLB7sWD41wkcb1WyUw8Z")
	//
	splTokenProgram := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	sysRent := solana.MustPublicKeyFromBase58("SysvarRent111111111111111111111111111111111")
	userOwner := solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS")
	tokenUSDC := solana.MustPublicKeyFromBase58("C5dr2fk9zLvjuSjxe3tdvFurcZHKbfU1mHbpdv9Bfr7T")
	tokenPolis := solana.MustPublicKeyFromBase58("2WfUhA3UtiJDTxWsvUjbk45qc61w3FUpZ38hKRMmMkwL")
	//tokenUSDT := solana.MustPublicKeyFromBase58("EDASG9z1jHYMxDqda9JeFM9E7UJAEg7j9LqJtqevwKL5")

	{
		getAccountInfoResponse , err := client.GetAccountInfo(ctx, tokenPolis)
		if err != nil {
			panic(err)
		}
		xx := decodeAccount(getAccountInfoResponse.Value)
		fmt.Printf("account: %s, amount: %d\n", tokenPolis.String(), xx.Amount)
	}

	{
		getAccountInfoResponse , err := client.GetAccountInfo(ctx, tokenUSDC)
		if err != nil {
			panic(err)
		}
		xx := decodeAccount(getAccountInfoResponse.Value)
		fmt.Printf("account: %s, amount: %d\n", tokenUSDC.String(), xx.Amount)
	}

	builder := solana.NewTransactionBuilder()
	{
		// serum, usdc -> sol
		serum := solana.MustPublicKeyFromBase58("9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin")
		serumMarket := solana.MustPublicKeyFromBase58("HxFLKUAmAMLz1jtT3hbvCMELwH5H9tpM2QugP8sKyfhW")
		openOrders := solana.MustPublicKeyFromBase58("AvigzQtp1YpSS3Wd3Kpz7Z1QpR651bnKkpvoutXK3tX3")
		//authority := solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS")
		request_quene := solana.MustPublicKeyFromBase58("88HGNyFaHEcNLfSbGQE4knRQ8gnjC3jo6YvPRAv477Uw")
		event_queue := solana.MustPublicKeyFromBase58("qeQC4u5vpo5QMC17V5UMkQfK67vu3DHtBYVT1hFSGCK")
		bids := solana.MustPublicKeyFromBase58("Bc5wovapX1tRjZfyZVpsGH73Gq5LGN4ANsj8kaEhfY7c")
		asks := solana.MustPublicKeyFromBase58("4EHg2ANFFEKLFkpLxgiyinJ1UDWsG2p8rVoAjFfjMDKc")
		//user_source := solana.MustPublicKeyFromBase58("C5dr2fk9zLvjuSjxe3tdvFurcZHKbfU1mHbpdv9Bfr7T")
		//user_dst := solana.MustPublicKeyFromBase58("E2tAVqGR7XQwkjy9HJ7cPWu4SJN5oeLW1B7ZPJE5dWP3")
		base_vault := solana.MustPublicKeyFromBase58("5XQ7xYE3ujVA21HGbvFGVG4pLgqVHSfR9anz2EfmZ3nA")
		quote_vault := solana.MustPublicKeyFromBase58("ArUDWPwzGQFfa7t7nSdkp1Dj6tYA3icXEq8K7goz9WoG")
		//user_base := solana.MustPublicKeyFromBase58("E2tAVqGR7XQwkjy9HJ7cPWu4SJN5oeLW1B7ZPJE5dWP3")
		//user_quote := solana.MustPublicKeyFromBase58("C5dr2fk9zLvjuSjxe3tdvFurcZHKbfU1mHbpdv9Bfr7T")
		//vault_signer := solana.MustPublicKeyFromBase58("9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin")
		nonce := make([]byte, 8)
		binary.LittleEndian.PutUint64(nonce, 0)
		vaultSigner, err := solana.CreateProgramAddress([][]byte{serumMarket.Bytes(), nonce}, serum)
		if err != nil {
			panic(err)
		}
		//
		amountIn := uint64(5000000000)
		data := make([]byte, 12)
		data[0] = 0
		binary.LittleEndian.PutUint64(data[1:], amountIn)
		data[9] = 2
		data[10] = 0 // buy
		data[11] = 0
		//
		builder.AddInstruction(&Instruction{
			IsAccounts: []*solana.AccountMeta{
				{PublicKey: exchange, IsSigner: false, IsWritable: true},
				{PublicKey: serum, IsSigner: false, IsWritable: false},
				{PublicKey: serumMarket, IsSigner: false, IsWritable: true},
				{PublicKey: openOrders, IsSigner: false, IsWritable: true},
				{PublicKey: userOwner, IsSigner: true, IsWritable: false},
				{PublicKey: request_quene, IsSigner: false, IsWritable: true},
				{PublicKey: event_queue, IsSigner: false, IsWritable: true},
				{PublicKey: bids, IsSigner: false, IsWritable: true},
				{PublicKey: asks, IsSigner: false, IsWritable: true},
				{PublicKey: tokenUSDC, IsSigner: false, IsWritable: true},
				{PublicKey: tokenPolis, IsSigner: false, IsWritable: true},
				{PublicKey: base_vault, IsSigner: false, IsWritable: true},
				{PublicKey: quote_vault, IsSigner: false, IsWritable: true},
				{PublicKey: tokenPolis, IsSigner: false, IsWritable: true},
				{PublicKey: tokenUSDC, IsSigner: false, IsWritable: true},
				{PublicKey: vaultSigner, IsSigner: false, IsWritable: false},
				{PublicKey: splTokenProgram, IsSigner: false, IsWritable: false},
				{PublicKey: sysRent, IsSigner: false, IsWritable: false},
			},
			IsData:      data,
			IsProgramID: arbitrage,
		})
	}
	//
	builder.SetRecentBlockHash(blockHash)
	builder.SetFeePayer(userOwner)
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

	if true {
		simulateTransactionResponse, err := client.SimulateTransactionWithOpts(ctx, trx, &rpc.SimulateTransactionOpts{
			SigVerify:              true,
			Commitment:             rpc.CommitmentFinalized,
			Accounts: &rpc.SimulateTransactionAccountsOpts{
				Encoding:  solana.EncodingBase64,
				Addresses: []solana.PublicKey{tokenPolis, tokenUSDC},
			},
		})
		//simulateTransactionResponse, err := client.SimulateTransaction(ctx, trx)
		if err != nil {
			panic(err)
		}
		fmt.Printf("simulate logs: %s\n", simulateTransactionResponse.Value.Logs)
		{
			xx := decodeAccount(simulateTransactionResponse.Value.Accounts[0])
			fmt.Printf("account: %s, amount: %d\n", tokenPolis.String(), xx.Amount)
		}

		{
			xx := decodeAccount(simulateTransactionResponse.Value.Accounts[1])
			fmt.Printf("account: %s, amount: %d\n", tokenUSDC.String(), xx.Amount)
		}
	}
}


