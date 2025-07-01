package cypher

import (
	"crypto/sha1"
	"encoding/hex"
)

func SHA1(input []byte) string {
	sum := sha1.Sum(input)
	return hex.EncodeToString(sum[:])
}