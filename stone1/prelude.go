package stone1

import (
	"encoding/binary"
	"errors"

	"github.com/der-eismann/libstone"
)

var (
	integrityCheck = [21]byte{0, 0, 1, 0, 0, 2, 0, 0, 3, 0, 0, 4, 0, 0, 5, 0, 0, 6, 0, 0, 7}
)

type StoneType uint8

const (
	BinaryStone StoneType = iota + 1
	DeltaStone
	RepositoryStone
	BuildManifestStone
)

type Prelude struct {
	NumPayloads uint16
	StoneType   StoneType
}

func NewPrelude(genericPre libstone.Prelude) (Prelude, error) {
	if genericPre.Version != libstone.V1 {
		return Prelude{}, errors.New("prelude version is not 1")
	}
	var prelude Prelude
	return prelude, prelude.UnmarshalBinary(genericPre.Data[:])
}

func (p *Prelude) UnmarshalBinary(data []byte) error {
	if len(data) <= len(libstone.PreludeData{}) {
		return errors.New("insufficient number of bytes to parse a V1 prelude")
	}
	if [21]byte(data[2:2+len(integrityCheck)]) != integrityCheck {
		return errors.New("V1 integrity check failed")
	}

	p.NumPayloads = binary.BigEndian.Uint16(data)
	p.StoneType = StoneType(data[2+len(integrityCheck)])
	return nil
}

func (h *Prelude) MarshalBinary() ([]byte, error) {
	var out libstone.PreludeData
	binary.BigEndian.PutUint16(out[:], h.NumPayloads)
	copy(out[2:], integrityCheck[:])
	out[2+len(integrityCheck)] = byte(h.StoneType)
	return out[:], nil
}
