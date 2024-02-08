package readers_test

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/der-eismann/libstone/internal/readers"
)

const (
	aheadDistance = 5
)

var (
	testData = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
)

func TestAhead(t *testing.T) {
	wlk := readers.ByteWalker(testData)
	testAhead(t, &wlk, testData)
}

func TestUint8(t *testing.T) {
	wlk := readers.ByteWalker(testData)
	testUint8(t, &wlk, testData)
}

func TestUint16(t *testing.T) {
	wlk := readers.ByteWalker(testData)
	testUint16(t, &wlk, testData)
}

func TestUint32(t *testing.T) {
	wlk := readers.ByteWalker(testData)
	testUint32(t, &wlk, testData)
}

func TestIsWalking(t *testing.T) {
	wlk := readers.ByteWalker(testData)
	data := testData
	testAhead(t, &wlk, data)
	testUint8(t, &wlk, data[aheadDistance:])
	testUint16(t, &wlk, data[aheadDistance+1:])
	testUint32(t, &wlk, data[aheadDistance+1+2:])
}

func testAhead(t *testing.T, wlk *readers.ByteWalker, data []byte) {
	expect := data[:aheadDistance]
	obtain := wlk.Ahead(aheadDistance)
	if !bytes.Equal(obtain, expect) {
		t.Fatalf("expected ahead slice %v. Got %v", expect, obtain)
	}
}

func testUint8(t *testing.T, wlk *readers.ByteWalker, data []byte) {
	expect := data[0]
	obtain := wlk.Uint8()
	if obtain != expect {
		t.Fatalf("expected uint8 %d. Got %d", expect, obtain)
	}
}

func testUint16(t *testing.T, wlk *readers.ByteWalker, data []byte) {
	expect := binary.BigEndian.Uint16(data)
	obtain := wlk.Uint16()
	if obtain != expect {
		t.Fatalf("expected uint16 %d. Got %d", expect, obtain)
	}
}

func testUint32(t *testing.T, wlk *readers.ByteWalker, data []byte) {
	expect := binary.BigEndian.Uint32(data)
	obtain := wlk.Uint32()
	if obtain != expect {
		t.Fatalf("expected uint32 %d. Got %d", expect, obtain)
	}
}
