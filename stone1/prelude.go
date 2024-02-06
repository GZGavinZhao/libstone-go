package stone1

import (
	"encoding/binary"
	"errors"

	"github.com/der-eismann/libstone"
)

var (
	integrityCheck = [21]byte{0, 0, 1, 0, 0, 2, 0, 0, 3, 0, 0, 4, 0, 0, 5, 0, 0, 6, 0, 0, 7}
)

type FileType uint8

const (
	Binary FileType = iota + 1
	Delta
	Repository
	BuildManifest
)

type Prelude struct {
	NumPayloads uint16
	FileType    FileType
}

func NewPrelude(genericPre libstone.Prelude) (Prelude, error) {
	if genericPre.Version != libstone.V1 {
		return Prelude{}, errors.New("header version is not 1")
	}
	var prelude Prelude
	return prelude, prelude.UnmarshalBinary(genericPre.Data[:])
}

func (p *Prelude) UnmarshalBinary(data []byte) error {
	if len(data) <= len(libstone.PreludeData{}) {
		return errors.New("insufficient number of bytes to parse a V1 header")
	}
	if [21]byte(data[2:2+len(integrityCheck)]) != integrityCheck {
		return errors.New("V1 integrity check failed")
	}

	p.NumPayloads = binary.BigEndian.Uint16(data)
	p.FileType = FileType(data[2+len(integrityCheck)])
	return nil
}

func (h *Prelude) MarshalBinary() ([]byte, error) {
	var out libstone.PreludeData
	binary.BigEndian.PutUint16(out[:], h.NumPayloads)
	copy(out[2:], integrityCheck[:])
	out[2+len(integrityCheck)] = byte(h.FileType)
	return out[:], nil
}
