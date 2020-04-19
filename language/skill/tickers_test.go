package skill

import (
	"fmt"
	"testing"
    "time"
)

func tTicker() {

    ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)

    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()

    time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")
}

func TestTTicker(t *testing.T){

	ticker := time.NewTicker(500 * time.Millisecond)
	
	done := make(chan bool)
	msgQueue := make(chan int, 3)

    go func() {
        for {
            select {
			case <-done:
				close(msgQueue)
                return
            case ticker := <-ticker.C:
				t.Log("Send Message at Tick:", ticker)
				msgQueue <- 5
				
            }
        }
	}()

	go func (){
		time.Sleep(1500 * time.Millisecond)
		ticker.Stop()
		done <- true
	}()

	for {
		msg, ok := <-msgQueue
		if ok{
			assertEq(5,msg)
			continue
		}
		break
	}

}