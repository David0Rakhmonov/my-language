package testing

import (
	"fmt"
	"reflect"
)

// Assert проверяет, что два значения равны
func Assert(expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		panic(fmt.Sprintf("Assertion failed: expected %v, got %v", expected, actual))
	}
}
