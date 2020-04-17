package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var t *testing.T

func init() {

	t = new(testing.T)
}

func expect(expression bool) {
	if !expression {
		t.Errorf("wrong expression result for %v", expression)
	}

}

func assertEq(expected interface{}, actual interface{}) {
	assert.Equal(t, expected, actual)
}
