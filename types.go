package utils

import (
	"reflect"
	"time"
)

// IsBool determines if a value is a boolean.
func IsBool(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Bool
}

// IsInt determines if a value is an integer.
func IsInt(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Int
}

// IsInt8 determines if a value is an 8 bit integer.
func IsInt8(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Int8
}

// IsInt16 determines if a value is a 16 bit integer.
func IsInt16(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Int16
}

// IsInt32 determines if a value is a 32 bit integer.
func IsInt32(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Int32
}

// IsInt64 determines if a value is a 64 bit integer.
func IsInt64(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Int64
}

// IsFloat32 determines if a value is a 32 bit float.
func IsFloat32(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Float32
}

// IsFloat64 determines if a value is a 64 bit float.
func IsFloat64(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Float64
}

// IsMap determines if a value is a map.
func IsMap(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Map
}

// IsSlice determines if a value is a slice.
func IsSlice(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Slice
}

// IsString determines if a value is a string.
func IsString(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.String
}

// IsStruct determines if a value is a struct.
func IsStruct(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Struct
}

// IsTime determines if a value is a time.Time object.
func IsTime(in any) bool {
	_, isTime := in.(time.Time)
	return isTime
}

// IsUint determines if a value is an unsigned integer.
func IsUint(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Uint
}

// IsUint8 determines if a value is an 8 bit unsigned integer.
func IsUint8(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Uint8
}

// IsUint16 determines if a value is an unsigned 16 bit integer.
func IsUint16(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Uint16
}

// IsUint32 determines if a value is an unsigned 32 bit integer.
func IsUint32(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Uint32
}

// IsUint64 determines if a value is an unsigned 64 bit integer.
func IsUint64(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Uint64
}

// SliceContains determines if a given slice contains a given item.
func SliceContains[T comparable](arr []T, item T) bool {
	for _, element := range arr {
		if element == item {
			return true
		}
	}
	return false
}

// MapHasKey determines if a given map has an entry for a given key.
func MapHasKey[T comparable](m map[T]T, key T) bool {
	_, hasKey := m[key]
	return hasKey
}
