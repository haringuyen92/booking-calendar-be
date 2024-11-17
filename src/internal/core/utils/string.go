package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func ReflectString(content interface{}) string {
	reflectedValue := reflect.ValueOf(content)
	switch reflectedValue.Kind() {
	case reflect.String:
		return reflectedValue.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", reflectedValue.Int())
	case reflect.Struct:
		content, err := json.Marshal(content)
		if err != nil {
			return ""
		}
		return string(content)
	default:
		return ""
	}
}
