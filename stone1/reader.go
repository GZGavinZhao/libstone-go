package stone1

import "io"

// Reader iterates over the content of a V1 stone archive.
type Reader struct {
	pre Prelude
	src io.Reader

	idxPayload int
}

// NewReader creates a new Reader.
func NewReader(pre Prelude, src io.Reader) *Reader {
	return &Reader{
		pre: pre,
		src: src,
	}
}

// Next advances to the next Header. It returns an empty Header
// and io.EOF if it reached the end of stone archive.
func (r *Reader) Next() (Header, error) {
	if r.idxPayload >= int(r.pre.NumPayloads) {
		return Header{}, io.EOF
	}
	hdr, err := r.readHeader()
	if err != nil {
		return Header{}, err
	}
	r.idxPayload += 1
	return hdr, nil
}

func (r *Reader) readHeader() (Header, error) {
	var buf [headerLen]byte
	_, err := io.ReadFull(r.src, buf[:])
	if err != nil {
		return Header{}, err
	}
	return newHeader(buf), nil
}
