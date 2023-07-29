// Code generated by "stringer -type=FontVariants"; DO NOT EDIT.

package gist

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FontVarNormal-0]
	_ = x[FontVarSmallCaps-1]
	_ = x[FontVariantsN-2]
}

const _FontVariants_name = "FontVarNormalFontVarSmallCapsFontVariantsN"

var _FontVariants_index = [...]uint8{0, 13, 29, 42}

func (i FontVariants) String() string {
	if i < 0 || i >= FontVariants(len(_FontVariants_index)-1) {
		return "FontVariants(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FontVariants_name[_FontVariants_index[i]:_FontVariants_index[i+1]]
}

func (i *FontVariants) FromString(s string) error {
	for j := 0; j < len(_FontVariants_index)-1; j++ {
		if s == _FontVariants_name[_FontVariants_index[j]:_FontVariants_index[j+1]] {
			*i = FontVariants(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: FontVariants")
}

var _FontVariants_descMap = map[FontVariants]string{
	0: ``,
	1: ``,
	2: ``,
}

func (i FontVariants) Desc() string {
	if str, ok := _FontVariants_descMap[i]; ok {
		return str
	}
	return "FontVariants(" + strconv.FormatInt(int64(i), 10) + ")"
}
