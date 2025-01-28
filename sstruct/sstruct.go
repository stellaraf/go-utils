package sstruct

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// GetTag gets the value of a struct tag by tag name.
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

// SetValue sets the value of a struct member by string key.
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

// FromMap parses a map and converts it to a struct.
//
// This is just a convenience wrapper that marshals the map to JSON and then unmarshals it to the
// provided struct type, and is probably wildly inefficient. But it's handy. Like your mom.
func FromMap[T any, MV any](m map[string]MV) (*T, error) {
	b, err := json.Marshal(&m)
	if err != nil {
		err := errors.Join(err, errors.New("failed to marshal input map to JSON"))
		return nil, err
	}
	var t *T
	err = json.Unmarshal(b, &t)
	if err != nil {
		err := errors.Join(err, errors.New("failed to unmarshal input map from JSON to struct"))
		return nil, err
	}
	return t, nil
}
