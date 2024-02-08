package stone1

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// RecordKind is the kind of the payload's records.
type RecordKind = uint8

const (
	// Meta indicates a [MetaRecord].
	Meta RecordKind = iota + 1
	// Content indicates a file.
	Content
	// Layout indicates a [LayoutRecord].
	Layout
	// Index indicates an [IndexRecord].
	Index
	// Attributes indicates an attribute store.
	Attributes
)

// Compression is the compression method of the archive content.
type Compression uint8

const (
	// Uncompressed indicates an uncompressed content.
	Uncompressed Compression = iota + 1
	// Uncompressed indicates a compressed content using zstd.
	ZSTD
)

// Header is the payload's header.
type Header struct {
	// StoredSize is the size, in bytes, of the content once stored in mass memory.
	StoredSize uint64
	// PlainSize is the size, in bytes, of the content as it is.
	PlainSize uint64
	// Checksum is the payload's checksum.
	Checksum [8]byte
	// NumRecords is the number of records contained in the payload.
	NumRecords uint32
	// Version is the version of the payload data format.
	Version uint16
	// Kind is the kind of payload's records.
	Kind RecordKind
	// Compression is the compression used for the payload.
	Compression Compression
}

func (h *Header) UnmarshalBinary(data []byte) error {
	if len(data) < binary.Size(h) {
		return errors.New("insufficient number of bytes to parse a V1 payload header")
	}
	rdr := bytes.NewReader(data)
	return binary.Read(rdr, binary.BigEndian, h)
}

func (h *Header) MarshalBinary() ([]byte, error) {
	out := make([]byte, binary.Size(h))
	wrt := bytes.NewBuffer(out)
	return out, binary.Write(wrt, binary.BigEndian, h)
}
