package bybit

import (
	"encoding/json"
	"fmt"
	"github.com/hirokisan/bybit"
	"testing"
)

func TestBalance(t *testing.T) {
	client := bybit.NewTestClient().WithAuth("lZFv30nrQnoagJY96k", "bPy3QsHGOcn2xHrEtCg1ivgeU4qXBUiEJfxn")
	res, err := client.Wallet().Balance(bybit.CoinUSDT)
	if err != nil {
		panic(err)
	}
	balance := res.Result.Balance[bybit.CoinUSDT]
	balanceJson, _ := json.MarshalIndent(balance, "", "    ")
	fmt.Printf("balance: %s\n", balanceJson)
}
