package loops

import (
	"encoding/json"
	"log"
	"testing"
)

//Show Wrong test
func TestForRangeWrongProc(t *testing.T) {

	src := []string{"a-> ", "b-> ", "c-> "}
	PShow(src)
	var dec []string
	for _, v := range src {

		PShow(v)
		dec = append(dec, v)
	}
	log.Println(dec)
}

//Show Right test
func TestForRangeRightProc(t *testing.T) {
	src := []string{"a-> ", "b-> ", "c-> "}
	PShow(src)
	var dec []string
	for _, v := range src {
		value := v
		dec = append(dec, value)
	}
	log.Println(dec)
}

//Show Right test
func TestForRangeIntProc(t *testing.T) {
	src := []int{5, 1, 2, 3}
	PShow(src)
	var dec []int
	for _, v := range src {
		dec = append(dec, v)
	}
	PShow(dec)
}

//Show Right test
func TestForRangeIntArrayProc(t *testing.T) {
	src := [][]int{
		{5, 1, -2, 3},
		{4, 5, 0, 7},
		{-4, 6, 2, -1},
	}
	PShow(src)
	var dec [][]int
	for _, v := range src {
		dec = append(dec, v)
	}
	PShow("after concat:", dec)
}

//Show Right test
func TestForLoopBigIndexIssue(t *testing.T) {
	src := []int{5, 1, -2, 3}

	PShow(src)

	for i := 0; i < 4; i++ {
		abc := src[i]
		PShow("num:", abc)
	}
}

type test struct {
	Name string
}

////////////////////////////////
//对于自定义类型的,引用和地址类型会存在下列问题,注意看
////////////////////////////////
func (test test) String() string {
	out, err := json.Marshal(test)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

////////////////////////////////
//引用和地址类型会存在下列问题,注意看
////////////////////////////////
func TestForRangeWrongProcMore(t *testing.T) {

	src := []test{{"a"}, {"b"}, {"c"}}
	PShow(src)
	var dec []*test
	for _, v := range src {
		PShow(v)
		dec = append(dec, &v) //----------------------> Watch here ,this is important
	}
	log.Println(dec)
}

func TestForRangeRightProcMore(t *testing.T) {
	src := []test{{"a"}, {"b"}, {"c"}}
	PShow(src)
	var dec []*test
	for _, v := range src {
		value := v
		dec = append(dec, &value) //----------------------> Compare here,this is important
	}
	log.Println(dec)
}

//for,range 的引用型和地址类型的v变量在整个遍历过程中共用
//不能直接进行引用传递，即不同直接用来进行地址和引用传递，
//但循环内定义的变量只能为当前循环所使用
