package main

import (
	"reflect"
)

const (
	KInt    = "цел"
	KFloat  = "вещ"
	KBool   = "лог"
	KChar   = "сим"
	KString = "лит"
)

func Declaration(identifier string, vtype string) string {
	return vtype + " " + identifier
}

func Assignment(identifier string, value string) string {
	return identifier + " := " + value
}

func TypeOf(value any) string {
	switch value := value.(type) {
	case string:
		return KString
	case int:
		return KInt
	case float64:
		return KFloat
	case float32:
		return KFloat
	case reflect.Value:
		switch value.Type().Name() {
		case "string":
			return KString
		case "int", "int32", "int64":
			return KInt
		case "float32", "float64":
			return KFloat
		}
	}
	return ""
}
