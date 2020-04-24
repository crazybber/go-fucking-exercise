package more

import (
	"testing"
)

var deepCopyArray = func(inputArray []int) []int {
	//intList := make([]int, len(inputArray))
	var intList []int
	for i := 0; i < len(inputArray); i++ {
		intList = append(intList, inputArray[i])
	}
	return intList
}
var deepCopyArray2 = func(inputArray []int) []int {
	intList := make([]int, len(inputArray))
	for i := 0; i < len(inputArray); i++ {
		intList[i] = inputArray[i]
	}
	return intList
}

//deepCopyMapData data from range
var deepCopyMapData = func(input [][]int, oneMore int) map[int][]int {
	cloneMap := make(map[int][]int)
	for k, v := range input {
		array := deepCopyArray(v)
		cloneMap[k] = append(array, oneMore)
	}
	return cloneMap
}

func TestDeepCopyMapWithArray(t *testing.T) {

	input := []int{10, 9, 2, 5, 3, 7, 101, 18}
	//output [2, 5, 7, 101]
	//output [2, 3, 7, 101]
	//output [2, 5, 7, 18]
	//output [2, 3, 7, 18]

	out := deepCopyArray(input)

	PShow(out)

	out = deepCopyArray2(input)

	PShow(out)
}
