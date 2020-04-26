package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func sSlices() {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func TestSSlices(t *testing.T) {

	sSlices()

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	assertEq("c", s[2])

	assertEq(3, len(s))

	s = append(s, "d")
	s = append(s, "e", "f")

	t.Logf("slice length: %v", len(s))

	assertEq(6, len(s))

	c := make([]string, len(s))

	copy(c, s)

	fmt.Println("cpy c:", c)

	assertEq(c, s)

	l := s[2:5] //index 2 to index 5

	fmt.Println("sl1:", l)

	t.Logf("slice l length: %v", len(l))

	assertEq("d", l[1])

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	assertEq("f", l[3])

	h := []string{"g", "h", "i"}

	fmt.Println("dcl:", h)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	assertEq(2, twoD[1][1])

	// no twoD[1][2]
	//assertEq(nil, twoD[1][2])

	assert.Panics(t, func() { assertEq(nil, twoD[1][2]) }, "index out of range")

}

func TestArraysSetsValue(t *testing.T) {

	bytesStack := []byte{}

	assertNotEq(nil, bytesStack)

	var bytesStack1 []byte

	assertEq(0, len(bytesStack1))

}
