package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func Bytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Reader.Read(b)
	if err != nil {
		panic(fmt.Sprintf("failed to generate rand bytes: %v", err))
	}
	return b
}

func Base64(n int) string {
	return base64.StdEncoding.EncodeToString(Bytes(n))
}
