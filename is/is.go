package is

import (
	"reflect"
	"time"
)

// Bool determines if a value is a boolean.
func Bool(in any) bool {
	_, ok := in.(bool)
	return ok
	// return reflect.TypeOf(in).Kind() == reflect.Bool
}

// Int determines if a value is an integer.
func Int(in any) bool {
	_, ok := in.(int)
	return ok
}

// Int8 determines if a value is an 8 bit integer.
func Int8(in any) bool {
	_, ok := in.(int8)
	return ok
}

// Int16 determines if a value is a 16 bit integer.
func Int16(in any) bool {
	_, ok := in.(int16)
	return ok
}

// Int32 determines if a value is a 32 bit integer.
func Int32(in any) bool {
	_, ok := in.(int32)
	return ok
}

// Int64 determines if a value is a 64 bit integer.
func Int64(in any) bool {
	_, ok := in.(int64)
	return ok
}

// Float32 determines if a value is a 32 bit float.
func Float32(in any) bool {
	_, ok := in.(float32)
	return ok
}

// Float64 determines if a value is a 64 bit float.
func Float64(in any) bool {
	_, ok := in.(float64)
	return ok
}

// Map determines if a value is a map.
func Map(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Map
}

// Slice determines if a value is a slice.
func Slice(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Slice
}

// String determines if a value is a string.
func String(in any) bool {
	_, ok := in.(string)
	return ok
}

// Struct determines if a value is a struct.
func Struct(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Struct
}

// Time determines if a value is a time.Time object.
func Time(in any) bool {
	_, ok := in.(time.Time)
	return ok
}

// Uint determines if a value is an unsigned integer.
func Uint(in any) bool {
	_, ok := in.(uint)
	return ok
}

// Uint8 determines if a value is an 8 bit unsigned integer.
func Uint8(in any) bool {
	_, ok := in.(uint8)
	return ok
}

// Uint16 determines if a value is an unsigned 16 bit integer.
func Uint16(in any) bool {
	_, ok := in.(uint16)
	return ok
}

// Uint32 determines if a value is an unsigned 32 bit integer.
func Uint32(in any) bool {
	_, ok := in.(uint32)
	return ok
}

// Uint64 determines if a value is an unsigned 64 bit integer.
func Uint64(in any) bool {
	_, ok := in.(uint64)
	return ok
}

// OneOf determines if the input value (first parameter) is equal to any of the trailing values.
//
// # Example
//
//	OneOf("one", "one", "two", "three") // true
//	OneOf(5, 1, 2, 3) // false
func OneOf[T comparable](searchFor T, searchIn ...T) bool {
	for _, item := range searchIn {
		if item == searchFor {
			return true
		}
	}
	return false
}
