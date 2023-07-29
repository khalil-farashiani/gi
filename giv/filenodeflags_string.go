// Code generated by "stringer -type=FileNodeFlags"; DO NOT EDIT.

package giv

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FileNodeOpen-11]
	_ = x[FileNodeSymLink-12]
	_ = x[FileNodeFlagsN-13]
}

const _FileNodeFlags_name = "FileNodeOpenFileNodeSymLinkFileNodeFlagsN"

var _FileNodeFlags_index = [...]uint8{0, 12, 27, 41}

func (i FileNodeFlags) String() string {
	i -= 11
	if i < 0 || i >= FileNodeFlags(len(_FileNodeFlags_index)-1) {
		return "FileNodeFlags(" + strconv.FormatInt(int64(i+11), 10) + ")"
	}
	return _FileNodeFlags_name[_FileNodeFlags_index[i]:_FileNodeFlags_index[i+1]]
}

func StringToFileNodeFlags(s string) (FileNodeFlags, error) {
	for i := 0; i < len(_FileNodeFlags_index)-1; i++ {
		if s == _FileNodeFlags_name[_FileNodeFlags_index[i]:_FileNodeFlags_index[i+1]] {
			return FileNodeFlags(i + 11), nil
		}
	}
	return 0, errors.New("String: " + s + " is not a valid option for type: FileNodeFlags")
}

var _FileNodeFlags_descMap = map[FileNodeFlags]string{
	11: `FileNodeOpen means file is open -- for directories, this means that sub-files should be / have been loaded -- for files, means that they have been opened e.g., for editing`,
	12: `FileNodeSymLink indicates that file is a symbolic link -- file info is all for the target of the symlink`,
	13: ``,
}

func (i FileNodeFlags) Desc() string {
	if str, ok := _FileNodeFlags_descMap[i]; ok {
		return str
	}
	return "FileNodeFlags(" + strconv.FormatInt(int64(i), 10) + ")"
}
