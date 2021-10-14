package random

import (
	"crypto/rand"
	"math/big"
	rrand "math/rand"
)

func RandomIntSafe(min, max int) int {
	if min > max {
		panic("min > max")
	}
	num, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		panic(err)
	}
	return int(num.Int64()) + min
}

func RandomInt64Safe(min, max int64) int64 {
	if min > max {
		panic("min > max")
	}
	num, err := rand.Int(rand.Reader, big.NewInt(max-min))
	if err != nil {
		panic(err)
	}
	return num.Int64() + min
}

func RandomInt(min, max int) int {
	if min > max {
		panic("min > max")
	}
	return rrand.Intn(max-min) + min
}

func RandomInt64(min, max int64) int64 {
	if min > max {
		panic("min > max")
	}
	return rrand.Int63n((max - min)) + min
}
