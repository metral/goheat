package util

import (
	"reflect"
	"runtime"
)

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func ExtractValue(details StackDetails, key string) interface{} {
	var value interface{}

	for _, i := range details.Stack.Outputs {
		if i.OutputKey == key {
			value = i.OutputValue
		}
	}

	return value
}
