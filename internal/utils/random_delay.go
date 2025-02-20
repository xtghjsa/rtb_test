package utils

import "math/rand"

func RandomDelay() int {
	return rand.Intn(500)
}
