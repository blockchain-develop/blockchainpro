package solana

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"math"
	"testing"
)

func TestMongo_PlaceOrder_BuySOL(t *testing.T) {
	ins := make([]solana.Instruction, 0)
	//
	mongo_group := solana.MustPublicKeyFromBase58("98pjRuQjK3qA6gXts96PqZT4Ze5QmnCmt3QYjhbUSPue")
	mongo_account := solana.MustPublicKeyFromBase58("2AhmR9rjGyiiHEExmC6w5wMTvUe5HVCKrzsdm1b24xkP")
	owner := solana.MustPublicKeyFromBase58("3pfNpRNu31FBzx84TnefG6iBkSqQxGtuL5G5v9aaxyv8")
	//sol_account := solana.MustPublicKeyFromBase58("AVn3JRGhifPCxjxZsU3tQuo4U4dTHizHzBDGW983tx47")
	//
	mongo_cache := solana.MustPublicKeyFromBase58("EBDRoayCDDUvDgCimta45ajQeXbexv7aKqJubruqpyvu")
	serum_dex := solana.MustPublicKeyFromBase58("9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin")
	market := solana.MustPublicKeyFromBase58("9wFFyRfZBsuAha4YcuxcXLKwMxJR43S7fPfQLusDBzvT")
	bids := solana.MustPublicKeyFromBase58("14ivtgssEBoBjuZJtSAPKYgpUK7DmnSwuPMqJoVTSgKJ")
	asks := solana.MustPublicKeyFromBase58("CEQdAFKdycHugujQg9k2wbmxjcpdYZyVLfV9WerTnafJ")
	request_queue := solana.MustPublicKeyFromBase58("AZG3tFCFtiCqEwyardENBQNpHqxgzbMw8uKeZEw2nRG5")
	event_queue := solana.MustPublicKeyFromBase58("5KKsLVU6TcbVDK4BS6K1DGDxnh4Q9xjYJ8XaDCG5t8ht")
	//
	mongo_base := solana.MustPublicKeyFromBase58("36c6YqAwyGKQG66XEp2dJc5JqjaBNv7sVghEtJv4c7u6")
	mongo_quote := solana.MustPublicKeyFromBase58("8CFo8bL8mZQK8abbFyypFMwEDd8tVJjHTTojMLgQTUSZ")
	//
	base_root_bank := solana.MustPublicKeyFromBase58("7jH1uLmiB2zbHNe6juZZYjQCrvquakTwd3yMaQpeP8rR")
	base_node_bank := solana.MustPublicKeyFromBase58("2bqJYcA1A8gw4qJFjyE2G4akiUunpd9rP6QzfnxHqSqr")
	base_vault := solana.MustPublicKeyFromBase58("AVn3JRGhifPCxjxZsU3tQuo4U4dTHizHzBDGW983tx47")
	quote_root_bank := solana.MustPublicKeyFromBase58("AMzanZxMirPCgGcBoH9kw4Jzi9LFMomyUCXbpzDeL2T8")
	quote_node_bank := solana.MustPublicKeyFromBase58("BGcwkj1WudQwUUjFk78hAjwd1uAm8trh1N4CJSa51euh")
	quote_vault := solana.MustPublicKeyFromBase58("8Vw25ZackDzaJzzBBqcgcpDsCsDfRSkMGgwFQ3gbReWF")
	//
	token_program := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	mongo_signer := solana.MustPublicKeyFromBase58("9BVcYqEQxyccuwznvxXqDkSJFavvTyheiTYk231T1A8S")
	dex_signer := solana.MustPublicKeyFromBase58("F8Vyqk3unwxkXukZFQeYyGmFfTG3CAX4v24iyrjEYBJV")
	srm_vault := solana.MustPublicKeyFromBase58("HiB9JtxgnA7G29EWVcPFhXbLCTaPitdVJKgD3BhH6TJj")
	//
	//open_order := solana.MustPublicKeyFromBase58("CyJ19NKrrHgo1j3KS49NJPhUzNZ73U4cqnabVvtuJbX8")
	marketIndex := uint64(3)
	marketIndexBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(marketIndexBytes, marketIndex)
	open_order, _, err := solana.FindProgramAddress([][]byte{mongo_account.Bytes(), marketIndexBytes, []byte("OpenOrders")}, MongoV3)
	if err != nil {
		panic(err)
	}
	// create spot open order account
	{
		data := make([]byte, 4)
		binary.LittleEndian.PutUint32(data[0:], uint32(60))
		//
		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{},
			IsData:      data,
			IsProgramID: MongoV3,
		}
		//
		accounts := []*solana.AccountMeta{
			{PublicKey: mongo_group, IsSigner: false, IsWritable: false},
			{PublicKey: mongo_account, IsSigner: false, IsWritable: true},
			{PublicKey: owner, IsSigner: true, IsWritable: false},
			{PublicKey: serum_dex, IsSigner: false, IsWritable: false},
			{PublicKey: open_order, IsSigner: false, IsWritable: true},
			{PublicKey: market, IsSigner: false, IsWritable: false},
			{PublicKey: mongo_signer, IsSigner: false, IsWritable: false},
			{PublicKey: System, IsSigner: false, IsWritable: false},
		}
		instruction.IsAccounts = append(instruction.IsAccounts, accounts...)
		ins = append(ins, instruction)
	}
	// palce order
	{
		//
		side := uint32(0) // buy
		//limitPrice := uint64(math.MaxUint64)
		limitPrice := uint64(32000)
		maxBaseQuantity := uint64(math.MaxUint64)
		maxQuoteQuantity := uint64(1000000000)
		orderType := uint32(1)         // ImmediateOrCancel
		//orderType := uint32(2)
		selfTradeBehavior := uint32(0) // decrementTake
		data := make([]byte, 46)
		binary.LittleEndian.PutUint32(data[0:], side)
		binary.LittleEndian.PutUint64(data[4:], limitPrice)
		binary.LittleEndian.PutUint64(data[12:], maxBaseQuantity)
		binary.LittleEndian.PutUint64(data[20:], maxQuoteQuantity)
		binary.LittleEndian.PutUint32(data[28:], selfTradeBehavior)
		binary.LittleEndian.PutUint32(data[32:], orderType)
		binary.LittleEndian.PutUint64(data[36:], 0)
		binary.LittleEndian.PutUint16(data[44:], 65535)
		//
		mongo_data := make([]byte, 50)
		binary.LittleEndian.PutUint32(mongo_data[0:], uint32(41))
		copy(mongo_data[4:], data)
		fmt.Printf("%s\n", hex.EncodeToString(mongo_data))
		//
		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{},
			IsData:      mongo_data,
			IsProgramID: MongoV3,
		}
		//
		accounts := []*solana.AccountMeta{
			{PublicKey: mongo_group, IsSigner: false, IsWritable: false},
			{PublicKey: mongo_account, IsSigner: false, IsWritable: true},
			{PublicKey: owner, IsSigner: true, IsWritable: false},
			{PublicKey: mongo_cache, IsSigner: false, IsWritable: false},
			{PublicKey: serum_dex, IsSigner: false, IsWritable: false},
			{PublicKey: market, IsSigner: false, IsWritable: true},
			{PublicKey: bids, IsSigner: false, IsWritable: true},
			{PublicKey: asks, IsSigner: false, IsWritable: true},
			{PublicKey: request_queue, IsSigner: false, IsWritable: true},
			{PublicKey: event_queue, IsSigner: false, IsWritable: true},
			{PublicKey: mongo_base, IsSigner: false, IsWritable: true},
			{PublicKey: mongo_quote, IsSigner: false, IsWritable: true},
			{PublicKey: base_root_bank, IsSigner: false, IsWritable: false},
			{PublicKey: base_node_bank, IsSigner: false, IsWritable: true},
			{PublicKey: base_vault, IsSigner: false, IsWritable: true},
			{PublicKey: quote_root_bank, IsSigner: false, IsWritable: false},
			{PublicKey: quote_node_bank, IsSigner: false, IsWritable: true},
			{PublicKey: quote_vault, IsSigner: false, IsWritable: true},
			{PublicKey: token_program, IsSigner: false, IsWritable: false},
			{PublicKey: mongo_signer, IsSigner: false, IsWritable: false},
			{PublicKey: dex_signer, IsSigner: false, IsWritable: false},
			{PublicKey: srm_vault, IsSigner: false, IsWritable: false},
			{PublicKey: open_order, IsSigner: false, IsWritable: true},
		}
		instruction.IsAccounts = append(instruction.IsAccounts, accounts...)
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
	builder.SetFeePayer(owner)
	trx, err := builder.Build()
	if err != nil {
		panic(err)
	}
	trx.Sign(getKey)
	//
	trxJson, _ := json.MarshalIndent(trx, "", "    ")
	fmt.Printf("tx: %s\n", trxJson)
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


func TestMongo_PlaceOrder_SellSOL(t *testing.T) {
	ins := make([]solana.Instruction, 0)
	//
	owner := solana.MustPublicKeyFromBase58("9BVcYqEQxyccuwznvxXqDkSJFavvTyheiTYk231T1A8S")
	sol_account := solana.MustPublicKeyFromBase58("AVn3JRGhifPCxjxZsU3tQuo4U4dTHizHzBDGW983tx47")
	//
	mongo_group := solana.MustPublicKeyFromBase58("98pjRuQjK3qA6gXts96PqZT4Ze5QmnCmt3QYjhbUSPue")
	//open order account
	accountNum := uint64(1)
	accountNumBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(accountNumBytes, accountNum)
	mongo_account, _, err := solana.FindProgramAddress([][]byte{mongo_group.Bytes(), owner.Bytes(), accountNumBytes}, MongoV3)
	if err != nil {
		panic(err)
	}
	//open order account
	marketIndex := uint64(3)
	marketIndexBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(marketIndexBytes, marketIndex)
	open_order, _, err := solana.FindProgramAddress([][]byte{mongo_account.Bytes(), marketIndexBytes, []byte("OpenOrders")}, MongoV3)
	if err != nil {
		panic(err)
	}
	//
	mongo_cache := solana.MustPublicKeyFromBase58("EBDRoayCDDUvDgCimta45ajQeXbexv7aKqJubruqpyvu")
	serum_dex := solana.MustPublicKeyFromBase58("9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin")
	market := solana.MustPublicKeyFromBase58("9wFFyRfZBsuAha4YcuxcXLKwMxJR43S7fPfQLusDBzvT")
	serum_bids := solana.MustPublicKeyFromBase58("14ivtgssEBoBjuZJtSAPKYgpUK7DmnSwuPMqJoVTSgKJ")
	serum_asks := solana.MustPublicKeyFromBase58("CEQdAFKdycHugujQg9k2wbmxjcpdYZyVLfV9WerTnafJ")
	request_queue := solana.MustPublicKeyFromBase58("AZG3tFCFtiCqEwyardENBQNpHqxgzbMw8uKeZEw2nRG5")
	event_queue := solana.MustPublicKeyFromBase58("5KKsLVU6TcbVDK4BS6K1DGDxnh4Q9xjYJ8XaDCG5t8ht")
	//
	mongo_base := solana.MustPublicKeyFromBase58("36c6YqAwyGKQG66XEp2dJc5JqjaBNv7sVghEtJv4c7u6")
	mongo_quote := solana.MustPublicKeyFromBase58("8CFo8bL8mZQK8abbFyypFMwEDd8tVJjHTTojMLgQTUSZ")
	//
	base_root_bank := solana.MustPublicKeyFromBase58("7jH1uLmiB2zbHNe6juZZYjQCrvquakTwd3yMaQpeP8rR")
	base_node_bank := solana.MustPublicKeyFromBase58("2bqJYcA1A8gw4qJFjyE2G4akiUunpd9rP6QzfnxHqSqr")
	base_vault := solana.MustPublicKeyFromBase58("AVn3JRGhifPCxjxZsU3tQuo4U4dTHizHzBDGW983tx47")
	quote_root_bank := solana.MustPublicKeyFromBase58("AMzanZxMirPCgGcBoH9kw4Jzi9LFMomyUCXbpzDeL2T8")
	quote_node_bank := solana.MustPublicKeyFromBase58("BGcwkj1WudQwUUjFk78hAjwd1uAm8trh1N4CJSa51euh")
	quote_vault := solana.MustPublicKeyFromBase58("8Vw25ZackDzaJzzBBqcgcpDsCsDfRSkMGgwFQ3gbReWF")
	//
	token_program := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	mongo_signer := solana.MustPublicKeyFromBase58("9BVcYqEQxyccuwznvxXqDkSJFavvTyheiTYk231T1A8S")
	dex_signer := solana.MustPublicKeyFromBase58("F8Vyqk3unwxkXukZFQeYyGmFfTG3CAX4v24iyrjEYBJV")
	srm_vault := solana.MustPublicKeyFromBase58("HiB9JtxgnA7G29EWVcPFhXbLCTaPitdVJKgD3BhH6TJj")
	//
	deposit_amount := uint64(10000000000000)
	sell_amount := uint64(80000000000000)
	//
	// create mongo account
	{
		data := make([]byte, 12)
		binary.LittleEndian.PutUint32(data[0:], uint32(55))
		binary.LittleEndian.PutUint64(data[4:], uint64(1))
		//
		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{},
			IsData:      data,
			IsProgramID: MongoV3,
		}
		//
		accounts := []*solana.AccountMeta{
			{PublicKey: mongo_group, IsSigner: false, IsWritable: true},
			{PublicKey: mongo_account, IsSigner: false, IsWritable: true},
			{PublicKey: owner, IsSigner: true, IsWritable: false},
			{PublicKey: System, IsSigner: false, IsWritable: false},
		}
		instruction.IsAccounts = append(instruction.IsAccounts, accounts...)
		ins = append(ins, instruction)
	}
	// deposit
	{
		data := make([]byte, 12)
		binary.LittleEndian.PutUint32(data[0:], uint32(2))
		binary.LittleEndian.PutUint64(data[4:], deposit_amount)
		//
		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{},
			IsData:      data,
			IsProgramID: MongoV3,
		}
		//
		accounts := []*solana.AccountMeta{
			{PublicKey: mongo_group, IsSigner: false, IsWritable: false},
			{PublicKey: mongo_account, IsSigner: false, IsWritable: true},
			{PublicKey: owner, IsSigner: true, IsWritable: false},
			{PublicKey: mongo_cache, IsSigner: false, IsWritable: true},
			{PublicKey: base_root_bank, IsSigner: false, IsWritable: true},
			{PublicKey: base_node_bank, IsSigner: false, IsWritable: true},
			{PublicKey: base_vault, IsSigner: false, IsWritable: true},
			{PublicKey: token_program, IsSigner: false, IsWritable: false},
			{PublicKey: sol_account, IsSigner: false, IsWritable: true},
		}
		instruction.IsAccounts = append(instruction.IsAccounts, accounts...)
		ins = append(ins, instruction)
	}
	// create spot open order account
	{
		data := make([]byte, 4)
		binary.LittleEndian.PutUint32(data[0:], uint32(60))
		//
		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{},
			IsData:      data,
			IsProgramID: MongoV3,
		}
		//
		accounts := []*solana.AccountMeta{
			{PublicKey: mongo_group, IsSigner: false, IsWritable: false},
			{PublicKey: mongo_account, IsSigner: false, IsWritable: true},
			{PublicKey: owner, IsSigner: true, IsWritable: false},
			{PublicKey: serum_dex, IsSigner: false, IsWritable: false},
			{PublicKey: open_order, IsSigner: false, IsWritable: true},
			{PublicKey: market, IsSigner: false, IsWritable: false},
			{PublicKey: mongo_signer, IsSigner: false, IsWritable: false},
			{PublicKey: System, IsSigner: false, IsWritable: false},
		}
		instruction.IsAccounts = append(instruction.IsAccounts, accounts...)
		ins = append(ins, instruction)
	}
	// palce sell sol order
	{
		//
		side := uint32(1) // sell
		//limitPrice := uint64(math.MaxUint64)
		limitPrice := uint64(10000)
		maxBaseQuantity := sell_amount / 100000000
		maxQuoteQuantity := uint64(math.MaxUint64)
		orderType := uint32(1)         // ImmediateOrCancel
		selfTradeBehavior := uint32(0) // decrementTake
		data := make([]byte, 46)
		binary.LittleEndian.PutUint32(data[0:], side)
		binary.LittleEndian.PutUint64(data[4:], limitPrice)
		binary.LittleEndian.PutUint64(data[12:], maxBaseQuantity)
		binary.LittleEndian.PutUint64(data[20:], maxQuoteQuantity)
		binary.LittleEndian.PutUint32(data[28:], selfTradeBehavior)
		binary.LittleEndian.PutUint32(data[32:], orderType)
		binary.LittleEndian.PutUint64(data[36:], 0)
		binary.LittleEndian.PutUint16(data[44:], 65535)
		//
		mongo_data := make([]byte, 50)
		binary.LittleEndian.PutUint32(mongo_data[0:], uint32(41))
		copy(mongo_data[4:], data)
		fmt.Printf("%s\n", hex.EncodeToString(mongo_data))
		//
		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{},
			IsData:      mongo_data,
			IsProgramID: MongoV3,
		}
		//
		accounts := []*solana.AccountMeta{
			{PublicKey: mongo_group, IsSigner: false, IsWritable: false},
			{PublicKey: mongo_account, IsSigner: false, IsWritable: true},
			{PublicKey: owner, IsSigner: true, IsWritable: false},
			{PublicKey: mongo_cache, IsSigner: false, IsWritable: false},
			{PublicKey: serum_dex, IsSigner: false, IsWritable: false},
			{PublicKey: market, IsSigner: false, IsWritable: true},
			{PublicKey: serum_bids, IsSigner: false, IsWritable: true},
			{PublicKey: serum_asks, IsSigner: false, IsWritable: true},
			{PublicKey: request_queue, IsSigner: false, IsWritable: true},
			{PublicKey: event_queue, IsSigner: false, IsWritable: true},
			{PublicKey: mongo_base, IsSigner: false, IsWritable: true},
			{PublicKey: mongo_quote, IsSigner: false, IsWritable: true},
			{PublicKey: base_root_bank, IsSigner: false, IsWritable: false},
			{PublicKey: base_node_bank, IsSigner: false, IsWritable: true},
			{PublicKey: base_vault, IsSigner: false, IsWritable: true},
			{PublicKey: quote_root_bank, IsSigner: false, IsWritable: false},
			{PublicKey: quote_node_bank, IsSigner: false, IsWritable: true},
			{PublicKey: quote_vault, IsSigner: false, IsWritable: true},
			{PublicKey: token_program, IsSigner: false, IsWritable: false},
			{PublicKey: mongo_signer, IsSigner: false, IsWritable: false},
			{PublicKey: dex_signer, IsSigner: false, IsWritable: false},
			{PublicKey: srm_vault, IsSigner: false, IsWritable: false},
			{PublicKey: open_order, IsSigner: false, IsWritable: true},
		}
		instruction.IsAccounts = append(instruction.IsAccounts, accounts...)
		ins = append(ins, instruction)
	}
	//
	ctx := context.Background()
	rpcClient := rpc.New(rpc.MainNetBeta_RPC)
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
	builder.SetFeePayer(owner)
	trx, err := builder.Build()
	if err != nil {
		panic(err)
	}
	trx.Sign(getKey)
	//
	trxJson, _ := json.MarshalIndent(trx, "", "    ")
	fmt.Printf("tx: %s\n", trxJson)
	{
		fmt.Printf("state before transaction\n")
		getMultipleAccountsRsp, err := rpcClient.GetMultipleAccountsWithOpts(ctx, []solana.PublicKey{serum_bids, serum_asks, base_vault, quote_vault},
			&rpc.GetMultipleAccountsOpts{Encoding: solana.EncodingBase64})
		if err != nil {
			panic(err)
		}
		{
			bidsAccount := getMultipleAccountsRsp.Value[0]
			bid := decodeOrderBookAccount(bidsAccount)
			xx := asks(bid, 1)
			bidJson, _ := json.MarshalIndent(xx, "", "    ")
			fmt.Printf("bid: %s\n", string(bidJson))
		}
		{
			bidsAccount := getMultipleAccountsRsp.Value[1]
			bid := decodeOrderBookAccount(bidsAccount)
			xx := bids(bid, 1)
			bidJson, _ := json.MarshalIndent(xx, "", "    ")
			fmt.Printf("ask: %s\n", string(bidJson))
		}
		{
			bidsAccount := getMultipleAccountsRsp.Value[2]
			bid := decodeAccount(bidsAccount)
			bidJson, _ := json.MarshalIndent(bid, "", "    ")
			fmt.Printf("base vault: %s\n", string(bidJson))
		}
		{
			bidsAccount := getMultipleAccountsRsp.Value[3]
			bid := decodeAccount(bidsAccount)
			bidJson, _ := json.MarshalIndent(bid, "", "    ")
			fmt.Printf("quote vault: %s\n", string(bidJson))
		}
	}
	//
	response, err := rpcClient.SimulateTransactionWithOpts(ctx, trx, &rpc.SimulateTransactionOpts{
		SigVerify:              false,
		Commitment:             rpc.CommitmentFinalized,
		ReplaceRecentBlockhash: true,
		Accounts: &rpc.SimulateTransactionAccountsOpts{
			solana.EncodingBase64,
			[]solana.PublicKey {serum_bids, serum_asks, base_vault, quote_vault},
		}})
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
	//
	{
		fmt.Printf("state after transaction\n")
		{
			bidsAccount := response.Value.Accounts[0]
			bid := decodeOrderBookAccount(bidsAccount)
			xx := asks(bid, 1)
			bidJson, _ := json.MarshalIndent(xx, "", "    ")
			fmt.Printf("bid: %s\n", string(bidJson))
		}
		{
			bidsAccount := response.Value.Accounts[1]
			bid := decodeOrderBookAccount(bidsAccount)
			xx := bids(bid, 1)
			bidJson, _ := json.MarshalIndent(xx, "", "    ")
			fmt.Printf("ask: %s\n", string(bidJson))
		}
		{
			bidsAccount := response.Value.Accounts[2]
			bid := decodeAccount(bidsAccount)
			bidJson, _ := json.MarshalIndent(bid, "", "    ")
			fmt.Printf("base vault: %s\n", string(bidJson))
		}
		{
			bidsAccount := response.Value.Accounts[3]
			bid := decodeAccount(bidsAccount)
			bidJson, _ := json.MarshalIndent(bid, "", "    ")
			fmt.Printf("quote vault: %s\n", string(bidJson))
		}
	}
	fmt.Printf("SimulateTransactionWithOpts successful!\n")
	return
}
