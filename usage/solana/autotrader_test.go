package solana

import (
	"context"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"io/ioutil"
	"os"
	"testing"
)

type AccountHolder struct {
	Mint                 solana.PublicKey
	Owner                solana.PublicKey
	Amount               uint64
	DelegateOption       [4]byte
	Delegate             solana.PublicKey
	State                byte
	IsNativeOption       [4]byte
	IsNative             uint64
	DelegatedAmount      uint64
	CloseAuthorityOption [4]byte
	CloseAuthority       solana.PublicKey
}

type Trader struct {
	ProgramID solana.PublicKey
	SwapInfo solana.PublicKey
	TokenA solana.PublicKey
	MintA solana.PublicKey
	TokenB solana.PublicKey
	MintB solana.PublicKey
	Pool solana.PublicKey
	Fee solana.PublicKey
}


type ConstantProductCurve struct {
	TokenAPubkey solana.PublicKey
	TokenBPubkey solana.PublicKey
	TokenAHolder solana.PublicKey
	TokenAAmount uint64
	TokenBHolder solana.PublicKey
	TokenBAmount uint64
	Pool solana.PublicKey
	Fee solana.PublicKey
	//
	TradeFeeNumerator        uint64
	TradeFeeDenominator      uint64
	OwnerTradeFeeNumerator   uint64
	OwnerTradeFeeDenominator uint64
}

type Pair struct {
	AmmId     solana.PublicKey
	PairId    solana.PublicKey
	CurveType byte
	Curve     *ConstantProductCurve
}


func readFile(fileName string) ([]byte, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("OpenFile %s error %s", fileName, err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("File %s close error %s", fileName, err)
		}
	}()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll %s error %s", fileName, err)
	}
	return data, nil
}

func ReadTraders() []*Pair {
	traders := make([]*Pair, 0)
	fileName := "./autotrader.json"
	data, err := readFile(fileName)
	if err != nil {
		fmt.Printf("read file err: %v\n", err)
	}
	err = json.Unmarshal(data, &traders)
	if err != nil {
		fmt.Printf("json.Unmarshal TestConfig:%s error:%s", data, err)
	}
	return traders
}

type UserAccountInfo struct {
	Holder solana.PublicKey
	Authority solana.PublicKey
}

var UserAccountInfos = map[solana.PublicKey]*UserAccountInfo {
	solana.MustPublicKeyFromBase58("So11111111111111111111111111111111111111112"): {
		solana.MustPublicKeyFromBase58("36c6YqAwyGKQG66XEp2dJc5JqjaBNv7sVghEtJv4c7u6"),
		solana.MustPublicKeyFromBase58("F8Vyqk3unwxkXukZFQeYyGmFfTG3CAX4v24iyrjEYBJV"),
	},
	solana.MustPublicKeyFromBase58("BQcdHdAQW1hczDbBi9hiegXAR7A98Q9jx3X3iBBBDiq4"): {
		solana.MustPublicKeyFromBase58("648kz2PXcnUwDWHqhu1k38gFLbnRMV88J9ByLtcesE2a"),
		solana.MustPublicKeyFromBase58("DkJ76rxxV4HnCm4TZAtPTykBjuthXZHjsk9NHyAUmUCT"),
	},
	solana.MustPublicKeyFromBase58("EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v"): {
		solana.MustPublicKeyFromBase58("btetJ8RQZ448nKtc7mNWGvZQAUXTJpWuZDo5h9qLr3Z"),
		solana.MustPublicKeyFromBase58("HJEXUEyvgYuwm8rBoLPc7ZdVty8aPqo7MUw9cK4v6afi"),
	},
	solana.MustPublicKeyFromBase58("4k3Dyjzvzp8eMZWUXbBCjEvwSkkk59S5iCNLY3QrkX6R"): {
		solana.MustPublicKeyFromBase58("HoVhs7NAzcAaavr2tc2aaTynJ6kwhdfC2B2Z7EthKpeo"),
		solana.MustPublicKeyFromBase58("8pFhUqCU7Fkxfg2DLytRDf7a9oK4XGtN92PrYwtVQc6G"),
	},
	solana.MustPublicKeyFromBase58("SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt"): {
		solana.MustPublicKeyFromBase58("2qryU73SHosH1vicYDfHUvKxAvNK7t6ivB8qowBSgG1z"),
		solana.MustPublicKeyFromBase58("3yFwqXBfZY4jBVUafQ1YEXw189y2dN3V5KQq9uzBDy1E"),
	},
	solana.MustPublicKeyFromBase58("9n4nbM75f5Ui33ZbPYXn59EwSgE8CGsHtAeTH5YFeJ9E"): {
		solana.MustPublicKeyFromBase58("DSf7hGudcxhhegMpZA1UtSiW4RqKgyEex9mqQECWwRgZ"),
		solana.MustPublicKeyFromBase58("EPzuCsSzHwhYWn2j69HQPKWuWz6wuv4ANZiVigLGMBoD"),
	},
	solana.MustPublicKeyFromBase58("2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk"): {
		solana.MustPublicKeyFromBase58("7Nw66LmJB6YzHsgEGQ8oDSSsJ4YzUkEVAvysQuQw7tC4"),
		solana.MustPublicKeyFromBase58("C5v68qSzDdGeRcs556YoEMJNsp8JiYEiEhw2hVUR8Z8y"),
	},
}

func FindUserAccountInfo(token solana.PublicKey) *UserAccountInfo {
	for k, userAccountInfo := range UserAccountInfos {
		if k == token {
			return userAccountInfo
		}
	}
	return nil
}

func TestAutoTrader(t *testing.T) {
	//orcaProgramID := solana.MustPublicKeyFromBase58("DjVE6JNiYqPL2QXyCUUh8rNjHrbz9hXHNYt99MQ59qw1")
	//tokenswapProgramID := solana.MustPublicKeyFromBase58("SwaPpA9LAaLfeLi3a68M4DjnLqgtticKg6CnyNwgAC8")
	tokenProgramID := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	feePlayer := solana.MustPublicKeyFromBase58("Y2akr3bXHRsqyP1QJtbm9G9N88ZV4t1KfaFeDzKRTfr")

	client := rpc.New(rpc.MainNetBetaSerum_RPC)
	ctx := context.Background()
	getRecentBlockHashResult, err := client.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}
	blockHash := getRecentBlockHashResult.Value.Blockhash

	amount := uint64(1000000000)
	token := solana.MustPublicKeyFromBase58("So11111111111111111111111111111111111111112")
	traders := ReadTraders()
	for _, trader := range traders {
		program := trader.AmmId
		swap := trader.PairId
		swap_authority, _, _ := solana.FindProgramAddress([][]byte{swap.Bytes()}, program)
		sourceToken := trader.Curve.TokenAPubkey
		sourceHolder := trader.Curve.TokenAHolder
		destinationToken := trader.Curve.TokenBPubkey
		destinationHolder := trader.Curve.TokenBHolder
		if destinationToken == token {
			sourceToken = trader.Curve.TokenBPubkey
			sourceHolder = trader.Curve.TokenBHolder
			destinationToken = trader.Curve.TokenAPubkey
			destinationHolder = trader.Curve.TokenAHolder
		}
		userAccountInfo := FindUserAccountInfo(sourceToken)
		if userAccountInfo == nil {
			panic(fmt.Sprintf("token a user account is not exist, token: %s", sourceToken))
		}
		userAccountInfo2 := FindUserAccountInfo(destinationToken)
		if userAccountInfo2 == nil {
			panic(fmt.Sprintf("token a user account is not exist, token: %s", destinationToken))
		}
		swap_user_authority := userAccountInfo.Authority
		swap_user_account := userAccountInfo.Holder
		swap_tokenA := sourceHolder
		swap_tokenB := destinationHolder
		swap_user2_account := userAccountInfo2.Holder
		swap_pool := trader.Curve.Pool
		swap_fee_account := trader.Curve.Fee

		var accountBefore AccountHolder
		err = client.GetAccountDataInto(ctx, swap_user2_account, &accountBefore)
		if err != nil {
			panic(err)
		}
		//
		data := make([]byte, 17)
		data[0] = 1
		binary.LittleEndian.PutUint64(data[1:], amount)
		binary.LittleEndian.PutUint64(data[9:], 0)
		fmt.Printf("swap in (%s, %d)\n", sourceToken.String(), amount)

		builder := solana.NewTransactionBuilder()
		builder.AddInstruction(&orcaTransactionInstructions{
			accounts: []*solana.AccountMeta{
				{PublicKey: swap, IsSigner: false, IsWritable: false},
				{PublicKey: swap_authority, IsSigner: true, IsWritable: false},
				{PublicKey: swap_user_authority, IsSigner: true, IsWritable: false},
				{PublicKey: swap_user_account, IsSigner: false, IsWritable: true},
				{PublicKey: swap_tokenA, IsSigner: false, IsWritable: true},
				{PublicKey: swap_tokenB, IsSigner: false, IsWritable: true},
				{PublicKey: swap_user2_account, IsSigner: false, IsWritable: true},
				{PublicKey: swap_pool, IsSigner: false, IsWritable: true},
				{PublicKey: swap_fee_account, IsSigner: false, IsWritable: true},
				{PublicKey: tokenProgramID, IsSigner: false, IsWritable: false},
				{PublicKey: swap_fee_account, IsSigner: false, IsWritable: false},
			},
			data:      data,
			programID: program,
		})

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
				Addresses: []solana.PublicKey{swap_user2_account},
			},
		})
		if err != nil {
			panic(err)
		}
		{
			rJson, _ := json.MarshalIndent(simulateTransactionResponse, "", "    ")
			fmt.Printf("SimulateTransactionResponse: %s\n", string(rJson))
		}

		/*
		var accountAfter AccountHolder
		err = bin.NewBorshDecoder(simulateTransactionResponse.Value.Accounts[0].Data.GetBinary()).Decode(&accountAfter)
		if err != nil {
			panic(err)
		}
		amount = accountAfter.Amount - accountBefore.Amount
		token = destinationToken
		fmt.Printf("swap out (%s, %d)\n", destinationToken.String(), amount)
		 */
	}
}
