// Code generated by "enumgen"; DO NOT EDIT.

package testdata

import (
	"errors"
	"strconv"
	"strings"

	"goki.dev/goki/enums"
)

var _TestEnumValues = []TestEnum{0, 1}

// TestEnumN is the highest valid value
// for type TestEnum, plus one.
const TestEnumN TestEnum = 2

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _TestEnumNoOp() {
	var x [1]struct{}
	_ = x[TestValue1-(0)]
	_ = x[TestValue2-(1)]
}

var _TestEnumNameToValueMap = map[string]TestEnum{
	`TestValue1`: 0,
	`testvalue1`: 0,
	`TestValue2`: 1,
	`testvalue2`: 1,
}

var _TestEnumDescMap = map[TestEnum]string{
	0: ``,
	1: ``,
}

var _TestEnumMap = map[TestEnum]string{
	0: `TestValue1`,
	1: `TestValue2`,
}

// String returns the string representation
// of this TestEnum value.
func (i TestEnum) String() string {
	if str, ok := _TestEnumMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the TestEnum value from its
// string representation, and returns an
// error if the string is invalid.
func (i *TestEnum) SetString(s string) error {
	if val, ok := _TestEnumNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _TestEnumNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type TestEnum")
}

// Int64 returns the TestEnum value as an int64.
func (i TestEnum) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the TestEnum value from an int64.
func (i *TestEnum) SetInt64(in int64) {
	*i = TestEnum(in)
}

// Desc returns the description of the TestEnum value.
func (i TestEnum) Desc() string {
	if str, ok := _TestEnumDescMap[i]; ok {
		return str
	}
	return i.String()
}

// TestEnumValues returns all possible values
// for the type TestEnum.
func TestEnumValues() []TestEnum {
	return _TestEnumValues
}

// Values returns all possible values
// for the type TestEnum.
func (i TestEnum) Values() []enums.Enum {
	res := make([]enums.Enum, len(_TestEnumValues))
	for i, d := range _TestEnumValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type TestEnum.
func (i TestEnum) IsValid() bool {
	_, ok := _TestEnumMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i TestEnum) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *TestEnum) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}
