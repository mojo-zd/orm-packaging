package tool

import (
	"reflect"
	"strings"
)

func NewInstanceSetValue(i interface{}, key string, value interface{}) interface{} {
	ty := GetType(reflect.TypeOf(i))
	va := GetValue(reflect.ValueOf(i))

	val := GetValue(reflect.ValueOf(reflect.New(ty).Interface()))
	for index := 0; index < ty.NumField(); index++ {
		if strings.ToLower(ty.Field(index).Name) == strings.ToLower(key) {
			switch va.Field(index).Kind() {
			case reflect.String:
			case reflect.Int64, reflect.Int:
				switch GetValue(reflect.ValueOf(value)).Kind() {
				case reflect.Int:
					v := value.(int)
					val.Field(index).SetInt(int64(v))
				case reflect.Int32:
					v := value.(int32)
					val.Field(index).SetInt(int64(v))
				case reflect.Int64:
					val.Field(index).SetInt(value.(int64))
				}
			}
		}
	}

	return val.Interface()
}

func SetValue(i interface{}, key string, value interface{}) {
	ty := GetType(reflect.TypeOf(i))
	va := GetValue(reflect.ValueOf(i))

	for index := 0; index < ty.NumField(); index++ {
		if strings.ToLower(ty.Field(index).Name) != strings.ToLower(key) {
			continue
		}

		switch va.Field(index).Kind() {
		case reflect.Int64, reflect.Int, reflect.Int32:
			switch GetValue(reflect.ValueOf(value)).Kind() {
			case reflect.Int:
				va.Field(index).SetInt(int64(value.(int)))
			case reflect.Int32:
				va.Field(index).SetInt(int64(value.(int32)))
			case reflect.Int64:
				va.Field(index).SetInt(value.(int64))
			case reflect.Float64:
				va.Field(index).SetInt(int64(value.(float64)))
			case reflect.Float32:
				va.Field(index).SetInt(int64(value.(float32)))
			}
		}

	}

}

func GetType(ty reflect.Type) (t reflect.Type) {
	if ty.Kind() == reflect.Ptr {
		t = ty.Elem()
		return
	}
	t = ty
	return
}

func GetValue(value reflect.Value) (v reflect.Value) {
	if value.Kind() == reflect.Ptr {
		v = value.Elem()
		return
	}
	v = value
	return
}
