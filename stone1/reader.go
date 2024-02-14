package stone1

import "io"

// Reader iterates over the content of a V1 stone archive.
type Reader struct {
	pre Prelude
	src io.Reader

	currHeader Header
	idxPayload int
	idxRecord  int
}

// NewReader creates a new Reader.
func NewReader(pre Prelude, src io.Reader) *Reader {
	return &Reader{
		pre:        pre,
		src:        src,
		idxPayload: -1,
		idxRecord:  -1,
	}
}

// NextPayload advances to the next payload Header.
// It returns an empty Header and io.EOF if it reached the end of stone archive.
func (r *Reader) NextPayload() (Header, error) {
	if r.idxPayload >= int(r.pre.NumPayloads) {
		return Header{}, io.EOF
	}
	hdr, err := r.readHeader()
	if err != nil {
		return Header{}, err
	}
	r.currHeader = hdr
	r.idxPayload += 1
	r.idxRecord = 0
	return hdr, nil
}

// NextRecord advances to the next payload record.
// It returns a nil record and io.EOF if it reached the end of the current payload.
// It panics if NextPayload was not called beforehand.
func (r *Reader) NextRecord() (any, error) {
	if r.idxPayload < 0 {
		panic("NextPayload was not called")
	}
	if r.idxRecord >= int(r.currHeader.NumRecords) {
		return nil, io.EOF
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
