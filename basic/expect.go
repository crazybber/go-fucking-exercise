package basic

import (
	"fmt"
	"runtime"

	"testing"

	"github.com/logrusorgru/aurora"
	"github.com/stretchr/testify/assert"
)

var t *testing.T

func init() {

	t = new(testing.T)
}

//need expression which will return true
func expect(expression bool) {
	if !expression {
		t.Logf("Failed!!!!!! ----------------------------------------------------------> expression for [%v] \n", expression)
		fmt.Printf("Failed!!!!!! ----------------------------------------------------------> expression for [%v] \n", expression)
	} else {
		t.Logf("OK!!! eq in expression: [%v] \n", expression)
		fmt.Printf("OK!!! eq in expression: [%v] \n", expression)
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
		_, file, line, _ := runtime.Caller(1)
		targetString := fmt.Sprintf("Failed!!!!!! in file: %s, line: %d -----> expect:[%v] , actual: [%v] \n", file, line, expected, actual)
		fmt.Println(aurora.Red(targetString))
		//fmt.Println(targetString)
		t.Logf("Failed!!!!!! ---------------------------------------------------------->  expect:[%v] , actual: [%v] \n", expected, actual)
	} else {

		t.Logf("OK!!! expect [%v] eq expression: [%v]", expected, actual)
		targetString := fmt.Sprintf("OK!!! expect [%v] eq expression: [%v] \n", expected, actual)
		fmt.Println(aurora.Green(targetString))
	}
}

//AssertNotEq ..
func assertNotEq(expected interface{}, actual interface{}) {
	if !assert.NotEqual(t, expected, actual) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("Failed!!!!!! in file: %s, line: %d-----> expect:[%v] , actual: [%v] \n", file, line, expected, actual)
		t.Logf("Failed!!!!!! in file: %s, line: %d------>  expect:[%v] , actual: [%v] \n", file, line, expected, actual)
	} else {
		t.Logf("OK!!! expect [%v] not eq expression: [%v]", expected, actual)
		fmt.Printf("OK!!! expect [%v] not eq expression: [%v] \n", expected, actual)
	}
}
