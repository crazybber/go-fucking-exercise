package skill

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRandomNumbers(t *testing.T) {

	randomNumbers()

	assert.GreaterOrEqual(t, 100, rand.Intn(100))

	assert.GreaterOrEqual(t, float64(1), rand.Float64())

	assert.GreaterOrEqual(t, float64(10), rand.Float64()*5+5)

	//new a random source
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	assert.GreaterOrEqual(t, 100, r1.Intn(100))

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)

	assert.GreaterOrEqual(t, 100, r2.Intn(100))

	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)

	assert.GreaterOrEqual(t, 100, r3.Intn(100))

}

func randomNumbers() {

	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	fmt.Println(rand.Float64())

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
