package random

import (
	"crypto/rand"
	"math/big"
)

func Int64(nMax int64) int64 {
	r, err := rand.Int(rand.Reader, big.NewInt(nMax))
	if err != nil {
		panic(err)
	}

	return r.Int64()
}
