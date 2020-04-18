package basic

import (
	"fmt"
	"testing"
)

func minus(a int, b int) int {

	return a - b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func fFunction() {

	res := minus(1, 2)
	fmt.Println("1-2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}

func TestFFunction(t *testing.T) {

	fFunction()

	assertEq(4, plusPlus(1, 1, 2))

	assertEq(1, minus(3, 2))

	assertEq(-4, minus(2, 6))
}
