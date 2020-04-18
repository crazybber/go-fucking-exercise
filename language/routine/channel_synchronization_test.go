package routine

import (
	"fmt"
	"testing"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func channelSynchronization() {

	done := make(chan bool, 1)

	go worker(done)

	<-done
}

func TestChannelSynchronization(t *testing.T) {
	channelSynchronization()
	done := make(chan bool, 1)
	go worker(done)
	exit := <-done

	assertEq(true, exit)
}
