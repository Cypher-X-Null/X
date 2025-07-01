package cypher

import (
	"hash/crc32"
)

var crc32cTable = crc32.MakeTable(crc32.Castagnoli)

func CRC32C(data []byte) uint32 {
	return crc32.Checksum(data, crc32cTable)
}