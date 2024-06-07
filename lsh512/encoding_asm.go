//go:build (arm64 || amd64 || amd64p32) && !purego
// +build arm64 amd64 amd64p32
// +build !purego

package lsh512

import (
	"encoding"
	"encoding/binary"
	"errors"
)

const (
	magic         = "lsh\x01asm"
	marshaledSize = len(magic) +
		/*data*/ contextDataSize +
		/*size*/ 2
)

func (ctx *lsh512ContextAsm) MarshalBinary() ([]byte, error) {
	b := make([]byte, 0, marshaledSize)
	b = append(b, magic...)
	b = append(b, ctx.data[:]...)
	b = binary.BigEndian.AppendUint16(b, uint16(ctx.size))
	return b, nil
}

func (ctx *lsh512ContextAsm) UnmarshalBinary(b []byte) error {
	if len(b) < len(magic) || string(b[:len(magic)]) != magic {
		return errors.New("krypto/hash160: invalid hash state identifier")
	}
	if len(b) != marshaledSize {
		return errors.New("krypto/hash160: invalid hash state size")
	}

	b = b[len(magic):]
	b = b[copy(ctx.data[:], b[:contextDataSize]):]
	ctx.size = int(binary.BigEndian.Uint16(b))
	return nil
}

var (
	_ encoding.BinaryMarshaler   = (*lsh512ContextAsm)(nil)
	_ encoding.BinaryUnmarshaler = (*lsh512ContextAsm)(nil)
)
