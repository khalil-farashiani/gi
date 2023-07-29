// Code generated by "stringer -type=FontWeights"; DO NOT EDIT.

package gist

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[WeightNormal-0]
	_ = x[Weight100-1]
	_ = x[WeightThin-2]
	_ = x[Weight200-3]
	_ = x[WeightExtraLight-4]
	_ = x[Weight300-5]
	_ = x[WeightLight-6]
	_ = x[Weight400-7]
	_ = x[Weight500-8]
	_ = x[WeightMedium-9]
	_ = x[Weight600-10]
	_ = x[WeightSemiBold-11]
	_ = x[Weight700-12]
	_ = x[WeightBold-13]
	_ = x[Weight800-14]
	_ = x[WeightExtraBold-15]
	_ = x[Weight900-16]
	_ = x[WeightBlack-17]
	_ = x[WeightBolder-18]
	_ = x[WeightLighter-19]
	_ = x[FontWeightsN-20]
}

const _FontWeights_name = "WeightNormalWeight100WeightThinWeight200WeightExtraLightWeight300WeightLightWeight400Weight500WeightMediumWeight600WeightSemiBoldWeight700WeightBoldWeight800WeightExtraBoldWeight900WeightBlackWeightBolderWeightLighterFontWeightsN"

var _FontWeights_index = [...]uint8{0, 12, 21, 31, 40, 56, 65, 76, 85, 94, 106, 115, 129, 138, 148, 157, 172, 181, 192, 204, 217, 229}

func (i FontWeights) String() string {
	if i < 0 || i >= FontWeights(len(_FontWeights_index)-1) {
		return "FontWeights(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FontWeights_name[_FontWeights_index[i]:_FontWeights_index[i+1]]
}

func (i *FontWeights) FromString(s string) error {
	for j := 0; j < len(_FontWeights_index)-1; j++ {
		if s == _FontWeights_name[_FontWeights_index[j]:_FontWeights_index[j+1]] {
			*i = FontWeights(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: FontWeights")
}

var _FontWeights_descMap = map[FontWeights]string{
	0:  ``,
	1:  ``,
	2:  ``,
	3:  ``,
	4:  ``,
	5:  ``,
	6:  ``,
	7:  ``,
	8:  ``,
	9:  ``,
	10: ``,
	11: ``,
	12: ``,
	13: ``,
	14: ``,
	15: ``,
	16: ``,
	17: ``,
	18: ``,
	19: ``,
	20: ``,
}

func (i FontWeights) Desc() string {
	if str, ok := _FontWeights_descMap[i]; ok {
		return str
	}
	return "FontWeights(" + strconv.FormatInt(int64(i), 10) + ")"
}
