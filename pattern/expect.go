package pattern

import (
	"fmt"
	"runtime"

	"testing"

	"github.com/stretchr/testify/assert"
)

var t *testing.T

func init() {

	t = new(testing.T)
}

func show(t *testing.T, tp interface{}) {
	t.Logf("type information : %v", tp)
}

//need expression which will return true
func expect(expression bool) {
	if !expression {
		funcName, file, line, _ := runtime.Caller(0)
		t.Logf("Failed!!!!!! ----------------------------------------------------------> expression for [%v]", expression)
		fmt.Printf("Failed!!!!!! in func %s file: %s, line: %d ----------------------------------------------------------> expression for [%v]", runtime.FuncForPC(funcName).Name(), file, line, expression)
	} else {
		t.Logf("OK!!! eq in expression: [%v]", expression)
	}

}

func expectEq(expected interface{}, actual interface{}) {
	assertEq(expected, actual)
}

func expectNotEq(expected interface{}, actual interface{}) {
	assertNotEq(expected, actual)
}

func assertEq(expected interface{}, actual interface{}) {

	if !assert.Equal(t, expected, actual) {
		_, file, line, _ := runtime.Caller(0)
		fmt.Printf("Failed!!!!!! in file: %s, line: %d ----------------------------------------------------------> expect:[%v] , actual: [%v] \n", file, line, expected, actual)
		t.Logf("Failed!!!!!! ---------------------------------------------------------->  expect:[%v] , actual: [%v] \n", expected, actual)
	} else {
		t.Logf("OK!!! expect [%v] eq expression: [%v]", expected, actual)
		fmt.Printf("OK!!! expect [%v] eq expression: [%v] \n", expected, actual)
	}
}

func assertNotEq(expected interface{}, actual interface{}) {
	if !assert.NotEqual(t, expected, actual) {
		fmt.Printf("Failed!!!!!! ----------------------------------------------------------> expect:[%v] , actual: [%v] \n", expected, actual)
		t.Logf("Failed!!!!!! ---------------------------------------------------------->  expect:[%v] , actual: [%v] \n", expected, actual)
	}
}
