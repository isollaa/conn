package status

import (
	"reflect"
	"runtime"
	"strings"
)

func getPackageName(method interface{}) string {
	path := strings.Split(runtime.FuncForPC(reflect.ValueOf(method).Pointer()).Name(), "/")
	value := strings.Split(path[len(path)-1], ".")
	if len(value) == 0 {
		return path[len(path)-1]
	}
	return value[0]
}
