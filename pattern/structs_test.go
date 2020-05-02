package pattern

import (
	"fmt"
	"testing"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 10000
	return &p
}

func sStruct() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}

func TestSStruct(t *testing.T) {
	sStruct()

	per := newPerson("crazybber")

	assertEq("crazybber", per.name)

	assertEq(10000, per.age)

	pper := person{name: "Alice", age: 30}

	assertEq("Alice", pper.name)

	assertEq(30, pper.age)

	pp := pper

	pp.age = -10

	assertEq(30, pper.age)

	s := &person{name: "Sean", age: 50}

	assertEq(50, s.age)

	assertEq("Sean", s.name)

	p := s

	p.name = "hello"

	assertEq("hello", s.name)

}
