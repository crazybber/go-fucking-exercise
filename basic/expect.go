package basic

import (
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
		t.Errorf("wrong expression result for %v", expression)
	} else {
		t.Logf("OK :%v", expression)
	}

}

func assertEq(expected interface{}, actual interface{}) {
	if !assert.Equal(t, expected, actual) {
		t.Logf("Failed expect %v but get %v", expected, actual)
	} else {
		t.Logf("OK expect %v eq %v", expected, actual)
		//fmt.Printf("OK expect %v eq %v \n", expected, actual)
	}
}

func assertNotEq(expected interface{}, actual interface{}) {
	assert.NotEqual(t, expected, actual)
}
