package utils

import "math/rand"

const (
	low  = 1000
	high = 9999
)

func GenerateVerifCode() int {
	return low + rand.Intn(high-low)
}
