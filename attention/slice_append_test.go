package attention

import (
	"fmt"
	"testing"
)

////////////////////////////////
//注意Slice不能直接使用赋值(=)的方式进行赋值，必须全部逐值copy.
//Slice的本质只是在保存了元素序列的引用和窗口大小关系
////////////////////////////////
func TestSliceBasic(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	//get a basic slice
	sliceSrc := nums[2:5]
	PShow(sliceSrc)

	sliceNew1 := append(sliceSrc, nums[1])

	AssertEq([]int{2, 5, 3, 9}, sliceNew1)
	//wish sliceNew1[2,5,3,9]
	PShow(sliceNew1)

	sliceNew2 := append(sliceSrc, nums[6])

	///问题在这里，问题在这里，问题在这里和最后
	///问题在这里，问题在这里，问题在这里和最后
	//wish sliceNew2[2,5,3,101]
	AssertEq([]int{2, 5, 3, 101}, sliceNew2)
	PShow(sliceNew2)

	sliceNew3 := append(sliceSrc, nums[7])

	//wish sliceNew3[2,5,3,18]
	AssertEq([]int{2, 5, 3, 18}, sliceNew3)
	PShow(sliceNew3)

	///问题在这里，问题在这里，问题在这里
	///问题在这里，问题在这里，问题在这里
	//现在sliceNew2 是多少 还是：[2,5,3,101] 吗?
	AssertNotEq([]int{2, 5, 3, 101}, sliceNew2)
	PShow(sliceNew2)
	//想想为什么现在sliceNew2 变成了：[2,5,3,18] ?

}

func TestSliceMore(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	//get a basic slice
	sliceSrc := nums[2:5]
	PShow(sliceSrc)

	sliceNew1 := append(sliceSrc, nums[1])

	AssertEq([]int{2, 5, 3, 9}, sliceNew1)
	//wish sliceNew1[2,5,3,9]
	PShow(sliceNew1)

	sliceNew2 := append(sliceNew1, nums[6])

	///问题在这里，问题在这里，问题在这里和最后
	///问题在这里，问题在这里，问题在这里和最后
	//wish sliceNew2[2,5,3,9,101]
	AssertEq([]int{2, 5, 3, 9, 101}, sliceNew2)
	PShow(sliceNew2)

	sliceNew3 := append(sliceNew1, nums[7])

	//wish sliceNew3[2,5,3,18]
	AssertEq([]int{2, 5, 3, 9, 18}, sliceNew3)
	PShow(sliceNew3)

	///问题在这里，问题在这里，问题在这里
	///问题在这里，问题在这里，问题在这里
	//现在重新看sliceNew2，你会发现覆盖问题，
	//wish sliceNew2[2,5,3,9,101]
	AssertNotEq([]int{2, 5, 3, 9, 101}, sliceNew2)
	PShow(sliceNew2)
	PShow(sliceNew1)
	PShow(sliceNew3)

}

func TestSliceWithMapCopy(t *testing.T) {

	type element struct {
		mp map[int][]int
	}

	basicDataSet := element{
		mp: map[int][]int{
			0: {2, 3 - 5, 6, 0, -1},
			1: {-4, 6, 8, -3, 55, 2},
		},
	}
	fmt.Printf("address0: %p\n", basicDataSet.mp[0])
	fmt.Printf("address0: %p\n", basicDataSet.mp[1])
	//创建一个新的结构
	newSet1 := element{map[int][]int{}}
	for k, v := range basicDataSet.mp {
		newSet1.mp[k] = append(v, 100) //每一列后面附加100
	}

	PShow(newSet1.mp)
	fmt.Printf("address0: %p\n", newSet1.mp[0])
	fmt.Printf("address1: %p\n", newSet1.mp[1])

	//创建第二个结构
	newSet2 := element{map[int][]int{}}
	for k, v := range basicDataSet.mp {
		newSet2.mp[k] = append(v, -9876) //每一列后面附加100
	}
	PShow(newSet2.mp)
	fmt.Printf("address0: %p\n", newSet2.mp[0])
	fmt.Printf("address1: %p\n", newSet2.mp[1])

	//再看一下第一个结构
	PShow(newSet1.mp)

	fmt.Printf("address0: %p\n", newSet1.mp[0])
	fmt.Printf("address1: %p\n", newSet1.mp[1])
}

//这一个有问题的，正确的
func TestSliceWithSliceInMap(t *testing.T) {

	type element struct {
		mp map[int][]int
	}

	srcArray := []int{2, 3, -5, 6, 2, -1}

	basicDataSet := element{
		mp: map[int][]int{
			0: srcArray[2:4], //切片初始化
			1: srcArray[1:5],
		},
	}
	PShow("basicDataSet:", basicDataSet.mp)
	fmt.Printf("address0: %p\n", basicDataSet.mp[0])
	fmt.Printf("address1: %p\n", basicDataSet.mp[1])

	//创建一个新的结构
	newSet1 := element{map[int][]int{}}
	for k, v := range basicDataSet.mp {
		newSet1.mp[k] = append(v, 100) //每一列后面附加100
	}

	PShow("newSet1:", newSet1.mp)
	fmt.Printf("address0: %p\n", newSet1.mp[0])
	fmt.Printf("address1: %p\n", newSet1.mp[1])

	//创建第二个结构
	newSet2 := element{map[int][]int{}}
	for k, v := range basicDataSet.mp {
		newSet2.mp[k] = append(v, -9876) //每一列后面附加100
	}
	PShow("newSet2:", newSet2.mp)
	fmt.Printf("address0: %p\n", newSet2.mp[0])
	fmt.Printf("address1: %p\n", newSet2.mp[1])

	//再重新看一下第一个结构,会发生变化
	PShow("newSet1:", newSet1.mp)
	//why ? please fix it

}

//这一个是正常的，正确的
func TestSliceWithArrayInMap(t *testing.T) {

	type element struct {
		mp map[int][]int
	}

	basicDataSet := element{
		mp: map[int][]int{
			0: {2, 3 - 5, 6, 0, -1},
			1: {-4, 6, 8, -3, 55, 2},
		},
	}
	fmt.Printf("address0: %p\n", basicDataSet.mp[0])
	fmt.Printf("address1: %p\n", basicDataSet.mp[1])

	//创建一个缓存指针的数组
	var tmpIndexPointers [][]int

	tmpIndexPointers = append(tmpIndexPointers, basicDataSet.mp[0])
	tmpIndexPointers = append(tmpIndexPointers, basicDataSet.mp[2])

	fmt.Printf("tmpIndexPointers0: %p\n", basicDataSet.mp[0])
	fmt.Printf("tmpIndexPointers1: %p\n", basicDataSet.mp[1])

	//创建一个新的结构
	newSet1 := element{map[int][]int{}}
	for k, v := range tmpIndexPointers {
		newSet1.mp[k] = append(v, 100) //每一列后面附加100
	}

	PShow("newSet1:", newSet1.mp)
	fmt.Printf("address0: %p\n", newSet1.mp[0])
	fmt.Printf("address1: %p\n", newSet1.mp[1])

	//创建第二个结构
	newSet2 := element{map[int][]int{}}
	for k, v := range tmpIndexPointers {
		newSet2.mp[k] = append(v, -9876) //每一列后面附加100
	}
	PShow("newSet2:", newSet2.mp)
	fmt.Printf("address0: %p\n", newSet2.mp[0])
	fmt.Printf("address1: %p\n", newSet2.mp[1])

	//再看一下第一个结构
	PShow("newSet1:", newSet1.mp)

	fmt.Printf("address0: %p\n", newSet1.mp[0])
	fmt.Printf("address1: %p\n", newSet1.mp[1])
}
