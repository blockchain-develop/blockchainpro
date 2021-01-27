package neo

import (
	"fmt"
	"testing"
)

func TestGetApplicationLog(t *testing.T) {
}

func TestGetLatestHeight(t *testing.T) {
	client := NewNeoClient()
	countrep := client.GetBlockCount()
	fmt.Printf("height: %d\n", countrep.Result)
}
