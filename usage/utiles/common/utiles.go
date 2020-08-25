package common

import (
	"encoding/hex"
)

func RevertBytes(data []byte) []byte {
	n := len(data)
	newdata := make([]byte, n)
	for i := 0;i < n;i ++ {
		newdata[i] = data[n - 1 - i]
	}
	return newdata
}

func HexStringReverse(value string) string {
	aa, _ := hex.DecodeString(value)
	bb := RevertBytes(aa)
	return hex.EncodeToString(bb)
}
