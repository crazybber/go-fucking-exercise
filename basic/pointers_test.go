package basic

import (
	"fmt"
	"testing"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func pPointer() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}

func emptyValue(str string) {
	str = ""
}

func emptyPtr(pstr *string) {
	*pstr = ""
}

func TestPPointer(t *testing.T) {
	pPointer()

	var intValue = 300

	zeroval(intValue)

	//a temp value var into zeroval
	assertEq(300, intValue)

	//set address of intvalue into zeroptr
	zeroptr(&intValue)

	assertEq(0, intValue)

	str := "i am crazybber"

	emptyValue(str)

	assertEq("i am crazybber", str)

	emptyPtr(&str)

	assertEq("", str)
}
