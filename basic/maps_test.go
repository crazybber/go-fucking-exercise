package basic

import (
	"fmt"
	"testing"
)

func mMap() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}

func TestMMap(t *testing.T) {

	mMap()

	m := make(map[string]int)

	assertEq(0, len(m))

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	expect(m["k2"] == 13)

	v1 := m["k1"]

	fmt.Println("v1: ", v1)

	assertEq(7, v1)

	fmt.Println("len--->:", len(m))

	assertEq(2, len(m))

	delete(m, "k2")

	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	assertEq(false, prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	assertEq(2, len(n))

	//wil return 0,if map can't find element
	assertEq(0, n["abc"])

}
