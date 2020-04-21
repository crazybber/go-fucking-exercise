package routine

import (
	"fmt"
	"testing"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func goroutinemain() {

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

func TestGoroutinemain(t *testing.T) {

	goroutinemain()

}

func TestGoroutine(t *testing.T) {

	go f("goroutine")

	abc := 1

	assertEq(1, abc)

	go func(msg *int) {
		for i := 0; i < 3; i++ {
			fmt.Println("single goroutine :", i)
			*msg += i
		}
	}(&abc)

	time.Sleep(time.Second)

	assertEq(4, abc)
}
