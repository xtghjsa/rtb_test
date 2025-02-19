package utils

import (
	"math/rand"
)

func RandomPrice() int {
	return rand.Intn(100)
}
