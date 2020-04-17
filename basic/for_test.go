package basic

import (
	"fmt"
	"testing"
)

func ffor() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func TestFfor(t *testing.T) {

	ffor()

	result := 0
	for j := 7; j <= 9; j++ {
		result = j
	}

	assertEq(9, result)

	target := 0
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		target = n % 2
	}

	expect(1 == target)

	assertEq(1, target)

}
