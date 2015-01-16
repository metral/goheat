package util

import (
	"reflect"
	"runtime"
)

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func ExtractOverlordIP(details StackDetails) string {
	overlordIP := ""

	for _, i := range details.Stack.Outputs {
		if i.OutputKey == "overlord_ip" {
			overlordIP = i.OutputValue.(string)
		}
	}

	return overlordIP
}
