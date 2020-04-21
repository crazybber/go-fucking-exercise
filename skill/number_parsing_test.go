package skill

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNumberParsing(t *testing.T) {

	f, _ := strconv.ParseFloat("1.234", 64)
	assertEq(float64(1.234), f)

	i, _ := strconv.ParseInt("123", 0, 64)
	assertEq(int64(123), i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	assertEq(int64(456), d)

	u, _ := strconv.ParseUint("789", 0, 64)
	assertEq(uint64(456), u)

	k, _ := strconv.Atoi("135")
	assertEq(135, k)

	_, e := strconv.Atoi("wat")
	assertNotEq(nil, e)
}

func numberParsing() {

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
