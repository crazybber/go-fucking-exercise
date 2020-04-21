package routine

import (
	"fmt"
	"testing"
)

func channels() {

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}

func TestChannel(t *testing.T) {

	channels()

	messages := make(chan int)

	go func() {
		messages <- 2
		messages <- 3
	}()

	msg := <-messages

	assertEq(2, msg)

	msg = <-messages

	assertEq(3, msg)

}
