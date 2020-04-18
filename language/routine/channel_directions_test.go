package routine

import (
	"fmt"
	"testing"
)

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func channelDirectionsmain() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func TestChannelDirectionsmain(t *testing.T) {

	channelDirectionsmain()

	pingsCh := make(chan string, 1)
	pongsCh := make(chan string, 1)
	ping(pingsCh, "passed message")

	pong(pingsCh, pongsCh)

	out := <-pongsCh

	assertEq("passed message", out)
}
