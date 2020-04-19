package skill

import (
    "fmt"
	"sort"
	"testing"
)


func TestSSorting(t *testing.T){
	//sSorting()

	array := []int{65,3,4,6,8,3,32}
	 
	sort.Ints(array)

	pShow(array)

	s := sort.IntsAreSorted(array)
	
	pShow(s)

	array = []int{65,3,2,4,1,8,25,3,32}

	sort.Ints(array)

	assertEq(1,array[0])

	pShow(array)

	ss := []string{"surface", "ipad", "mac pro", "mac air", "think pad", "idea pad"}
	
	sort.Strings(ss)

	assertEq("idea pad",ss[0])

	pShow(ss)

}


func TestSSortingReverse(t *testing.T){
	//sSorting()


	ss := []string{"surface", "ipad", "mac pro", "mac air", "think pad", "idea pad"}
	
	sort.Strings(ss)

	assertEq("idea pad",ss[0])

	pShow(ss)

	//sort.Sort may be not stable
	sort.Sort(sort.Reverse(sort.StringSlice(ss)))

	assertEq("think pad",ss[0])

	pShow(ss)

	array := []int{65,3,2,4,1,8,25,3,32}

	sort.Ints(array)

	assertEq(1,array[0])

	pShow(array)

	array = []int{3, 5, 4, -1, 9, 11, -14}

	sort.Ints(array)
	
	pShow(array)

	sort.Stable(sort.Reverse(sort.IntSlice(array)))

	assertEq(11,array[0])

	pShow(array)

}



func sSorting() {

    strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs)

    ints := []int{7, 2, 4}
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)

    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)
}
