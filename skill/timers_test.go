package skill

import (
	"fmt"
	"testing"
	"time"
)

func tTimer() {

	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}

func TestTTimer(t *testing.T) {

	time.Sleep(time.Millisecond * 100)

	done := make(chan struct{})

	select {
	case tm := <-time.After(time.Second * 2):
		t.Log("Timeout event:", tm)
	}

	timer3 := time.NewTimer(time.Millisecond * 500)

	go func() {
		evt := <-timer3.C
		t.Log("Event from Timer 3 ", evt)
		done <- struct{}{}
		timer3.Stop()
	}()

	exit := <-done

	//duck type (not including `func type`)
	assertEq(struct{}{}, exit)
}
