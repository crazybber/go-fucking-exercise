package basic

import (
	"fmt"
	"testing"
)

//阶乘

//递归思想
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

//非递归思想
func factNoRecursion(n int) int {
	if n == 0 {
        return 1
	}
	result := 1
	for i:=1; i<= n; i++ {
		result = result *i
	}

	return result
}


func RRecursion() {
    fmt.Println(fact(7))
}

func TestRRecursion(t *testing.T) {

	RRecursion()

	assertEq(6,fact(3))

	assertEq(24,fact(4))

	assertEq(6,factNoRecursion(3))

	assertEq(24,factNoRecursion(4))
}