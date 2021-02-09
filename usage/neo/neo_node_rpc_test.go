package neo

import (
	"fmt"
	"testing"
)

func TestGetApplicationLog(t *testing.T) {
	client := NewNeoClient()
	res := client.GetApplicationLog("063c9a998a6bfa4054ad82ebd3799c9f8ae941b28183b6b1b75b5cfbe8af891e")
	fmt.Printf("%v\n", res)
}

func TestGetLatestHeight(t *testing.T) {
	client := NewNeoClient()
	countrep := client.GetBlockCount()
	fmt.Printf("height: %d\n", countrep.Result)
}

