package pattern

import (
	"fmt"
	"testing"
)

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2*r.width + 2*r.height
}

func mMethods() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}

func TestMMethod(t *testing.T) {
	mMethods()

	r := rect{width: 10, height: 5}

	assertEq(50, r.area())
	assertEq(30, r.perim())

	rp := &r
	rp.height = 10

	t.Logf("rect: %v", r)
	expect(r.height == 10)

	expect(r.area() == 100)

	expectEq(40, r.perim())
}
