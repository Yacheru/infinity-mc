package utils

import (
	"crypto/sha512"
	"fmt"
)

type Hasher interface {
	Hash(string) string
}

type SHA512Hasher struct {
	salt string
}

func NewSHA512Hasher(salt string) *SHA512Hasher {
	return &SHA512Hasher{salt: salt}
}

func (sh *SHA512Hasher) Hash(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(sh.salt)))
}
