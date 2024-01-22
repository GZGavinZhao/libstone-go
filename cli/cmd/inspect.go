package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/der-eismann/libstone/pkg/header"
	"github.com/der-eismann/libstone/pkg/payload"
	"github.com/klauspost/compress/zstd"
	"github.com/spf13/cobra"
)

func Inspect(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("one stone file as argument required")
	}

	var pos int64

	absPath, err := filepath.Abs(args[0])
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	file, err := os.Open(absPath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	fmt.Printf("\"%s\" = stone container version V1\n", absPath)

	packageHeader, err := header.ReadHeader(io.NewSectionReader(file, 0, 32))
	if err != nil {
		return fmt.Errorf("failed to read package header: %w", err)
	}

	pos += 32

	for i := 0; i < int(packageHeader.Data.NumPayloads); i++ {
		payloadheader, err := payload.ReadPayloadHeader(io.NewSectionReader(file, pos, 32))
		if err != nil {
			return fmt.Errorf("failed to read payload header: %w", err)
		}
		//payloadheader.Print()

		pos += 32

		payloadReader, err := getCompressionReader(file, payloadheader.Compression, pos, int64(payloadheader.StoredSize))
		if err != nil {
			return fmt.Errorf("failed to get compression reader: %w", err)
		}

		pos += int64(payloadheader.StoredSize)

		switch payloadheader.Kind {
		case payload.KindMeta:
			err = payload.PrintMetaPayload(payloadReader, int(payloadheader.NumRecords))
		case payload.KindLayout:
			err = payload.PrintLayoutPayload(payloadReader, int(payloadheader.NumRecords))
		// case payload.KindIndex:
		// 	err = payload.PrintIndexPayload(payloadReader, int(payloadheader.NumRecords))
		default:
			continue
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func getCompressionReader(r io.ReaderAt, compressionType payload.Compression, offset, length int64) (io.Reader, error) {
	switch compressionType {
	case payload.CompressionNone:
		return io.NewSectionReader(r, offset, length), nil
	case payload.CompressionZstd:
		return zstd.NewReader(io.NewSectionReader(r, offset, length))
	}
	return nil, errors.New("Unknown compression type")
}
