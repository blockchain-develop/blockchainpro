package bitcoin

import (
	"fmt"
	"testing"
)

func TestGetBlock(t *testing.T) {
	//btc := NewBtcTools("http://172.168.3.10:20336", "test", "test")
	btc := NewBtcTools("http://18.140.187.37:18332", "omnicorerpc", "EzriglUqnFC!")
	height, err := btc.GetCurrentHeight()
	if err != nil {
		fmt.Printf("GetCurrentHeight err: %v\n", err)
		return
	}
	fmt.Printf("current height: %d\n", height)
}
