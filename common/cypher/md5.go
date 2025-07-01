package cypher

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func MD5(input []byte) [16]byte {
	return md5.Sum(input)
}

func MD5Hex(input []byte) string {
	sum := md5.Sum(input)
	return hex.EncodeToString(sum[:])
}

func MD5File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}