package pattern

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type geometry interface {
	area() float64
	perim() float64
	name() string
}

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}
func (r rectangle) name() string {
	return "rectangle"
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func (c circle) name() string {
	return "circle"
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func iInterfaces() {
	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}

func TestIInterfaces(t *testing.T) {

	iInterfaces()

	var g geometry = rectangle{width: 3.0, height: 4.0}

	assertEq(12.0, g.area())

	assertEq(14.0, g.perim())

	g = circle{radius: 5.0}

	assert.LessOrEqual(t, g.perim(), 2*math.Pi*5.0*5.0)

	assert.Less(t, g.perim()-2*math.Pi*5.0*5.0, 0.0000000001)
}
