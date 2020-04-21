package routine

import (
	"fmt"
	"testing"
	"time"
)

/*

Basic sends and receives on channels are blocking.
However, we can use with a clause to implement non-blocking sends, receives,
and even non-blocking multi-way s.select default select
*/

//default case is the key branch to go in non blocking
func TestAAsyncChannelOperations(t *testing.T) {

	messages := make(chan string)
	signals := make(chan bool)

	//test receive messages without sending
	select {
	case msg := <-messages:
		assertEq(nil, "")
		fmt.Println("received message", msg)
	default:
		//will directly go here!
		t.Log("no message received")
	}

	msg := "hi"

	select {
	case messages <- msg: //here wii block ,so it will will never go ,in current case
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	time.Sleep(time.Second * 1)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
		assertEq("hi", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

func TestManyTimes(t *testing.T) {

	for i := 0; i < 200; i++ {
		TestAAsyncChannelOperations(t)
	}

}
