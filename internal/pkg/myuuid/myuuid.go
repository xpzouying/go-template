package myuuid

import (
	"crypto/rand"
	"encoding/hex"
)

func New() string {
	b := make([]byte, 4)
	rand.Read(b) // Doesn’t actually fail
	s := hex.EncodeToString(b)

	return s
}
