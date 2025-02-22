package utils

import "github.com/google/uuid"

func GenerateBidID() string {
	return uuid.New().String()
}
