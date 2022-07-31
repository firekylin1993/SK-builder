package kits

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"
)

func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

func IntToBytes(n uint32) []byte {
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, n) //nolint:errcheck
	return bytebuf.Bytes()
}

func GetCRC32Key(strKey []byte) uint32 {
	table := crc32.MakeTable(crc32.IEEE)
	return crc32.Checksum(strKey, table)
}
