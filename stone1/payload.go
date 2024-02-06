package stone1

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// ArchiveType is the kind of the archive content.
type ArchiveType = uint8

const (
	// MetaArchive indicates a metadata store.
	MetaArchive ArchiveType = iota + 1
	// ContentArchive indicates a file store, i.e. hash indexed.
	ContentArchive
	// LayoutArchive indicates a file-to-disk map with basic UNIX permissions and types.
	LayoutArchive
	// Index indicates an index for deduplicated store.
	IndexArchive
	// AttributeArchive indicates an attribute store.
	AttributeArchive
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
	// Kind is the kind of payload.
	Kind ArchiveType
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
