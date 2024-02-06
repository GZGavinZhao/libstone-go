package libstone

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

// Version is the stone format version contained inside the [Prelude].
type Version uint32

const (
	// V1 is the first version of the stone format.
	V1 Version = iota + 1
)

type magicNumber = [4]byte

var (
	// MagicNumber is the magic number of a stone archive.
	MagicNumber = magicNumber{0, 'm', 'o', 's'}
)

var (
	// ErrNoStone is returned when the magic number doesn't match
	// [MagicNumber].
	ErrNoStone = errors.New("data is not a stone archive")
)

// PreludeData is an agnostic array of bytes extending the base Prelude.
// Its meaning varies according to Version.
type PreludeData [24]byte

// Prelude is the header of the stone format.
type Prelude struct {
	Data PreludeData

	// Version is the version of this stone archive.
	Version Version
}

func (h *Prelude) UnmarshalBinary(data []byte) error {
	r := bytes.NewReader(data)

	var magic magicNumber
	_, err := io.ReadFull(r, magic[:])
	if err != nil {
		return err
	}
	if magic != MagicNumber {
		return ErrNoStone
	}

	return binary.Read(r, binary.BigEndian, &h)
}
