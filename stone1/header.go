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

type Header struct {
	NumPayloads uint16
	FileType    FileType
}

func NewHeader(genericHdr libstone.Header) (Header, error) {
	if genericHdr.Version != libstone.V1 {
		return Header{}, errors.New("header version is not 1")
	}
	var hdr Header
	return hdr, hdr.UnmarshalBinary(genericHdr.Data[:])
}

func (h *Header) UnmarshalBinary(data []byte) error {
	if len(data) <= len(libstone.HeaderData{}) {
		return errors.New("insufficient number of bytes to parse a V1 header")
	}
	if [21]byte(data[2:2+len(integrityCheck)]) != integrityCheck {
		return errors.New("V1 integrity check failed")
	}

	h.NumPayloads = binary.BigEndian.Uint16(data)
	h.FileType = FileType(data[2+len(integrityCheck)])
	return nil
}

func (h *Header) MarshalBinary() ([]byte, error) {
	var out libstone.HeaderData
	binary.BigEndian.PutUint16(out[:], h.NumPayloads)
	copy(out[2:], integrityCheck[:])
	out[2+len(integrityCheck)] = byte(h.FileType)
	return out[:], nil
}
