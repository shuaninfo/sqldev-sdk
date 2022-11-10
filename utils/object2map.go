package utils

import (
	"errors"
	"fmt"
	"reflect"
)

func ObjectToMap(obj interface{}) (map[string]string, error) {
	m := make(map[string]string)
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, errors.New(`obj is not a struct`)
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Kind() == reflect.Struct {
			continue
		}
		m[t.Field(i).Tag.Get(`json`)] = fmt.Sprintf("%v", v.Field(i).Interface())
	}
	return m, nil
}
