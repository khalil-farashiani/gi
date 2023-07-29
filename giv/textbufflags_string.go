// Code generated by "stringer -type=TextBufFlags"; DO NOT EDIT.

package giv

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TextBufAutoSaving-24]
	_ = x[TextBufMarkingUp-25]
	_ = x[TextBufChanged-26]
	_ = x[TextBufFileModOk-27]
	_ = x[TextBufFlagsN-28]
}

const _TextBufFlags_name = "TextBufAutoSavingTextBufMarkingUpTextBufChangedTextBufFileModOkTextBufFlagsN"

var _TextBufFlags_index = [...]uint8{0, 17, 33, 47, 63, 76}

func (i TextBufFlags) String() string {
	i -= 24
	if i < 0 || i >= TextBufFlags(len(_TextBufFlags_index)-1) {
		return "TextBufFlags(" + strconv.FormatInt(int64(i+24), 10) + ")"
	}
	return _TextBufFlags_name[_TextBufFlags_index[i]:_TextBufFlags_index[i+1]]
}

func StringToTextBufFlags(s string) (TextBufFlags, error) {
	for i := 0; i < len(_TextBufFlags_index)-1; i++ {
		if s == _TextBufFlags_name[_TextBufFlags_index[i]:_TextBufFlags_index[i+1]] {
			return TextBufFlags(i + 24), nil
		}
	}
	return 0, errors.New("String: " + s + " is not a valid option for type: TextBufFlags")
}

var _TextBufFlags_descMap = map[TextBufFlags]string{
	24: `TextBufAutoSaving is used in atomically safe way to protect autosaving`,
	25: `TextBufMarkingUp indicates current markup operation in progress -- don&#39;t redo`,
	26: `TextBufChanged indicates if the text has been changed (edited) relative to the original, since last save`,
	27: `TextBufFileModOk have already asked about fact that file has changed since being opened, user is ok`,
	28: ``,
}

func (i TextBufFlags) Desc() string {
	if str, ok := _TextBufFlags_descMap[i]; ok {
		return str
	}
	return "TextBufFlags(" + strconv.FormatInt(int64(i), 10) + ")"
}
