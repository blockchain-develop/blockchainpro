package solana

import (
	"encoding/base64"
	//"encoding/base64"
	"encoding/binary"
	"fmt"
	"github.com/mr-tron/base58"
	"testing"
)

func parseData1(data []byte) {
	version := data[0]
	command := binary.LittleEndian.Uint32(data[1:])
	side := binary.LittleEndian.Uint32(data[5:])
	price := binary.LittleEndian.Uint64(data[9:])
	base := binary.LittleEndian.Uint64(data[17:])
	quote := binary.LittleEndian.Uint64(data[25:])

	fmt.Printf("%d, %d, %d, %d, %d, %d\n", version, command, side, price, base, quote)
}

func TestInstruction_DataDecode(t *testing.T) {
	contentString := "189VEfQCdeLgv9KiurKGEPcPUeDLAurQmhCYQNGMMJBFzFZwGEoEREskDvb6zubEMj2eA"
	content, err :=  base58.Decode(contentString)
	if err != nil {
		panic(err)
	}
	parseData1(content)
}

func TestDataDecode(t *testing.T) {
	contentString := "ZKMujQAAAAA="
	content, err :=  base64.StdEncoding.DecodeString(contentString)
	if err != nil {
		panic(err)
	}
	value := binary.LittleEndian.Uint64(content)
	//value := new(big.Int).SetBytes(content)
	fmt.Printf("price: %d\n",value)
}
