package common

import (
	"crypto/sha256"
	"encoding/hex"
)

type Sha256Hash struct {
}

func NewSha256Hash() *Sha256Hash {
	return &Sha256Hash{}
}
func (h *Sha256Hash) Hash(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
