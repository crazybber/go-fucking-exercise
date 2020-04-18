package basic

import (
	"fmt"
	"math"
	"testing"
	"github.com/stretchr/testify/assert"
)

const s string = "constant"

func constants() {

	fmt.Println(s)
	
	const n = 500000000

	const d = 3e20 / n

	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

func TestConstants(t *testing.T) {

	constants()

	expect(s == "constant")

	const n = 500000000
	
	expect(n == 500000000)

	const d = 3e20 / n

	expect(d == 600000000000)

	assert.Equal(t,6e11,d)
	
	expect(int64(d) == 600000000000)

	assert.Equal(t,int64(600000000000),int64(d))


	expect(math.Abs(math.Sin(n) + 0.28470407323754404)  < 0.000000001)

	assert.Equal(t,0.0,math.Sin(n) + 0.28470407323754404)


}
