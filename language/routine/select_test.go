package routine

import (
	"fmt"
	"testing"
	"time"
)

func sSelect() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

func TestSSelect(t *testing.T) {
	c1 := make(chan int)
	c2 := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- 1
	}()
	go func() {
		time.Sleep(1500 * time.Millisecond)
		c2 <- 2
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("received", msg1)
		expect(msg1 == 1)
	case msg2 := <-c2:
		fmt.Println("received", msg2)
		assertEq(msg2, 2)
	}
	//time.Sleep(time.Second)
	select {
	case msg1 := <-c1:
		expect(msg1 == 1)
	case msg2 := <-c2:
		assertEq(msg2, 2)
	}

}
