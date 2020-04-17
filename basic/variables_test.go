package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func variables() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)
	f := "apple"
	fmt.Println(f)
}

func TestVariables(t *testing.T) {

	variables()

	var a = "initial"

	assert.EqualValues(t, a, "initial")

	var b, c int = 1, 2

	expect(b == 1)
	expect(c == 2)

	var e int

	assert.Equal(t, 0, e)

	f := "apple"

	assert.Equal(t, "apple", f)

}
