package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func variables() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)
	f := "apple"
	fmt.Println(f)
}

func TestVariables(t *testing.T) {

	variables()

	var a = "initial"

	assert.EqualValues(t, a, "initial")

	var b, c int = 1, 2

	expect(b == 1)
	expect(c == 2)

	b++
	expect(b == 2)

	c--

	assert.Equal(t, 1, c)

	fmt.Println("c--", c)

	var e int

	assert.Equal(t, 0, e)

	f := "apple"

	assert.Equal(t, "apple", f)

}

//32位浮点数float类型是7为有效数字,6位小数。
//32位浮点数的0，实际是0.000000
//float遵从的是IEEE 754标准 float32 是其中的R32.24 ,而float64 是R64.53。
//https://blog.csdn.net/u012184539/article/details/88020956
//https://www.cnblogs.com/jillzhang/archive/2007/06/24/793901.html
func TestVariablesFloats(t *testing.T) {

	var b, c float64 = 1.0, 2.0

	assertEq(1.0, b) //OK
	assertEq(2.0, c) //OK

	var d, e float32 = 1.0, 2.0

	assertEq(1.0, d) // must Failed in float32,Failed is right
	assertEq(2.0, e) //  must Failed in float32,Failed is right

	result := d - 1.0

	t.Logf("d-1.0= %f", result)

	assert.LessOrEqual(t, result, float32(0.000000))

	var f float64

	//Special 0 and 0.0
	assertEq(0, f)   //will Fail ,0 for int
	assertEq(0.0, f) //this is ok ,0.0 for float

	f = 1
	assertEq(1.0, f)

	f = 1.5
	assertEq(1.5, f)

	f = 2.0

	t.Logf("f-2=%f", f-2.0)

}
