package routine

import (
	"fmt"
	"testing"
)

func channelsBuffer() {

	messages := make(chan string, 1)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}

func TestChannelsWithBuffer(t *testing.T) {

	channelsBuffer()

	messages := make(chan string, 2)

	go func() {
		messages <- "pinging 1"
		messages <- "pinging 2"
		messages <- "pinging 3"
	}()

	msg := <-messages

	assertEq("pinging 1", msg)

	msg = <-messages

	assertEq("pinging 2", msg)

	msg = <-messages

	assertEq("pinging 3", msg)

}
