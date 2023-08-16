package solana

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"math"
	"testing"
)

func TestSerum_ErrorCode(t *testing.T) {
	code := 0x1000ca4
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, uint32(code))
	file := data[3]
	line := binary.LittleEndian.Uint16(data[0:2])
	fmt.Printf("file: %d, line: %d\n", file, line)
}


func TestSerum_NewOrder(t *testing.T) {
	ins := make([]solana.Instruction, 0)
	//
	splTokenProgram := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	sysRent := solana.MustPublicKeyFromBase58("SysvarRent111111111111111111111111111111111")
	//userOwner := solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS")
	userOwner := solana.MustPublicKeyFromBase58("F8Vyqk3unwxkXukZFQeYyGmFfTG3CAX4v24iyrjEYBJV")
	//tokenUSDC := solana.MustPublicKeyFromBase58("C5dr2fk9zLvjuSjxe3tdvFurcZHKbfU1mHbpdv9Bfr7T")
	tokenUSDC := solana.MustPublicKeyFromBase58("8CFo8bL8mZQK8abbFyypFMwEDd8tVJjHTTojMLgQTUSZ")

	{
		// serum, usdc -> sol
		serumMarket := solana.MustPublicKeyFromBase58("9wFFyRfZBsuAha4YcuxcXLKwMxJR43S7fPfQLusDBzvT")
		openOrders := solana.MustPublicKeyFromBase58("fYaDiwbBY4ZWrAMnryNEZK2ADdHunWWycT7nDzqRAQJ")
		//authority := solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS")
		request_quene := solana.MustPublicKeyFromBase58("AZG3tFCFtiCqEwyardENBQNpHqxgzbMw8uKeZEw2nRG5")
		event_queue := solana.MustPublicKeyFromBase58("5KKsLVU6TcbVDK4BS6K1DGDxnh4Q9xjYJ8XaDCG5t8ht")
		bids := solana.MustPublicKeyFromBase58("14ivtgssEBoBjuZJtSAPKYgpUK7DmnSwuPMqJoVTSgKJ")
		asks := solana.MustPublicKeyFromBase58("CEQdAFKdycHugujQg9k2wbmxjcpdYZyVLfV9WerTnafJ")
		//user_source := solana.MustPublicKeyFromBase58("C5dr2fk9zLvjuSjxe3tdvFurcZHKbfU1mHbpdv9Bfr7T")
		//user_dst := solana.MustPublicKeyFromBase58("E2tAVqGR7XQwkjy9HJ7cPWu4SJN5oeLW1B7ZPJE5dWP3")
		base_vault := solana.MustPublicKeyFromBase58("36c6YqAwyGKQG66XEp2dJc5JqjaBNv7sVghEtJv4c7u6")
		quote_vault := solana.MustPublicKeyFromBase58("8CFo8bL8mZQK8abbFyypFMwEDd8tVJjHTTojMLgQTUSZ")
		//user_base := solana.MustPublicKeyFromBase58("E2tAVqGR7XQwkjy9HJ7cPWu4SJN5oeLW1B7ZPJE5dWP3")
		//user_quote := solana.MustPublicKeyFromBase58("C5dr2fk9zLvjuSjxe3tdvFurcZHKbfU1mHbpdv9Bfr7T")
		//vault_signer := solana.MustPublicKeyFromBase58("9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin")
		//
		side := uint32(0) // buy
		limitPrice := uint64(math.MaxUint64)
		maxBaseQuantity := uint64(math.MaxUint64)
		maxQuoteQuantity := uint64(1000000000)
		orderType := uint32(1)         // ImmediateOrCancel
		selfTradeBehavior := uint32(0) // decrementTake
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
		//
		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{
			},
			IsData:      data,
			IsProgramID: Serum,
		}
		//
		accounts := []*solana.AccountMeta{
			{PublicKey: serumMarket, IsSigner: false, IsWritable: true},
			{PublicKey: openOrders, IsSigner: false, IsWritable: true},
			{PublicKey: request_quene, IsSigner: false, IsWritable: true},
			{PublicKey: event_queue, IsSigner: false, IsWritable: true},
			{PublicKey: bids, IsSigner: false, IsWritable: true},
			{PublicKey: asks, IsSigner: false, IsWritable: true},
			{PublicKey: tokenUSDC, IsSigner: false, IsWritable: true},
			{PublicKey: userOwner, IsSigner: true, IsWritable: true},
			{PublicKey: base_vault, IsSigner: false, IsWritable: true},
			{PublicKey: quote_vault, IsSigner: false, IsWritable: true},
			{PublicKey: splTokenProgram, IsSigner: false, IsWritable: false},
			{PublicKey: sysRent, IsSigner: false, IsWritable: false},
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
	builder.SetFeePayer(Player)
	trx, err := builder.Build()
	if err != nil {
		panic(err)
	}
	trx.Sign(getKey)
	//
	response, err := rpcClient.SimulateTransactionWithOpts(ctx, trx, &rpc.SimulateTransactionOpts{
		SigVerify:              false,
		Commitment:             rpc.CommitmentFinalized,
		ReplaceRecentBlockhash: true,
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

func TestSerum_CreateOpenOrder(t *testing.T) {
	ctx := context.Background()
	rpcClient := rpc.New(rpc.MainNetBetaSerum_RPC)
	ins := make([]solana.Instruction, 0)


	// create a private key
	wallet := solana.NewWallet()
	openorder := wallet.PublicKey()
	Keys[openorder] = wallet.PrivateKey
	fmt.Printf("new open order: %s\n", openorder.String())
	fmt.Printf("new open order pri: %s\n", wallet.PrivateKey.String())

	// create account
	{
		space := uint64(3228)
		lamports, err := rpcClient.GetMinimumBalanceForRentExemption(ctx, space, rpc.CommitmentFinalized)
		if err != nil {
			panic(err)
		}
		data := make([]byte, 52)
		binary.LittleEndian.PutUint32(data[0:], 0)
		binary.LittleEndian.PutUint64(data[4:], lamports)
		binary.LittleEndian.PutUint64(data[12:], space)
		copy(data[20:], Serum.Bytes())
		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{
				{PublicKey: Player, IsSigner: true, IsWritable: true},
				{PublicKey: openorder, IsSigner: false, IsWritable: true},
			},
			IsData:      data,
			IsProgramID: System,
		}
		ins = append(ins, instruction)
	}
	{
		owner := solana.MustPublicKeyFromBase58("F8Vyqk3unwxkXukZFQeYyGmFfTG3CAX4v24iyrjEYBJV")
		serumMarket := solana.MustPublicKeyFromBase58("9wFFyRfZBsuAha4YcuxcXLKwMxJR43S7fPfQLusDBzvT")

		data := make([]byte, 5)
		data[0] = 0
		binary.LittleEndian.PutUint32(data[1:], uint32(15))
		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{
				{PublicKey: openorder, IsSigner: true, IsWritable: true},
				{PublicKey: owner, IsSigner: false, IsWritable: false},
				{PublicKey: serumMarket, IsSigner: false, IsWritable: true},
				{PublicKey: SysRent, IsSigner: false, IsWritable: false},
			},
			IsData:      data,
			IsProgramID: Serum,
		}
		ins = append(ins, instruction)
	}
	{
		//
		splTokenProgram := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
		sysRent := solana.MustPublicKeyFromBase58("SysvarRent111111111111111111111111111111111")
		//userOwner := solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS")
		userOwner := solana.MustPublicKeyFromBase58("F8Vyqk3unwxkXukZFQeYyGmFfTG3CAX4v24iyrjEYBJV")
		//tokenUSDC := solana.MustPublicKeyFromBase58("C5dr2fk9zLvjuSjxe3tdvFurcZHKbfU1mHbpdv9Bfr7T")
		tokenUSDC := solana.MustPublicKeyFromBase58("8CFo8bL8mZQK8abbFyypFMwEDd8tVJjHTTojMLgQTUSZ")


		// serum, usdc -> sol
		serumMarket := solana.MustPublicKeyFromBase58("9wFFyRfZBsuAha4YcuxcXLKwMxJR43S7fPfQLusDBzvT")
		//openOrders := solana.MustPublicKeyFromBase58("fYaDiwbBY4ZWrAMnryNEZK2ADdHunWWycT7nDzqRAQJ")
		openOrders := openorder
		//authority := solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS")
		request_quene := solana.MustPublicKeyFromBase58("AZG3tFCFtiCqEwyardENBQNpHqxgzbMw8uKeZEw2nRG5")
		event_queue := solana.MustPublicKeyFromBase58("5KKsLVU6TcbVDK4BS6K1DGDxnh4Q9xjYJ8XaDCG5t8ht")
		bids := solana.MustPublicKeyFromBase58("14ivtgssEBoBjuZJtSAPKYgpUK7DmnSwuPMqJoVTSgKJ")
		asks := solana.MustPublicKeyFromBase58("CEQdAFKdycHugujQg9k2wbmxjcpdYZyVLfV9WerTnafJ")
		//user_source := solana.MustPublicKeyFromBase58("C5dr2fk9zLvjuSjxe3tdvFurcZHKbfU1mHbpdv9Bfr7T")
		//user_dst := solana.MustPublicKeyFromBase58("E2tAVqGR7XQwkjy9HJ7cPWu4SJN5oeLW1B7ZPJE5dWP3")
		base_vault := solana.MustPublicKeyFromBase58("36c6YqAwyGKQG66XEp2dJc5JqjaBNv7sVghEtJv4c7u6")
		quote_vault := solana.MustPublicKeyFromBase58("8CFo8bL8mZQK8abbFyypFMwEDd8tVJjHTTojMLgQTUSZ")
		//user_base := solana.MustPublicKeyFromBase58("E2tAVqGR7XQwkjy9HJ7cPWu4SJN5oeLW1B7ZPJE5dWP3")
		//user_quote := solana.MustPublicKeyFromBase58("C5dr2fk9zLvjuSjxe3tdvFurcZHKbfU1mHbpdv9Bfr7T")
		//vault_signer := solana.MustPublicKeyFromBase58("9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin")
		//
		side := uint32(0) // buy
		limitPrice := uint64(math.MaxUint64)
		maxBaseQuantity := uint64(math.MaxUint64)
		maxQuoteQuantity := uint64(1000000000)
		orderType := uint32(1)         // ImmediateOrCancel
		selfTradeBehavior := uint32(0) // decrementTake
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
		//
		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{
			},
			IsData:      data,
			IsProgramID: Serum,
		}
		//
		accounts := []*solana.AccountMeta{
			{PublicKey: serumMarket, IsSigner: false, IsWritable: true},
			{PublicKey: openOrders, IsSigner: false, IsWritable: true},
			{PublicKey: request_quene, IsSigner: false, IsWritable: true},
			{PublicKey: event_queue, IsSigner: false, IsWritable: true},
			{PublicKey: bids, IsSigner: false, IsWritable: true},
			{PublicKey: asks, IsSigner: false, IsWritable: true},
			{PublicKey: tokenUSDC, IsSigner: false, IsWritable: true},
			{PublicKey: userOwner, IsSigner: false, IsWritable: true},
			{PublicKey: base_vault, IsSigner: false, IsWritable: true},
			{PublicKey: quote_vault, IsSigner: false, IsWritable: true},
			{PublicKey: splTokenProgram, IsSigner: false, IsWritable: false},
			{PublicKey: sysRent, IsSigner: false, IsWritable: false},
		}
		instruction.IsAccounts = append(instruction.IsAccounts, accounts...)
		//
		ins = append(ins, instruction)
	}

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

func TestSerum_SetFunds(t *testing.T) {
	ins := make([]solana.Instruction, 0)
	{
		splTokenProgram := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
		userOwner := solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS")
		//userOwner := solana.MustPublicKeyFromBase58("F8Vyqk3unwxkXukZFQeYyGmFfTG3CAX4v24iyrjEYBJV")
		//tokenUSDC := solana.MustPublicKeyFromBase58("C5dr2fk9zLvjuSjxe3tdvFurcZHKbfU1mHbpdv9Bfr7T")
		tokenUSDC := solana.MustPublicKeyFromBase58("8CFo8bL8mZQK8abbFyypFMwEDd8tVJjHTTojMLgQTUSZ")
		//
		tokenSol := solana.MustPublicKeyFromBase58("37FHaj6rRL4qhjXJN9PxeDRqyuTdeB5QXW5yPxiBJxmu")

		// serum, usdc -> sol
		serumMarket := solana.MustPublicKeyFromBase58("9wFFyRfZBsuAha4YcuxcXLKwMxJR43S7fPfQLusDBzvT")
		//openOrders := solana.MustPublicKeyFromBase58("fYaDiwbBY4ZWrAMnryNEZK2ADdHunWWycT7nDzqRAQJ")
		openOrders := solana.MustPublicKeyFromBase58("fYaDiwbBY4ZWrAMnryNEZK2ADdHunWWycT7nDzqRAQJ")
		//authority := solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS")
		//user_source := solana.MustPublicKeyFromBase58("C5dr2fk9zLvjuSjxe3tdvFurcZHKbfU1mHbpdv9Bfr7T")
		//user_dst := solana.MustPublicKeyFromBase58("E2tAVqGR7XQwkjy9HJ7cPWu4SJN5oeLW1B7ZPJE5dWP3")
		base_vault := solana.MustPublicKeyFromBase58("36c6YqAwyGKQG66XEp2dJc5JqjaBNv7sVghEtJv4c7u6")
		quote_vault := solana.MustPublicKeyFromBase58("8CFo8bL8mZQK8abbFyypFMwEDd8tVJjHTTojMLgQTUSZ")

		nonce := make([]byte, 8)
		binary.LittleEndian.PutUint64(nonce, uint64(1))
		vaultSigner, err := solana.CreateProgramAddress([][]byte{serumMarket.Bytes(), nonce}, Serum)
		if err != nil {
			panic(err)
		}

		data := make([]byte, 5)
		data[0] = 0
		binary.LittleEndian.PutUint32(data[1:], uint32(5))
		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{
				{PublicKey: serumMarket, IsSigner: false, IsWritable: true},
				{PublicKey: openOrders, IsSigner: false, IsWritable: true},
				{PublicKey: userOwner, IsSigner: true, IsWritable: false},
				{PublicKey: base_vault, IsSigner: false, IsWritable: true},
				{PublicKey: quote_vault, IsSigner: false, IsWritable: true},
				{PublicKey: tokenSol, IsSigner: false, IsWritable: true},
				{PublicKey: tokenUSDC, IsSigner: false, IsWritable: true},
				{PublicKey: vaultSigner, IsSigner: false, IsWritable: false},
				{PublicKey: splTokenProgram, IsSigner: false, IsWritable: false},
			},
			IsData:      data,
			IsProgramID: Serum,
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
