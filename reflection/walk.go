package reflection

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := range val.NumField() {
		item := val.Field(i)
		if item.Kind() == reflect.String {

			fn(item.String())
		}
		if item.Kind() == reflect.Struct {
			walk(item.Interface(), fn)
		}
	}

}
