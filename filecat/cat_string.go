// Code generated by "stringer -type=Cat"; DO NOT EDIT.

package filecat

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[Folder-1]
	_ = x[Archive-2]
	_ = x[Backup-3]
	_ = x[Code-4]
	_ = x[Doc-5]
	_ = x[Sheet-6]
	_ = x[Data-7]
	_ = x[Text-8]
	_ = x[Image-9]
	_ = x[Model-10]
	_ = x[Audio-11]
	_ = x[Video-12]
	_ = x[Font-13]
	_ = x[Exe-14]
	_ = x[Bin-15]
	_ = x[CatN-16]
}

const _Cat_name = "UnknownFolderArchiveBackupCodeDocSheetDataTextImageModelAudioVideoFontExeBinCatN"

var _Cat_index = [...]uint8{0, 7, 13, 20, 26, 30, 33, 38, 42, 46, 51, 56, 61, 66, 70, 73, 76, 80}

func (i Cat) String() string {
	if i < 0 || i >= Cat(len(_Cat_index)-1) {
		return "Cat(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Cat_name[_Cat_index[i]:_Cat_index[i+1]]
}

func (i *Cat) FromString(s string) error {
	for j := 0; j < len(_Cat_index)-1; j++ {
		if s == _Cat_name[_Cat_index[j]:_Cat_index[j+1]] {
			*i = Cat(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: Cat")
}
