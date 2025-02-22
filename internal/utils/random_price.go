package utils

import (
	"math/rand"
)

func RandomPrice() int64 {
	return int64(rand.Intn(100))
}
