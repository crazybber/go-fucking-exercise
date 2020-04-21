package skill

import (
	"fmt"
	"strings"
	"testing"
)

// a generic usually have some basic method/interface: index/any/all/filter/may
type person struct {
	id     int32
	name   string
	age    int
	gender bool
}

func index(vs []person, t person) int {
	for i, v := range vs {
		if v.id == t.id {
			return i
		}
	}
	return -1
}

func include(vs []person, t person) bool {
	return index(vs, t) >= 0
}

//if any meet f
func any(vs []person, f func(person) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

//if all meet f
func all(vs []person, f func(person) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

//select target element sets
func filter(vs []person, f func(person) bool) []person {
	vsf := make([]person, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// map curent elements to a new list (with a new elements)
func maps(vs []person, f func(person) person) []person {
	vsm := make([]person, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func TestCollectionFunctions(t *testing.T) {

	var personList = []person{
		{3, "crazy", 16, true},
		{1, "bber", 10, false},
		{4, "new", 40, false},
		{4, "zhang", 88, true},
	}

	bber := person{1, "bber", 10, false}

	assertEq(1, index(personList, bber))

	news := person{4, "new", 40, false}

	assertEq(true, include(personList, news))

	t.Log(any(personList, func(p person) bool {
		return strings.HasPrefix(p.name, "b")
	}))

	allcheck := all(personList, func(p person) bool {
		return p.age > 9
	})

	assertEq(true, allcheck)

	t.Log(filter(personList, func(p person) bool {
		return len(p.name) > 3
	}))

	newpersonlist := maps(personList, func(p person) person {

		return person{
			id:     p.id + 1,
			name:   p.name,
			age:    p.age - 3,
			gender: p.gender,
		}

	})
	assertEq(int32(2), newpersonlist[1].id)

	t.Log(newpersonlist)

}

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func cCollectionFunctions() {

	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(Map(strs, strings.ToUpper))

}
