package reflection

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := range val.NumField() {
		item := val.Field(i)
		switch item.Kind() {
		case reflect.String:
			fn(item.String())
		case reflect.Struct:
			walk(item.Interface(), fn)
		}
	}

}
