// Code generated by "stringer -type=VectorEffects"; DO NOT EDIT.

package gist

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[VecEffNone-0]
	_ = x[VecEffNonScalingStroke-1]
	_ = x[VecEffN-2]
}

const _VectorEffects_name = "VecEffNoneVecEffNonScalingStrokeVecEffN"

var _VectorEffects_index = [...]uint8{0, 10, 32, 39}

func (i VectorEffects) String() string {
	if i < 0 || i >= VectorEffects(len(_VectorEffects_index)-1) {
		return "VectorEffects(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _VectorEffects_name[_VectorEffects_index[i]:_VectorEffects_index[i+1]]
}

func (i *VectorEffects) FromString(s string) error {
	for j := 0; j < len(_VectorEffects_index)-1; j++ {
		if s == _VectorEffects_name[_VectorEffects_index[j]:_VectorEffects_index[j+1]] {
			*i = VectorEffects(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: VectorEffects")
}

var _VectorEffects_descMap = map[VectorEffects]string{
	0: ``,
	1: `VecEffNonScalingStroke means that the stroke width is not affected by transform properties`,
	2: ``,
}

func (i VectorEffects) Desc() string {
	if str, ok := _VectorEffects_descMap[i]; ok {
		return str
	}
	return "VectorEffects(" + strconv.FormatInt(int64(i), 10) + ")"
}
