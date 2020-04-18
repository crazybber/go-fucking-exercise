package basic

import (
	"fmt"

	"testing"

	"github.com/stretchr/testify/assert"
)

var t *testing.T

func init() {

	t = new(testing.T)
}

//need expression which will return true
func expect(expression bool) {
	if !expression {
		t.Logf("Failed!!!!!! ----------------------------------------------------------> expression result for [%v]", expression)
		fmt.Printf("Failed!!!!!! ----------------------------------------------------------> expression result for [%v]", expression)
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
		fmt.Printf("Failed!!!!!! ----------------------------------------------------------> expect:[%v] eq actual: [%v] \n", expected, actual)
		t.Logf("Failed!!!!!! ---------------------------------------------------------->  expect:[%v] eq actual: [%v] \n", expected, actual)
	} else {
		t.Logf("OK!!! expect [%v] eq in expression [%v]", expected, actual)
		fmt.Printf("OK!!! expect [%v] eq in expression [%v] \n", expected, actual)
	}
}

func assertNotEq(expected interface{}, actual interface{}) {
	if !assert.NotEqual(t, expected, actual) {
		fmt.Printf("Failed!!!!!! ----------------------------------------------------------> expect:[%v] eq actual: [%v] \n", expected, actual)
		t.Logf("Failed!!!!!! ---------------------------------------------------------->  expect:[%v] eq actual: [%v] \n", expected, actual)
	}
}
