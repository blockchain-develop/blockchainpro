package solana

import (
	"encoding/binary"
	"github.com/gagliardetto/solana-go"
	"testing"
)

var (
	//ARB = solana.MustPublicKeyFromBase58("8xfm8n7D3RN9jJ8DFzs3QEUqyqzdgdYi39o3cePf7SSw")
	//ARB = solana.MustPublicKeyFromBase58("A8uqVMNXnYE5vNc1Fkbpd144ezvkqW58wMvYUYhD9rHo")
	ARB = solana.MustPublicKeyFromBase58("7H4ShpibmzrKS8yPJX9wi1ZyrRYzw5tLym7RjWvAxcHA")

	Orca = solana.MustPublicKeyFromBase58("SwaPpA9LAaLfeLi3a68M4DjnLqgtticKg6CnyNwgAC8")
	//Saber = solana.MustPublicKeyFromBase58("")
	Serum = solana.MustPublicKeyFromBase58("9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin")
	//
	MongoV3 = solana.MustPublicKeyFromBase58("mv3ekLzLbnVPNxjSKvqBpU3ZeZXPQdEC3bp5MDEBG68")
)

func ArbitrageTokenSwap() []solana.Instruction {
	amountIn := uint64(100000)
	data := make([]byte, 18)
	data[0] = 0
	binary.LittleEndian.PutUint64(data[1:], amountIn)
	data[9] = 1
	//
	program := ARB
	tokenSwap := solana.MustPublicKeyFromBase58("9bHt39hwT8yXECSknLXbkGAAEDDdmLQodGtXd8NHD52u")
	market := solana.MustPublicKeyFromBase58("BF5BXLK5SEerE6fMRtdEmodJ51NV1weqaoicmohZQwGp")
	authority, _, err := solana.FindProgramAddress([][]byte{market.Bytes()}, tokenSwap)
	if err != nil {
		panic(err)
	}
	userOwner := solana.MustPublicKeyFromBase58("8ZqMeFqC2tAGPpywsyA9apwnKMJcZqFpgBbh6tN88jNq")
	userSrc := solana.MustPublicKeyFromBase58("6iY3fHNZVVhzjXkjZDVyNzavTAPEW9sZMvfoz44E4biv")
	tokenSwapSrc := solana.MustPublicKeyFromBase58("617iL9wJkwZ2LCrFtb74DgCLKqPqCctyUtXiUhxesSq2")
	tokenSwapDst := solana.MustPublicKeyFromBase58("7xp1c5BGDcBWg16eeHCZcwXHFP2mrrrJDXegSuKKdkpP")
	userDst := solana.MustPublicKeyFromBase58("BncwuoeqtYpR5S9SgqeLddZzjSGAW194mAu94KQX7NAL")
	poolMint := solana.MustPublicKeyFromBase58("AirFfe8hej7WMZExQkVLhPrhRcxkLeLNSKeWXFnDTPnA")
	poolAccount := solana.MustPublicKeyFromBase58("DHXQLCbFRRAJH3JyZSShH8ev5Qbn1GgvDCRdLtYH6hXs")
	splTokenProgram := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")

	//

	instruction := &Instruction{
		IsAccounts: []*solana.AccountMeta{
			{PublicKey: tokenSwap, IsSigner: false, IsWritable: false},
			{PublicKey: market, IsSigner: false, IsWritable: false},
			{PublicKey: authority, IsSigner: false, IsWritable: false},
			{PublicKey: userOwner, IsSigner: true, IsWritable: false},
			{PublicKey: userSrc, IsSigner: false, IsWritable: true},
			{PublicKey: tokenSwapSrc, IsSigner: false, IsWritable: true},
			{PublicKey: tokenSwapDst, IsSigner: false, IsWritable: true},
			{PublicKey: userDst, IsSigner: false, IsWritable: true},
			{PublicKey: poolMint, IsSigner: false, IsWritable: true},
			{PublicKey: poolAccount, IsSigner: false, IsWritable: true},
			{PublicKey: splTokenProgram, IsSigner: false, IsWritable: false},
		},
		IsData:      data,
		IsProgramID: program,
	}
	return []solana.Instruction{instruction}
}

func TestArbitrageTokenSwap(t *testing.T) {
	ins := ArbitrageTokenSwap()
	SendTransaction(ins)
}



func ArbitrageSaberSwap() []solana.Instruction {
	//
	ins := make([]solana.Instruction, 0)
	//
	{
		amountIn := uint64(100000)
		data := make([]byte, 12)
		data[0] = 0
		binary.LittleEndian.PutUint64(data[1:], amountIn)
		data[9] = 1
		data[10] = 0
		data[11] = 0
		//
		program := ARB
		exchange := solana.MustPublicKeyFromBase58("CE2w3KYUUYYAjpJ2QeMMQCRT8Fe7s2fhqyQbhhnWnrRe")
		saberSwap := solana.MustPublicKeyFromBase58("SSwpkEEcbUqx4vtoEByFjSkhKdCT862DNVb52nZg1UZ")
		market := solana.MustPublicKeyFromBase58("B94iYzzWe7Q3ksvRnt5yJm6G5YquerRFKpsUVUvasdmA")
		authority, _, err := solana.FindProgramAddress([][]byte{market.Bytes()}, saberSwap)
		if err != nil {
			panic(err)
		}
		userOwner := solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS")
		userSrc := solana.MustPublicKeyFromBase58("HRsjKYQoEzJkpuhb6YKzDvHtha55iMfUZanmBH7PwuMY")
		tokenSwapSrc := solana.MustPublicKeyFromBase58("DjX8KKu5bHz3zz7oJDhZbSDoksGbDr6EZFaooYFrPK4u")
		tokenSwapDst := solana.MustPublicKeyFromBase58("654z3VDWzK7BuehVQSmyftqm6TxHnJDDJF8eHdCXUEcs")
		userDst := solana.MustPublicKeyFromBase58("FoKc24SoEmEVr2wub6g9RkHCKi1jGHTssVooaYEC2YDS")
		feeAdmin := solana.MustPublicKeyFromBase58("2v5K4N1D8Jx5BXpbLHBs8zc7xh1Li2QWL4ZDCU89cp9R")
		splTokenProgram := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")

		//
		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{
				{PublicKey: exchange, IsSigner: false, IsWritable: true},
				{PublicKey: saberSwap, IsSigner: false, IsWritable: false},
				{PublicKey: market, IsSigner: false, IsWritable: false},
				{PublicKey: authority, IsSigner: false, IsWritable: false},
				{PublicKey: userOwner, IsSigner: true, IsWritable: false},
				{PublicKey: userSrc, IsSigner: false, IsWritable: true},
				{PublicKey: tokenSwapSrc, IsSigner: false, IsWritable: true},
				{PublicKey: tokenSwapDst, IsSigner: false, IsWritable: true},
				{PublicKey: userDst, IsSigner: false, IsWritable: true},
				{PublicKey: feeAdmin, IsSigner: false, IsWritable: true},
				{PublicKey: splTokenProgram, IsSigner: false, IsWritable: false},
				{PublicKey: SysClock, IsSigner: false, IsWritable: false},
			},
			IsData:      data,
			IsProgramID: program,
		}
		ins = append(ins, instruction)
	}

	{
		amountIn := uint64(0)
		data := make([]byte, 12)
		data[0] = 0
		binary.LittleEndian.PutUint64(data[1:], amountIn)
		data[9] = 1
		data[10] = 0
		data[11] = 1
		//
		program := ARB
		exchange := solana.MustPublicKeyFromBase58("CE2w3KYUUYYAjpJ2QeMMQCRT8Fe7s2fhqyQbhhnWnrRe")
		saberSwap := solana.MustPublicKeyFromBase58("SSwpkEEcbUqx4vtoEByFjSkhKdCT862DNVb52nZg1UZ")
		market := solana.MustPublicKeyFromBase58("B94iYzzWe7Q3ksvRnt5yJm6G5YquerRFKpsUVUvasdmA")
		authority, _, err := solana.FindProgramAddress([][]byte{market.Bytes()}, saberSwap)
		if err != nil {
			panic(err)
		}
		userOwner := solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS")
		userSrc := solana.MustPublicKeyFromBase58("HRsjKYQoEzJkpuhb6YKzDvHtha55iMfUZanmBH7PwuMY")
		tokenSwapSrc := solana.MustPublicKeyFromBase58("DjX8KKu5bHz3zz7oJDhZbSDoksGbDr6EZFaooYFrPK4u")
		tokenSwapDst := solana.MustPublicKeyFromBase58("654z3VDWzK7BuehVQSmyftqm6TxHnJDDJF8eHdCXUEcs")
		userDst := solana.MustPublicKeyFromBase58("FoKc24SoEmEVr2wub6g9RkHCKi1jGHTssVooaYEC2YDS")
		feeAdmin := solana.MustPublicKeyFromBase58("2v5K4N1D8Jx5BXpbLHBs8zc7xh1Li2QWL4ZDCU89cp9R")
		splTokenProgram := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")

		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{
				{PublicKey: exchange, IsSigner: false, IsWritable: true},
				{PublicKey: saberSwap, IsSigner: false, IsWritable: false},
				{PublicKey: market, IsSigner: false, IsWritable: false},
				{PublicKey: authority, IsSigner: false, IsWritable: false},
				{PublicKey: userOwner, IsSigner: true, IsWritable: false},
				{PublicKey: userSrc, IsSigner: false, IsWritable: true},
				{PublicKey: tokenSwapSrc, IsSigner: false, IsWritable: true},
				{PublicKey: tokenSwapDst, IsSigner: false, IsWritable: true},
				{PublicKey: userDst, IsSigner: false, IsWritable: true},
				{PublicKey: feeAdmin, IsSigner: false, IsWritable: true},
				{PublicKey: splTokenProgram, IsSigner: false, IsWritable: false},
				{PublicKey: SysClock, IsSigner: false, IsWritable: false},
			},
			IsData:      data,
			IsProgramID: program,
		}
		ins = append(ins, instruction)
	}

	{
		amountIn := uint64(0)
		data := make([]byte, 12)
		data[0] = 0
		binary.LittleEndian.PutUint64(data[1:], amountIn)
		data[9] = 1
		data[10] = 0
		data[11] = 2
		//
		program := ARB
		exchange := solana.MustPublicKeyFromBase58("CE2w3KYUUYYAjpJ2QeMMQCRT8Fe7s2fhqyQbhhnWnrRe")
		saberSwap := solana.MustPublicKeyFromBase58("SSwpkEEcbUqx4vtoEByFjSkhKdCT862DNVb52nZg1UZ")
		market := solana.MustPublicKeyFromBase58("B94iYzzWe7Q3ksvRnt5yJm6G5YquerRFKpsUVUvasdmA")
		authority, _, err := solana.FindProgramAddress([][]byte{market.Bytes()}, saberSwap)
		if err != nil {
			panic(err)
		}
		userOwner := solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS")
		userSrc := solana.MustPublicKeyFromBase58("HRsjKYQoEzJkpuhb6YKzDvHtha55iMfUZanmBH7PwuMY")
		tokenSwapSrc := solana.MustPublicKeyFromBase58("DjX8KKu5bHz3zz7oJDhZbSDoksGbDr6EZFaooYFrPK4u")
		tokenSwapDst := solana.MustPublicKeyFromBase58("654z3VDWzK7BuehVQSmyftqm6TxHnJDDJF8eHdCXUEcs")
		userDst := solana.MustPublicKeyFromBase58("FoKc24SoEmEVr2wub6g9RkHCKi1jGHTssVooaYEC2YDS")
		feeAdmin := solana.MustPublicKeyFromBase58("2v5K4N1D8Jx5BXpbLHBs8zc7xh1Li2QWL4ZDCU89cp9R")
		splTokenProgram := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")

		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{
				{PublicKey: exchange, IsSigner: false, IsWritable: true},
				{PublicKey: saberSwap, IsSigner: false, IsWritable: false},
				{PublicKey: market, IsSigner: false, IsWritable: false},
				{PublicKey: authority, IsSigner: false, IsWritable: false},
				{PublicKey: userOwner, IsSigner: true, IsWritable: false},
				{PublicKey: userSrc, IsSigner: false, IsWritable: true},
				{PublicKey: tokenSwapSrc, IsSigner: false, IsWritable: true},
				{PublicKey: tokenSwapDst, IsSigner: false, IsWritable: true},
				{PublicKey: userDst, IsSigner: false, IsWritable: true},
				{PublicKey: feeAdmin, IsSigner: false, IsWritable: true},
				{PublicKey: splTokenProgram, IsSigner: false, IsWritable: false},
				{PublicKey: SysClock, IsSigner: false, IsWritable: false},
			},
			IsData:      data,
			IsProgramID: program,
		}
		ins = append(ins, instruction)
	}

	return ins
}

func TestArbitrageSaberSwap(t *testing.T) {
	ins := ArbitrageSaberSwap()
	SendTransaction(ins)
}


func ArbitrageAll() []solana.Instruction {
	//
	ins := make([]solana.Instruction, 0)
	//
	arbitrage := solana.MustPublicKeyFromBase58("7H4ShpibmzrKS8yPJX9wi1ZyrRYzw5tLym7RjWvAxcHA")
	exchange := solana.MustPublicKeyFromBase58("HhUVfHYvGby6k7zHrAcmA52YQLB7sWD41wkcb1WyUw8Z")
	splTokenProgram := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	sysRent := solana.MustPublicKeyFromBase58("SysvarRent111111111111111111111111111111111")
	userOwner := solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS")
	tokenUSDC := solana.MustPublicKeyFromBase58("C5dr2fk9zLvjuSjxe3tdvFurcZHKbfU1mHbpdv9Bfr7T")
	tokenSol := solana.MustPublicKeyFromBase58("E2tAVqGR7XQwkjy9HJ7cPWu4SJN5oeLW1B7ZPJE5dWP3")
	tokenUSDT := solana.MustPublicKeyFromBase58("EDASG9z1jHYMxDqda9JeFM9E7UJAEg7j9LqJtqevwKL5")

	{
		// saber, usdc -> usdt
		saberSwap := solana.MustPublicKeyFromBase58("SSwpkEEcbUqx4vtoEByFjSkhKdCT862DNVb52nZg1UZ")
		saberMarket := solana.MustPublicKeyFromBase58("YAkoNb6HKmSxQN9L8hiBE5tPJRsniSSMzND1boHmZxe")
		saberAuthority, _, err := solana.FindProgramAddress([][]byte{saberMarket.Bytes()}, saberSwap)
		if err != nil {
			panic(err)
		}
		//userOwner := solana.MustPublicKeyFromBase58("sjcCd2jMG9BfmwwTDiitmJp8vSTJEF4M2wuBm3hR1M9")
		//userSrc := solana.MustPublicKeyFromBase58("92ev3jbJowvFpi6BaaXS3GkzjZhiwaFQ2nxUs8i61YKL")
		saberSwapSrc := solana.MustPublicKeyFromBase58("CfWX7o2TswwbxusJ4hCaPobu2jLCb1hfXuXJQjVq3jQF")
		saberSwapDst := solana.MustPublicKeyFromBase58("EnTrdMMpdhugeH6Ban6gYZWXughWxKtVGfCwFn78ZmY3")
		//userDst := solana.MustPublicKeyFromBase58("9aqNrRocaMkFE2qeq685JVcxpdHpcmeHF5pJGSs6d5or")
		feeAdmin := solana.MustPublicKeyFromBase58("2SL8iP8EjnUr6qTkbkfZt9tauXwJgc4GKXkYCCbLGbVP")
		//splTokenProgram := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
		//
		amountIn := uint64(100000000)
		data := make([]byte, 12)
		data[0] = 0
		binary.LittleEndian.PutUint64(data[1:], amountIn)
		data[9] = 1
		data[10] = 0
		data[11] = 0
		//
		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{
			},
			IsData:      data,
			IsProgramID: arbitrage,
		}
		accounts := []*solana.AccountMeta{
			{PublicKey: exchange, IsSigner: false, IsWritable: true},
			{PublicKey: saberSwap, IsSigner: false, IsWritable: false},
			{PublicKey: saberMarket, IsSigner: false, IsWritable: false},
			{PublicKey: saberAuthority, IsSigner: false, IsWritable: false},
			{PublicKey: userOwner, IsSigner: true, IsWritable: false},
			{PublicKey: tokenUSDC, IsSigner: false, IsWritable: true},
			{PublicKey: saberSwapSrc, IsSigner: false, IsWritable: true},
			{PublicKey: saberSwapDst, IsSigner: false, IsWritable: true},
			{PublicKey: tokenUSDT, IsSigner: false, IsWritable: true},
			{PublicKey: feeAdmin, IsSigner: false, IsWritable: true},
			{PublicKey: splTokenProgram, IsSigner: false, IsWritable: false},
			{PublicKey: SysClock, IsSigner: false, IsWritable: false},
		}
		instruction.IsAccounts = append(instruction.IsAccounts, accounts...)
		ins = append(ins, instruction)
	}

	{
		// orca, usdt -> sol
		tokenSwap := solana.MustPublicKeyFromBase58("9W959DqEETiGZocYWCQPaJ6sBmUzgfxXfqGeTEdp3aQP")
		tokenSwapMarket := solana.MustPublicKeyFromBase58("Dqk7mHQBx2ZWExmyrR2S8X6UG75CrbbpK2FSBZsNYsw6")
		tokenSwapAuthority, _, err := solana.FindProgramAddress([][]byte{tokenSwapMarket.Bytes()}, tokenSwap)
		if err != nil {
			panic(err)
		}
		//userOwner := solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS")
		//userSrc := solana.MustPublicKeyFromBase58("E2tAVqGR7XQwkjy9HJ7cPWu4SJN5oeLW1B7ZPJE5dWP3")
		tokenSwapSrc := solana.MustPublicKeyFromBase58("E8erPjPEorykpPjFV9yUYMYigEWKQUxuGfL2rJKLJ3KU")
		tokenSwapDst := solana.MustPublicKeyFromBase58("DTb8NKsfhEJGY1TrA7RXN6MBiTrjnkdMAfjPEjtmTT3M")
		//userDst := solana.MustPublicKeyFromBase58("EDASG9z1jHYMxDqda9JeFM9E7UJAEg7j9LqJtqevwKL5")
		poolMint := solana.MustPublicKeyFromBase58("FZthQCuYHhcfiDma7QrX7buDHwrZEd7vL8SjS6LQa3Tx")
		poolAccount := solana.MustPublicKeyFromBase58("BBKgw75FivTYXj85D2AWyVdaTdTWuSuHVXRm1Xu7fipb")
		//
		//
		amountIn := uint64(0)
		data := make([]byte, 12)
		data[0] = 0
		binary.LittleEndian.PutUint64(data[1:], amountIn)
		data[9] = 0
		data[10] = 0
		data[11] = 1
		//
		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{
			},
			IsData:      data,
			IsProgramID: arbitrage,
		}
		//
		accounts := []*solana.AccountMeta{
			{PublicKey: exchange, IsSigner: false, IsWritable: true},
			{PublicKey: tokenSwap, IsSigner: false, IsWritable: false},
			{PublicKey: tokenSwapMarket, IsSigner: false, IsWritable: false},
			{PublicKey: tokenSwapAuthority, IsSigner: false, IsWritable: false},
			{PublicKey: userOwner, IsSigner: true, IsWritable: false},
			{PublicKey: tokenUSDT, IsSigner: false, IsWritable: true},
			{PublicKey: tokenSwapSrc, IsSigner: false, IsWritable: true},
			{PublicKey: tokenSwapDst, IsSigner: false, IsWritable: true},
			{PublicKey: tokenSol, IsSigner: false, IsWritable: true},
			{PublicKey: poolMint, IsSigner: false, IsWritable: true},
			{PublicKey: poolAccount, IsSigner: false, IsWritable: true},
			{PublicKey: splTokenProgram, IsSigner: false, IsWritable: false},
		}
		instruction.IsAccounts = append(instruction.IsAccounts, accounts...)
		ins = append(ins, instruction)
	}

	{
		// serum, sol -> usdc
		serum := solana.MustPublicKeyFromBase58("9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin")
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
		nonce := make([]byte, 8)
		binary.LittleEndian.PutUint64(nonce, 1)
		vaultSigner, err := solana.CreateProgramAddress([][]byte{serumMarket.Bytes(), nonce}, serum)
		if err != nil {
			panic(err)
		}
		//
		amountIn := uint64(0)
		data := make([]byte, 12)
		data[0] = 0
		binary.LittleEndian.PutUint64(data[1:], amountIn)
		data[9] = 2
		data[10] = 1 // sell
		data[11] = 2
		//
		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{
			},
			IsData:      data,
			IsProgramID: arbitrage,
		}
		//
		accounts := []*solana.AccountMeta{
			{PublicKey: exchange, IsSigner: false, IsWritable: true},
			{PublicKey: serum, IsSigner: false, IsWritable: false},
			{PublicKey: serumMarket, IsSigner: false, IsWritable: true},
			{PublicKey: openOrders, IsSigner: false, IsWritable: true},
			{PublicKey: userOwner, IsSigner: true, IsWritable: false},
			{PublicKey: request_quene, IsSigner: false, IsWritable: true},
			{PublicKey: event_queue, IsSigner: false, IsWritable: true},
			{PublicKey: bids, IsSigner: false, IsWritable: true},
			{PublicKey: asks, IsSigner: false, IsWritable: true},
			{PublicKey: tokenSol, IsSigner: false, IsWritable: true},
			{PublicKey: tokenUSDC, IsSigner: false, IsWritable: true},
			{PublicKey: base_vault, IsSigner: false, IsWritable: true},
			{PublicKey: quote_vault, IsSigner: false, IsWritable: true},
			{PublicKey: tokenSol, IsSigner: false, IsWritable: true},
			{PublicKey: tokenUSDC, IsSigner: false, IsWritable: true},
			{PublicKey: vaultSigner, IsSigner: false, IsWritable: false},
			{PublicKey: splTokenProgram, IsSigner: false, IsWritable: false},
			{PublicKey: sysRent, IsSigner: false, IsWritable: false},
		}
		instruction.IsAccounts = append(instruction.IsAccounts, accounts...)
		ins = append(ins, instruction)
	}

	return ins
}

func TestArbitrageAll(t *testing.T) {
	ins := ArbitrageAll()
	SendTransaction(ins)
}


func ArbitrageAll1() []solana.Instruction {
	ins := make([]solana.Instruction, 0)
	//
	arbitrage := solana.MustPublicKeyFromBase58("7H4ShpibmzrKS8yPJX9wi1ZyrRYzw5tLym7RjWvAxcHA")
	exchange := solana.MustPublicKeyFromBase58("HhUVfHYvGby6k7zHrAcmA52YQLB7sWD41wkcb1WyUw8Z")
	//
	splTokenProgram := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	sysRent := solana.MustPublicKeyFromBase58("SysvarRent111111111111111111111111111111111")
	userOwner := solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS")
	tokenUSDC := solana.MustPublicKeyFromBase58("C5dr2fk9zLvjuSjxe3tdvFurcZHKbfU1mHbpdv9Bfr7T")
	tokenSol := solana.MustPublicKeyFromBase58("E2tAVqGR7XQwkjy9HJ7cPWu4SJN5oeLW1B7ZPJE5dWP3")
	tokenUSDT := solana.MustPublicKeyFromBase58("EDASG9z1jHYMxDqda9JeFM9E7UJAEg7j9LqJtqevwKL5")

	{
		// serum, usdc -> sol
		serum := solana.MustPublicKeyFromBase58("9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin")
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
		nonce := make([]byte, 8)
		binary.LittleEndian.PutUint64(nonce, 1)
		vaultSigner, err := solana.CreateProgramAddress([][]byte{serumMarket.Bytes(), nonce}, serum)
		if err != nil {
			panic(err)
		}
		//
		amountIn := uint64(1000000000)
		data := make([]byte, 12)
		data[0] = 0
		binary.LittleEndian.PutUint64(data[1:], amountIn)
		data[9] = 2
		data[10] = 0 // buy
		data[11] = 0
		//
		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{
			},
			IsData:      data,
			IsProgramID: arbitrage,
		}
		//
		accounts := []*solana.AccountMeta{
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
			{PublicKey: tokenSol, IsSigner: false, IsWritable: true},
			{PublicKey: base_vault, IsSigner: false, IsWritable: true},
			{PublicKey: quote_vault, IsSigner: false, IsWritable: true},
			{PublicKey: tokenSol, IsSigner: false, IsWritable: true},
			{PublicKey: tokenUSDC, IsSigner: false, IsWritable: true},
			{PublicKey: vaultSigner, IsSigner: false, IsWritable: false},
			{PublicKey: splTokenProgram, IsSigner: false, IsWritable: false},
			{PublicKey: sysRent, IsSigner: false, IsWritable: false},
		}
		instruction.IsAccounts = append(instruction.IsAccounts, accounts...)
		ins = append(ins, instruction)
	}

	{
		//
		amountIn := uint64(0)
		data := make([]byte, 12)
		data[0] = 0
		binary.LittleEndian.PutUint64(data[1:], amountIn)
		data[9] = 0
		data[10] = 0
		data[11] = 1
		//
		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{
			},
			IsData:      data,
			IsProgramID: arbitrage,
		}
		// orca, sol -> usdt
		tokenSwap := solana.MustPublicKeyFromBase58("9W959DqEETiGZocYWCQPaJ6sBmUzgfxXfqGeTEdp3aQP")
		tokenSwapMarket := solana.MustPublicKeyFromBase58("Dqk7mHQBx2ZWExmyrR2S8X6UG75CrbbpK2FSBZsNYsw6")
		tokenSwapAuthority, _, err := solana.FindProgramAddress([][]byte{tokenSwapMarket.Bytes()}, tokenSwap)
		if err != nil {
			panic(err)
		}
		//userOwner := solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS")
		//userSrc := solana.MustPublicKeyFromBase58("E2tAVqGR7XQwkjy9HJ7cPWu4SJN5oeLW1B7ZPJE5dWP3")
		tokenSwapSrc := solana.MustPublicKeyFromBase58("DTb8NKsfhEJGY1TrA7RXN6MBiTrjnkdMAfjPEjtmTT3M")
		tokenSwapDst := solana.MustPublicKeyFromBase58("E8erPjPEorykpPjFV9yUYMYigEWKQUxuGfL2rJKLJ3KU")
		//userDst := solana.MustPublicKeyFromBase58("EDASG9z1jHYMxDqda9JeFM9E7UJAEg7j9LqJtqevwKL5")
		poolMint := solana.MustPublicKeyFromBase58("FZthQCuYHhcfiDma7QrX7buDHwrZEd7vL8SjS6LQa3Tx")
		poolAccount := solana.MustPublicKeyFromBase58("BBKgw75FivTYXj85D2AWyVdaTdTWuSuHVXRm1Xu7fipb")
		//
		accounts := []*solana.AccountMeta{
			{PublicKey: exchange, IsSigner: false, IsWritable: true},
			{PublicKey: tokenSwap, IsSigner: false, IsWritable: false},
			{PublicKey: tokenSwapMarket, IsSigner: false, IsWritable: false},
			{PublicKey: tokenSwapAuthority, IsSigner: false, IsWritable: false},
			{PublicKey: userOwner, IsSigner: true, IsWritable: false},
			{PublicKey: tokenSol, IsSigner: false, IsWritable: true},
			{PublicKey: tokenSwapSrc, IsSigner: false, IsWritable: true},
			{PublicKey: tokenSwapDst, IsSigner: false, IsWritable: true},
			{PublicKey: tokenUSDT, IsSigner: false, IsWritable: true},
			{PublicKey: poolMint, IsSigner: false, IsWritable: true},
			{PublicKey: poolAccount, IsSigner: false, IsWritable: true},
			{PublicKey: splTokenProgram, IsSigner: false, IsWritable: false},
		}
		instruction.IsAccounts = append(instruction.IsAccounts, accounts...)
		ins = append(ins, instruction)
	}

	{
		//
		amountIn := uint64(0)
		data := make([]byte, 12)
		data[0] = 0
		binary.LittleEndian.PutUint64(data[1:], amountIn)
		data[9] = 1
		data[10] = 0
		data[11] = 2
		//
		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{
			},
			IsData:      data,
			IsProgramID: arbitrage,
		}
		// saber, usdt -> usdc
		saberSwap := solana.MustPublicKeyFromBase58("SSwpkEEcbUqx4vtoEByFjSkhKdCT862DNVb52nZg1UZ")
		saberMarket := solana.MustPublicKeyFromBase58("YAkoNb6HKmSxQN9L8hiBE5tPJRsniSSMzND1boHmZxe")
		saberAuthority, _, err := solana.FindProgramAddress([][]byte{saberMarket.Bytes()}, saberSwap)
		if err != nil {
			panic(err)
		}
		//userOwner := solana.MustPublicKeyFromBase58("sjcCd2jMG9BfmwwTDiitmJp8vSTJEF4M2wuBm3hR1M9")
		//userSrc := solana.MustPublicKeyFromBase58("92ev3jbJowvFpi6BaaXS3GkzjZhiwaFQ2nxUs8i61YKL")
		saberSwapSrc := solana.MustPublicKeyFromBase58("EnTrdMMpdhugeH6Ban6gYZWXughWxKtVGfCwFn78ZmY3")
		saberSwapDst := solana.MustPublicKeyFromBase58("CfWX7o2TswwbxusJ4hCaPobu2jLCb1hfXuXJQjVq3jQF")
		//userDst := solana.MustPublicKeyFromBase58("9aqNrRocaMkFE2qeq685JVcxpdHpcmeHF5pJGSs6d5or")
		feeAdmin := solana.MustPublicKeyFromBase58("GLztedC76MeBXjAmVXMezcHQzdmQaVLiXCZr9KEBSR6Y")
		//splTokenProgram := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
		//
		accounts := []*solana.AccountMeta{
			{PublicKey: exchange, IsSigner: false, IsWritable: true},
			{PublicKey: saberSwap, IsSigner: false, IsWritable: false},
			{PublicKey: saberMarket, IsSigner: false, IsWritable: false},
			{PublicKey: saberAuthority, IsSigner: false, IsWritable: false},
			{PublicKey: userOwner, IsSigner: true, IsWritable: false},
			{PublicKey: tokenUSDT, IsSigner: false, IsWritable: true},
			{PublicKey: saberSwapSrc, IsSigner: false, IsWritable: true},
			{PublicKey: saberSwapDst, IsSigner: false, IsWritable: true},
			{PublicKey: tokenUSDC, IsSigner: false, IsWritable: true},
			{PublicKey: feeAdmin, IsSigner: false, IsWritable: true},
			{PublicKey: splTokenProgram, IsSigner: false, IsWritable: false},
			{PublicKey: SysClock, IsSigner: false, IsWritable: false},
		}
		instruction.IsAccounts = append(instruction.IsAccounts, accounts...)
		ins = append(ins, instruction)
	}
	return ins
}

func TestArbitrageAll1(t *testing.T) {
	ins := ArbitrageAll1()
	SendTransaction(ins)
}


func TestSerumNewOrderOne(t *testing.T) {
	ins := make([]solana.Instruction, 0)
	//
	splTokenProgram := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	sysRent := solana.MustPublicKeyFromBase58("SysvarRent111111111111111111111111111111111")
	userOwner := solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS")
	tokenUSDC := solana.MustPublicKeyFromBase58("C5dr2fk9zLvjuSjxe3tdvFurcZHKbfU1mHbpdv9Bfr7T")
	tokenSol := solana.MustPublicKeyFromBase58("E2tAVqGR7XQwkjy9HJ7cPWu4SJN5oeLW1B7ZPJE5dWP3")
	tokenUSDT := solana.MustPublicKeyFromBase58("EDASG9z1jHYMxDqda9JeFM9E7UJAEg7j9LqJtqevwKL5")

	{
		// serum, usdc -> sol
		serum := solana.MustPublicKeyFromBase58("9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin")
		serumMarket := solana.MustPublicKeyFromBase58("9wFFyRfZBsuAha4YcuxcXLKwMxJR43S7fPfQLusDBzvT")
		openOrders := solana.MustPublicKeyFromBase58("CcANeZLGr8VbRco3v8hxyAbsgiG3fiNDNWDLZf5Ru1vY")
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
		nonce := make([]byte, 8)
		binary.LittleEndian.PutUint64(nonce, 1)
		vaultSigner, err := solana.CreateProgramAddress([][]byte{serumMarket.Bytes(), nonce}, serum)
		if err != nil {
			panic(err)
		}
		//
		amountIn := uint64(1000000000)
		data := make([]byte, 12)
		data[0] = 0
		binary.LittleEndian.PutUint64(data[1:], amountIn)
		data[9] = 2
		data[10] = 0 // buy
		data[11] = 0
		//
		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{
			},
			IsData:      data,
			IsProgramID: ARB,
		}
		//
		accounts := []*solana.AccountMeta{
			{PublicKey: ARB, IsSigner: false, IsWritable: true},
			{PublicKey: serum, IsSigner: false, IsWritable: false},
			{PublicKey: serumMarket, IsSigner: false, IsWritable: true},
			{PublicKey: openOrders, IsSigner: false, IsWritable: true},
			{PublicKey: userOwner, IsSigner: true, IsWritable: false},
			{PublicKey: request_quene, IsSigner: false, IsWritable: true},
			{PublicKey: event_queue, IsSigner: false, IsWritable: true},
			{PublicKey: bids, IsSigner: false, IsWritable: true},
			{PublicKey: asks, IsSigner: false, IsWritable: true},
			{PublicKey: tokenUSDC, IsSigner: false, IsWritable: true},
			{PublicKey: tokenSol, IsSigner: false, IsWritable: true},
			{PublicKey: base_vault, IsSigner: false, IsWritable: true},
			{PublicKey: quote_vault, IsSigner: false, IsWritable: true},
			{PublicKey: tokenSol, IsSigner: false, IsWritable: true},
			{PublicKey: tokenUSDC, IsSigner: false, IsWritable: true},
			{PublicKey: vaultSigner, IsSigner: false, IsWritable: false},
			{PublicKey: splTokenProgram, IsSigner: false, IsWritable: false},
			{PublicKey: sysRent, IsSigner: false, IsWritable: false},
		}
		instruction.IsAccounts = append(instruction.IsAccounts, accounts...)
		ins = append(ins, instruction)
	}

	{
		//
		amountIn := uint64(0)
		data := make([]byte, 12)
		data[0] = 0
		binary.LittleEndian.PutUint64(data[1:], amountIn)
		data[9] = 0
		data[10] = 0
		data[11] = 1
		//
		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{
			},
			IsData:      data,
			IsProgramID: ARB,
		}
		// orca, sol -> usdt
		tokenSwap := solana.MustPublicKeyFromBase58("9W959DqEETiGZocYWCQPaJ6sBmUzgfxXfqGeTEdp3aQP")
		tokenSwapMarket := solana.MustPublicKeyFromBase58("Dqk7mHQBx2ZWExmyrR2S8X6UG75CrbbpK2FSBZsNYsw6")
		tokenSwapAuthority, _, err := solana.FindProgramAddress([][]byte{tokenSwapMarket.Bytes()}, tokenSwap)
		if err != nil {
			panic(err)
		}
		//userOwner := solana.MustPublicKeyFromBase58("FrJZ4DP12Tg7r8rpjMqknkpCbJihqbEhfEBBQkpFimaS")
		//userSrc := solana.MustPublicKeyFromBase58("E2tAVqGR7XQwkjy9HJ7cPWu4SJN5oeLW1B7ZPJE5dWP3")
		tokenSwapSrc := solana.MustPublicKeyFromBase58("DTb8NKsfhEJGY1TrA7RXN6MBiTrjnkdMAfjPEjtmTT3M")
		tokenSwapDst := solana.MustPublicKeyFromBase58("E8erPjPEorykpPjFV9yUYMYigEWKQUxuGfL2rJKLJ3KU")
		//userDst := solana.MustPublicKeyFromBase58("EDASG9z1jHYMxDqda9JeFM9E7UJAEg7j9LqJtqevwKL5")
		poolMint := solana.MustPublicKeyFromBase58("FZthQCuYHhcfiDma7QrX7buDHwrZEd7vL8SjS6LQa3Tx")
		poolAccount := solana.MustPublicKeyFromBase58("BBKgw75FivTYXj85D2AWyVdaTdTWuSuHVXRm1Xu7fipb")
		//
		accounts := []*solana.AccountMeta{
			{PublicKey: ARB, IsSigner: false, IsWritable: true},
			{PublicKey: tokenSwap, IsSigner: false, IsWritable: false},
			{PublicKey: tokenSwapMarket, IsSigner: false, IsWritable: false},
			{PublicKey: tokenSwapAuthority, IsSigner: false, IsWritable: false},
			{PublicKey: userOwner, IsSigner: true, IsWritable: false},
			{PublicKey: tokenSol, IsSigner: false, IsWritable: true},
			{PublicKey: tokenSwapSrc, IsSigner: false, IsWritable: true},
			{PublicKey: tokenSwapDst, IsSigner: false, IsWritable: true},
			{PublicKey: tokenUSDT, IsSigner: false, IsWritable: true},
			{PublicKey: poolMint, IsSigner: false, IsWritable: true},
			{PublicKey: poolAccount, IsSigner: false, IsWritable: true},
			{PublicKey: splTokenProgram, IsSigner: false, IsWritable: false},
		}
		instruction.IsAccounts = append(instruction.IsAccounts, accounts...)
		ins = append(ins, instruction)
	}

	{
		//
		amountIn := uint64(0)
		data := make([]byte, 12)
		data[0] = 0
		binary.LittleEndian.PutUint64(data[1:], amountIn)
		data[9] = 1
		data[10] = 0
		data[11] = 2
		//
		instruction := &Instruction{
			IsAccounts: []*solana.AccountMeta{
			},
			IsData:      data,
			IsProgramID: ARB,
		}
		// saber, usdt -> usdc
		saberSwap := solana.MustPublicKeyFromBase58("SSwpkEEcbUqx4vtoEByFjSkhKdCT862DNVb52nZg1UZ")
		saberMarket := solana.MustPublicKeyFromBase58("YAkoNb6HKmSxQN9L8hiBE5tPJRsniSSMzND1boHmZxe")
		saberAuthority, _, err := solana.FindProgramAddress([][]byte{saberMarket.Bytes()}, saberSwap)
		if err != nil {
			panic(err)
		}
		//userOwner := solana.MustPublicKeyFromBase58("sjcCd2jMG9BfmwwTDiitmJp8vSTJEF4M2wuBm3hR1M9")
		//userSrc := solana.MustPublicKeyFromBase58("92ev3jbJowvFpi6BaaXS3GkzjZhiwaFQ2nxUs8i61YKL")
		saberSwapSrc := solana.MustPublicKeyFromBase58("EnTrdMMpdhugeH6Ban6gYZWXughWxKtVGfCwFn78ZmY3")
		saberSwapDst := solana.MustPublicKeyFromBase58("CfWX7o2TswwbxusJ4hCaPobu2jLCb1hfXuXJQjVq3jQF")
		//userDst := solana.MustPublicKeyFromBase58("9aqNrRocaMkFE2qeq685JVcxpdHpcmeHF5pJGSs6d5or")
		feeAdmin := solana.MustPublicKeyFromBase58("GLztedC76MeBXjAmVXMezcHQzdmQaVLiXCZr9KEBSR6Y")
		//splTokenProgram := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
		//
		accounts := []*solana.AccountMeta{
			{PublicKey: ARB, IsSigner: false, IsWritable: true},
			{PublicKey: saberSwap, IsSigner: false, IsWritable: false},
			{PublicKey: saberMarket, IsSigner: false, IsWritable: false},
			{PublicKey: saberAuthority, IsSigner: false, IsWritable: false},
			{PublicKey: userOwner, IsSigner: true, IsWritable: false},
			{PublicKey: tokenUSDT, IsSigner: false, IsWritable: true},
			{PublicKey: saberSwapSrc, IsSigner: false, IsWritable: true},
			{PublicKey: saberSwapDst, IsSigner: false, IsWritable: true},
			{PublicKey: tokenUSDC, IsSigner: false, IsWritable: true},
			{PublicKey: feeAdmin, IsSigner: false, IsWritable: true},
			{PublicKey: splTokenProgram, IsSigner: false, IsWritable: false},
			{PublicKey: SysClock, IsSigner: false, IsWritable: false},
		}
		instruction.IsAccounts = append(instruction.IsAccounts, accounts...)
		ins = append(ins, instruction)
	}
	return
}




