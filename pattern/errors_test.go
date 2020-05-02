package pattern

import (
	"errors"
	"fmt"
	"testing"
)

func f1(arg int) (int, error) {
	if arg == 42 {

		return -1, errors.New("can't work with 42")

	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func eErrors() {

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

func TestEErrors(t *testing.T) {
	eErrors()

	_, e := f1(10)

	assertEq(nil, e)

	_, e1 := f1(42)

	assertNotEq(nil, e1)

	_, e3 := f2(42)

	if ae, ok := e3.(*argError); ok {
		show(t, ae)
		t.Logf("got a into type cast %v", ae)
		assertEq(42, ae.arg)
		assertEq("can't work with it", ae.prob)
	}

	_, e4 := f2(3)

	assertEq(nil, e4)
}
