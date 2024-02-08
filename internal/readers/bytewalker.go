package readers

import (
	"encoding/binary"
)

var (
	byteOrder = binary.BigEndian
)

type ByteWalker []byte

func (r *ByteWalker) Ahead(n int) []byte {
	val := (*r)[:n]
	*r = (*r)[n:]
	return val
}

func (r *ByteWalker) Uint8() uint8 {
	val := (*r)[0]
	*r = (*r)[1:]
	return val
}

func (r *ByteWalker) Uint16() uint16 {
	val := byteOrder.Uint16(*r)
	*r = (*r)[2:]
	return val
}

func (r *ByteWalker) Uint32() uint32 {
	val := byteOrder.Uint32(*r)
	*r = (*r)[4:]
	return val
}
