package sstruct

import (
	"fmt"
	"reflect"
)

func GetTag[T any](s T, f, t string) (string, error) {
	typeof := reflect.TypeOf(s)
	var field reflect.StructField
	if typeof.Kind() == reflect.Pointer {
		rf, ok := typeof.Elem().FieldByName(f)
		if !ok {
			return "", fmt.Errorf("field '%s' does not exist on struct", f)
		}
		field = rf
	} else {
		rf, ok := typeof.FieldByName(f)
		if !ok {
			return "", fmt.Errorf("field '%s' does not exist on struct", f)
		}
		field = rf
	}
	tag := field.Tag.Get(t)
	if tag == "" {
		return "", fmt.Errorf("field '%s' exists, but the value of tag '%s' is empty or the tag is not specified", f, t)
	}
	return tag, nil
}

func SetValue[S any, V any](s S, k string, v V) (S, error) {
	r := reflect.ValueOf(s)
	if r.Kind() != reflect.Pointer {
		return s, fmt.Errorf("struct passed must be a pointer")
	}
	f := reflect.Indirect(r).FieldByName(k)
	vr := reflect.ValueOf(v)
	if f.Kind() != reflect.Invalid {
		f.Set(vr)
	}
	return s, nil
}
