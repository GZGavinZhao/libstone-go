package stone1

import "github.com/zeebo/xxh3"

// IndexRecord records offsets to unique files within the content when decompressed.
// This is used to split the file into the content store on disk before promoting
// to a transaction.
type IndexRecord struct {
	// Start is the index where the content starts.
	Start uint64
	// End is the index where the content ends.
	End uint64
	// Hash is the XXH3_128 hash of the content.
	Hash xxh3.Uint128
}
