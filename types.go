package utils

import (
	"reflect"
	"time"
)

func IsBool(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Bool
}

func IsInt(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Int
}

func IsInt8(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Int8
}

func IsInt16(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Int16
}

func IsInt32(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Int32
}

func IsInt64(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Int64
}

func IsFloat32(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Float32
}

func IsFloat64(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Float64
}

func IsMap(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Map
}

func IsSlice(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Slice
}

func IsString(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.String
}

func IsStruct(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Struct
}

func IsTime(in any) bool {
	_, isTime := in.(time.Time)
	return isTime
}

func IsUint(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Uint
}

func IsUint8(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Uint8
}

func IsUint16(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Uint16
}

func IsUint32(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Uint32
}

func IsUint64(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Uint64
}

func SliceContains[T comparable](arr []T, item T) bool {
	for _, element := range arr {
		if element == item {
			return true
		}
	}
	return false
}

func MapHasKey[T comparable](m map[T]T, key T) bool {
	_, hasKey := m[key]
	return hasKey
}
