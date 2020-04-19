package routine

import (
	"fmt"
	"testing"
	"time"
)

func tTimeout() {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
func TestTTimeout(t *testing.T) {

	c1 := make(chan string, 1)

	go func() {
		time.Sleep(200 * time.Millisecond)
		c1 <- "result 1"
	}()
	//will block(witout default) ,until a IO occurred
	select {
	case res := <-c1:
		assertEq("result 1", res)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(200 * time.Millisecond)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		assertEq("result 2", res)
	case <-time.After(300 * time.Millisecond):
		fmt.Println("timeout 2")
	}

}
