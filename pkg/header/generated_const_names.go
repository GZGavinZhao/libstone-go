// Code generated by "stringer -type FileType -output generated_const_names.go"; DO NOT EDIT.

package header

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FileTypeUnknown-0]
	_ = x[FileTypeBinary-1]
	_ = x[FileTypeDelta-2]
	_ = x[FileTypeRepository-3]
	_ = x[FileTypeBuildManifest-4]
}

const _FileType_name = "FileTypeUnknownFileTypeBinaryFileTypeDeltaFileTypeRepositoryFileTypeBuildManifest"

var _FileType_index = [...]uint8{0, 15, 29, 42, 60, 81}

func (i FileType) String() string {
	if i >= FileType(len(_FileType_index)-1) {
		return "FileType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FileType_name[_FileType_index[i]:_FileType_index[i+1]]
}
