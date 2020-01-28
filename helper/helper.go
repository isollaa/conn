package helper

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/hokaccha/go-prettyjson"
)

func GetPackageName(method interface{}) string {
	path := strings.Split(runtime.FuncForPC(reflect.ValueOf(method).Pointer()).Name(), "/")
	value := strings.Split(path[len(path)-1], ".")
	if len(value) == 0 {
		return path[len(path)-1]
	}
	return value[0]
}

func PrintPretty(result interface{}) error {
	v, err := prettyjson.Marshal(result)
	if err != nil {
		return err
	}
	fmt.Println(string(v))
	return nil
}
