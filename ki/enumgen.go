// Code generated by "goki generate"; DO NOT EDIT.

package ki

import (
	"errors"
	"strconv"
	"strings"
	"sync/atomic"

	"goki.dev/goki/enums"
)

var _FlagsValues = []Flags{0, 1, 2, 3, 4, 5, 6}

// FlagsN is the highest valid value
// for type Flags, plus one.
const FlagsN Flags = 7

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _FlagsNoOp() {
	var x [1]struct{}
	_ = x[Field-(0)]
	_ = x[Updating-(1)]
	_ = x[Deleted-(2)]
	_ = x[Destroyed-(3)]
	_ = x[ChildAdded-(4)]
	_ = x[ChildDeleted-(5)]
	_ = x[ChildrenDeleted-(6)]
}

var _FlagsNameToValueMap = map[string]Flags{
	`Field`:           0,
	`field`:           0,
	`Updating`:        1,
	`updating`:        1,
	`Deleted`:         2,
	`deleted`:         2,
	`Destroyed`:       3,
	`destroyed`:       3,
	`ChildAdded`:      4,
	`childadded`:      4,
	`ChildDeleted`:    5,
	`childdeleted`:    5,
	`ChildrenDeleted`: 6,
	`childrendeleted`: 6,
}

var _FlagsDescMap = map[Flags]string{
	0: `Field indicates a node is a field in its parent node, not a child in children.`,
	1: `Updating flag is set at UpdateStart and cleared if we were the first updater at UpdateEnd.`,
	2: `Deleted means this node has been deleted (removed from Parent) Set just prior to calling OnDelete()`,
	3: `Destroyed means this node has been destroyed. It should be skipped in all further processing, if there are remaining pointers to it.`,
	4: `ChildAdded means one or more new children were added to the node.`,
	5: `ChildDeleted means one or more children were deleted from the node.`,
	6: `ChildrenDeleted means all children were deleted.`,
}

var _FlagsMap = map[Flags]string{
	0: `Field`,
	1: `Updating`,
	2: `Deleted`,
	3: `Destroyed`,
	4: `ChildAdded`,
	5: `ChildDeleted`,
	6: `ChildrenDeleted`,
}

// String returns the string representation
// of this Flags value.
func (i Flags) String() string {
	str := ""
	for _, ie := range _FlagsValues {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	return str
}

// BitIndexString returns the string
// representation of this Flags value
// if it is a bit index value
// (typically an enum constant), and
// not an actual bit flag value.
func (i Flags) BitIndexString() string {
	if str, ok := _FlagsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Flags value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Flags) SetString(s string) error {
	*i = 0
	return i.SetStringOr(s)
}

// SetStringOr sets the Flags value from its
// string representation while preserving any
// bit flags already set, and returns an
// error if the string is invalid.
func (i *Flags) SetStringOr(s string) error {
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _FlagsNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if val, ok := _FlagsNameToValueMap[strings.ToLower(flg)]; ok {
			i.SetFlag(true, &val)
		} else {
			return errors.New(flg + " is not a valid value for type Flags")
		}
	}
	return nil
}

// Int64 returns the Flags value as an int64.
func (i Flags) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Flags value from an int64.
func (i *Flags) SetInt64(in int64) {
	*i = Flags(in)
}

// Desc returns the description of the Flags value.
func (i Flags) Desc() string {
	if str, ok := _FlagsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// FlagsValues returns all possible values
// for the type Flags.
func FlagsValues() []Flags {
	return _FlagsValues
}

// Values returns all possible values
// for the type Flags.
func (i Flags) Values() []enums.Enum {
	res := make([]enums.Enum, len(_FlagsValues))
	for i, d := range _FlagsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Flags.
func (i Flags) IsValid() bool {
	_, ok := _FlagsMap[i]
	return ok
}

// HasFlag returns whether these
// bit flags have the given bit flag set.
func (i Flags) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(&i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given
// flags in these flags to the given value.
func (i *Flags) SetFlag(on bool, f ...enums.BitFlag) {
	var mask int64
	for _, v := range f {
		mask |= 1 << v.Int64()
	}
	in := int64(*i)
	if on {
		in |= mask
		atomic.StoreInt64((*int64)(i), in)
	} else {
		in &^= mask
		atomic.StoreInt64((*int64)(i), in)
	}
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Flags) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Flags) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}
