// Code generated by "goki generate"; DO NOT EDIT.

package textbuf

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"goki.dev/enums"
)

var _CasesValues = []Cases{0, 1, 2, 3, 4, 5, 6}

// CasesN is the highest valid value
// for type Cases, plus one.
const CasesN Cases = 7

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _CasesNoOp() {
	var x [1]struct{}
	_ = x[LowerCase-(0)]
	_ = x[UpperCase-(1)]
	_ = x[CamelCase-(2)]
	_ = x[LowerCamelCase-(3)]
	_ = x[SnakeCase-(4)]
	_ = x[UpperSnakeCase-(5)]
	_ = x[KebabCase-(6)]
}

var _CasesNameToValueMap = map[string]Cases{
	`LowerCase`:      0,
	`lowercase`:      0,
	`UpperCase`:      1,
	`uppercase`:      1,
	`CamelCase`:      2,
	`camelcase`:      2,
	`LowerCamelCase`: 3,
	`lowercamelcase`: 3,
	`SnakeCase`:      4,
	`snakecase`:      4,
	`UpperSnakeCase`: 5,
	`uppersnakecase`: 5,
	`KebabCase`:      6,
	`kebabcase`:      6,
}

var _CasesDescMap = map[Cases]string{
	0: ``,
	1: ``,
	2: `CamelCase is init-caps`,
	3: `LowerCamelCase has first letter lower-case`,
	4: `SnakeCase is snake_case -- lower with underbars`,
	5: `UpperSnakeCase is SNAKE_CASE -- upper with underbars`,
	6: `KebabCase is kebab-case -- lower with -&#39;s`,
}

var _CasesMap = map[Cases]string{
	0: `LowerCase`,
	1: `UpperCase`,
	2: `CamelCase`,
	3: `LowerCamelCase`,
	4: `SnakeCase`,
	5: `UpperSnakeCase`,
	6: `KebabCase`,
}

// String returns the string representation
// of this Cases value.
func (i Cases) String() string {
	if str, ok := _CasesMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Cases value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Cases) SetString(s string) error {
	if val, ok := _CasesNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _CasesNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Cases")
}

// Int64 returns the Cases value as an int64.
func (i Cases) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Cases value from an int64.
func (i *Cases) SetInt64(in int64) {
	*i = Cases(in)
}

// Desc returns the description of the Cases value.
func (i Cases) Desc() string {
	if str, ok := _CasesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// CasesValues returns all possible values
// for the type Cases.
func CasesValues() []Cases {
	return _CasesValues
}

// Values returns all possible values
// for the type Cases.
func (i Cases) Values() []enums.Enum {
	res := make([]enums.Enum, len(_CasesValues))
	for i, d := range _CasesValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Cases.
func (i Cases) IsValid() bool {
	_, ok := _CasesMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Cases) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Cases) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println(err)
	}
	return nil
}