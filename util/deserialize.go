package util

import "encoding/binary"

func DeserialzeInt64(b []byte) int64 {

	return int64(binary.LittleEndian.Uint64(b))
}
