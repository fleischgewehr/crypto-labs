package lib

import "math/big"

func IntToBytes(i int) []byte {
	if i > 0 {
		return append(big.NewInt(int64(i)).Bytes(), byte(1))
	}
	return append(big.NewInt(int64(i)).Bytes(), byte(0))
}

func BytesToInt(b []byte) int {
	if b[len(b)-1] == 0 {
		return -int(big.NewInt(0).SetBytes(b[:len(b)-1]).Int64())
	}
	return int(big.NewInt(0).SetBytes(b[:len(b)-1]).Int64())
}
