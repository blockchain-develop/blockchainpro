package bybit

import (
	"encoding/json"
	"fmt"
	"github.com/hirokisan/bybit"
	"testing"
)

func TestPlaceOrder(t *testing.T) {
	client := bybit.NewTestClient().WithAuth("lZFv30nrQnoagJY96k", "bPy3QsHGOcn2xHrEtCg1ivgeU4qXBUiEJfxn")
	{
		orderParam := bybit.CreateLinearOrderParam{}
		orderParam.OrderType = bybit.OrderTypeLimit
		orderParam.Symbol = bybit.SymbolUSDTETH
		price := float64(1500)
		orderParam.Price = &price
		qty := float64(1)
		orderParam.Qty = qty
		orderParam.Side = bybit.SideBuy
		//
		createdOrder, err := client.Account().CreateLinearOrder(orderParam)
		if err != nil {
			panic(err)
		}
		createdOrderJson, _ := json.MarshalIndent(createdOrder.Result, "", "    ")
		fmt.Printf("created order: %s\n", createdOrderJson)
	}
	{
		position, err := client.Account().ListLinearPosition(bybit.SymbolUSDTETH)
		if err != nil {
			panic(err)
		}
		positionJson, _ := json.MarshalIndent(position.Result, "", "    ")
		fmt.Printf("position: %s\n", positionJson)
	}
}
