package reflection

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(v reflect.Value) {
		walk(v.Interface(), fn)
	}

	var getField func(int) reflect.Value
	var iterate int

	switch val.Kind() {
	case reflect.Struct:
		iterate = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		iterate = val.Len()
		getField = val.Index
	case reflect.String:
		fn(val.String())
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for {
			v, ok := val.Recv()

			if ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}

	for i := range iterate {
		walk(getField(i).Interface(), fn)
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
