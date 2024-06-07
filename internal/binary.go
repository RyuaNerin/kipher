package internal

import "encoding/binary"

func ConsumeUint64(b []byte) ([]byte, uint64) {
	return b[8:], binary.BigEndian.Uint64(b)
}

func ConsumeUint32(b []byte) ([]byte, uint32) {
	return b[4:], binary.BigEndian.Uint32(b)
}
