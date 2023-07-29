// Code generated by "stringer -type=FontStyles"; DO NOT EDIT.

package gist

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FontNormal-0]
	_ = x[FontItalic-1]
	_ = x[FontOblique-2]
	_ = x[FontStylesN-3]
}

const _FontStyles_name = "FontNormalFontItalicFontObliqueFontStylesN"

var _FontStyles_index = [...]uint8{0, 10, 20, 31, 42}

func (i FontStyles) String() string {
	if i < 0 || i >= FontStyles(len(_FontStyles_index)-1) {
		return "FontStyles(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FontStyles_name[_FontStyles_index[i]:_FontStyles_index[i+1]]
}

func (i *FontStyles) FromString(s string) error {
	for j := 0; j < len(_FontStyles_index)-1; j++ {
		if s == _FontStyles_name[_FontStyles_index[j]:_FontStyles_index[j+1]] {
			*i = FontStyles(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: FontStyles")
}

var _FontStyles_descMap = map[FontStyles]string{
	0: ``,
	1: ``,
	2: ``,
	3: ``,
}

func (i FontStyles) Desc() string {
	if str, ok := _FontStyles_descMap[i]; ok {
		return str
	}
	return "FontStyles(" + strconv.FormatInt(int64(i), 10) + ")"
}
