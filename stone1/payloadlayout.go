package stone1

import (
	"io/fs"

	"github.com/zeebo/xxh3"
)

type FileType uint8

const (
	// Regular is a regular file.
	Regular FileType = iota + 1
	// Symlink is a symbolic link (source, target pair).
	Symlink
	// Directory is a directory node.
	Directory
	// CharacterDevice is a character device.
	CharacterDevice
	// BlockDevice is a block device.
	BlockDevice
	// FIFO is a FIFO node.
	FIFO
	// Socket is a UNIX socket.
	Socket
)

type Entry struct {
	FileType FileType
	value    any
}

func (e Entry) Source() []byte {
	switch e.FileType {
	case Regular:
		hashAndTarget := e.value.(tuple[xxh3.Uint128, string])
		hash := hashAndTarget.val1.Bytes()
		return hash[:]
	case Symlink:
		sourceAndTarget := e.value.(tuple[string, string])
		return []byte(sourceAndTarget.val1)
	case Directory,
		CharacterDevice,
		BlockDevice,
		FIFO,
		Socket:
		return nil
	default:
		panic("unknown value of FileType")
	}
}

func (e Entry) Target() []byte {
	switch e.FileType {
	case Regular:
		hashAndTarget := e.value.(tuple[xxh3.Uint128, string])
		return []byte(hashAndTarget.val2)
	case Symlink:
		sourceAndTarget := e.value.(tuple[string, string])
		return []byte(sourceAndTarget.val2)
	case Directory,
		CharacterDevice,
		BlockDevice,
		FIFO,
		Socket:
		target := e.value.(string)
		return []byte(target)
	default:
		panic("unknown value of FileType")
	}
}

type LayoutRecord struct {
	UID   uint32
	GID   uint32
	Mode  fs.FileMode
	Tag   uint32
	Entry Entry
}

// tuple mimics the tuple type from other languages.
type tuple[T1, T2 any] struct {
	val1 T1
	val2 T2
}
