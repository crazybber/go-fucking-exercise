package basic

import (
	"fmt"
	"reflect"
	"testing"
)

func multipleReturnValues() (int, int) {
	return 3, 7
}

func main() {

	a, b := multipleReturnValues()
	fmt.Println(a)
	fmt.Println(b)

	_, c := multipleReturnValues()
	fmt.Println(c)
}

func multipleReturnMoreValues() (int, string, func()) {
	return 8, "string", func() {}
}

func TestMultiReturnValues(t *testing.T) {

	a, b := multipleReturnValues()

	assertEq(3, a)
	assertEq(7, b)

	c, d, e := multipleReturnMoreValues()

	assertEq(8, c)

	assertEq("string", d)

	//func must be deep equal<with sample address> ,
	//func can't be compared by value assertion
	//expectNotEq(func() {}, e)
	//assertEq(func() {}, func() {})

	expectNotEq(reflect.ValueOf(func() {}).Pointer(), reflect.ValueOf(e).Pointer())

	assertEq(reflect.TypeOf(func() {}), reflect.TypeOf(e))
}
