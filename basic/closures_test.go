package basic

import (
	"fmt"
	"strconv"
	"testing"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func helloClosure(name string) func(int) string {
	return func(num int) string {
		return "hello " + name + " " + strconv.Itoa(num)
	}
}

func cClosures() {

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}

func TestCClosures(t *testing.T) {
	cClosures()

	nextInt := intSeq()

	assertEq(1, nextInt())

	hello := helloClosure("crazybber")

	assertEq("hello crazybber 3", hello(3))
}
