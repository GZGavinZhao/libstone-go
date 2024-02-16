package stone1

import (
	"io"

	"github.com/klauspost/compress/zstd"
)

// Reader iterates over the content of a V1 stone archive.
type Reader struct {
	Err error

	pre Prelude   // pre is the archive's prelude.
	src io.Reader // src is the reader from which the archive content is read.

	currHeader Header // currHeader is the header of the current payload.
	idxPayload int    // idxPayload points to the current payload.
	idxRecord  int    // idxRecord points to the current record.

	decomp *zstd.Decoder      // decomp decompresses payloads.
	cache  io.ReadWriteSeeker // cache is where the payloads are extracted.
}

// NewReader creates a new Reader which continues to read a stone archive from src.
// pre is the previously-written Prelude of the archive.
// Since stone payloads may be big in size, a cache is required to temporarily store data.
func NewReader(pre Prelude, src io.Reader, cache io.ReadWriteSeeker) *Reader {
	decomp, _ := zstd.NewReader(nil)
	return &Reader{
		pre:        pre,
		src:        src,
		idxPayload: -1,
		decomp:     decomp,
		cache:      cache,
	}
}

// NextPayload advances to the next payload Header.
// It returns true if it advanced to the next payload Header, false otherwise.
// If false was returned and r.Err is nil, it reached the end of the stone archive.
func (r *Reader) NextPayload() bool {
	if r.Err != nil {
		return false
	}
	if r.idxPayload >= int(r.pre.NumPayloads) {
		r.Err = nil
		return false
	}

	hdr, err := r.readHeader()
	if err != nil {
		r.Err = err
		return false
	}
	r.currHeader = hdr
	r.idxPayload += 1
	r.idxRecord = -1
	return true
}

// NextRecord advances to the next payload record.
// It returns a nil record and io.EOF if it reached the end of the current payload.
// It panics if NextPayload was not called beforehand.
func (r *Reader) NextRecord() (Record, error) {
	if r.Err != nil {
		return nil, r.Err
	}
	if r.idxPayload < 0 {
		panic("NextPayload was not called")
	}
	if r.idxRecord >= int(r.currHeader.NumRecords) {
		return nil, io.EOF
	}

	if r.idxRecord < 0 {
		err := r.extractPayload()
		if err != nil {
			r.Err = err
			return nil, err
		}
		r.idxRecord = 0
	}
	// TODO: Read.
	r.idxRecord += 1
	return nil, nil
}

func (r *Reader) readHeader() (Header, error) {
	var buf [headerLen]byte
	_, err := io.ReadFull(r.src, buf[:])
	if err != nil {
		return Header{}, err
	}
	return newHeader(buf), nil
}

func (r *Reader) extractPayload() error {
	return nil
}

func (r *Reader) readRecord() (Record, error) {
	return nil, nil
}
