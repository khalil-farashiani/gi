// Package enums provides common interfaces for enums
// and bit flag enums and utilities for using them
package enums

import "fmt"

// Enumer is the interface that all enum types satisfy.
// Enum types must be convertable to and from strings,
// and must be able to return all possible enum values
// and string representations.
type Enumer interface {
	fmt.Stringer
	// SetString sets the enum value from its
	// string representation, and returns an
	// error if the string is invalid.
	SetString(s string) error
	// Values returns all possible values this
	// enum type has.
	Values() []int64
	// Strings returns the string encodings of
	// all possible values this enum type has.
	Strings() []string
}

// BitEnumer is the interface that all bit flag enum types
// satisfy. Bit flag enum types support all of the operations
// that standard enums do, and additionally can get and set
// bit flags.
type BitEnumer interface {
	Enumer
	// HasBitFlag returns whether these flags
	// have the given flag set.
	HasBitFlag(f int64) bool
	// SetBitFlag sets the value of the given
	// flag in these flags to the given value.
	SetBitFlag(f int64, b bool)
}
