package cypher

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(input []byte) string {
	sum := sha256.Sum256(input)
	return hex.EncodeToString(sum[:])
}