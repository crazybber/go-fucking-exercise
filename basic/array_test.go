package basic

import (
	"fmt"
	"testing"
)

var twoD [2][3]int

func aArray() {

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
func TestAArray(t *testing.T) {
	aArray()
	var a [5]int
	assertEq(0, a[4])

	a[4] = 100

	expect(a[4] != 0)

	assertNotEq(0, a[4])

	b := [5]int{1, 2, 3, 4, 5}

	assertEq(3, b[2])

	assertEq(2, twoD[1][1])

	t.Logf("2d array length: %v", len(twoD))
	expect(len(twoD) == 2)
}
