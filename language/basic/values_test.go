package basic

import (
	"fmt"
	"testing"

	e "github.com/karlseguin/expect"
	"github.com/stretchr/testify/assert"
)

var stringsum string

func value() {

	stringsum = "go" + "lang"

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func TestValue(t *testing.T) {

	value()

	if "golang" != stringsum {
		t.Errorf("wrong result for %v", "golang")
	}

	e.Expect(true && false).ToEql(false)

	expect(7.0/3.0-2.333333333333333333333333 < 0.0000000000001)

	e.Expect(7.0/3.0 - 2.333333333333333333333333).LessThan(0.0000000000001)

	assert.False(t, true && false)
	assert.True(t, true || false)

	assert.False(t, !true)

}
