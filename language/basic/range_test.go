package basic

import (
	"fmt"
	"testing"
)

func rRange() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func TestRRange(t *testing.T) {
	rRange()

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	assertEq(9, sum)

	sum = 0
	for k := range nums {
		sum += k
	}

	assertEq(3, sum)

	assertEq(3, sum)

	i, c := 100, 'c'
	for i, c = range "go" {
		fmt.Println(i, c)
	}

	assertEq(1, i)
	assertEq('o', c)
}
