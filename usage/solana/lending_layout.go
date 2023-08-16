package solana

import (
	"github.com/gagliardetto/solana-go"
)

var (
	LendingMarketLen = uint64(258)
	ReserveLen = uint64(571)
)

type LendingMarketLayout struct {
	Version uint8
	BumpSeed uint8
	Owner solana.PublicKey
	QuoteCurrency solana.PublicKey
	TokenProgram solana.PublicKey
	OracleProgram solana.PublicKey
	_ [4]solana.PublicKey
}

type LastUpdateLayout struct {
	Slot uint64
	Stale bool
}

type ReserveLiquidity struct {
	MintKey solana.PublicKey
	MintDecimal uint8
	SupplyKey solana.PublicKey
	FeeReceiver solana.PublicKey
	OracleKey solana.PublicKey
	AvailableAmount uint64
	BorrowedAmount Decimal
	CumulativeBorrowRate Decimal
	MarketPrice Decimal
}

type Decimal struct {
	Data [2]uint64
}

type ReserveCollateral struct {
	MintKey solana.PublicKey
	MintTotalSupply uint64
	SupplyKey solana.PublicKey
}

type ReserveConfig struct {
	Rate uint8
	Ratio uint8
	Bonus uint8
	Threshold uint8
	BorrowRate uint8
	BorrowRate1 uint8
	BorrowRate2 uint8
	Fee ReserveFees
}

type ReserveFees struct {
	BorrowFee uint64
	FlashLoan uint64
	Fee uint8
}

type ReserveLayout struct {
	Version uint8
	LastUpdate LastUpdateLayout
	LendingMarket solana.PublicKey
	Liquidity ReserveLiquidity
	Collateral ReserveCollateral
	Config ReserveConfig
	_ [30]uint64
}
