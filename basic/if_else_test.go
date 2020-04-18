package basic

import (
	"fmt"
	"testing"
)

func ifElse() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

}

func TestIfElse(t *testing.T) {

	ifElse()

	expect(7%2 != 0)

	assertEq(8%4, 0)

	fn := func(num int) int {
		if num < 0 {
			return -1
		} else if num < 10 {
			return 5
		} else {
			return 10
		}
	}

	assertEq(-1, fn(-10))

	assertEq(5, fn(6))

	assertEq(5, fn(0))

	assertEq(10, fn(100))

}
