package basic

import (
	"fmt"
	"testing"

	. "github.com/karlseguin/expect"
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

	//if 7.0/3.0-2.333333333333333333333333 > 0.0000000000001 {
	//}

	Expect(7.0/3.0 - 2.333333333333333333333333).LessThan(0.0000000000001)

}
