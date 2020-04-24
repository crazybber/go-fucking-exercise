package loop

import (
	"testing"
	"time"
)

func selectCannotBreakFor(chExit chan bool) {
	for {
		select {
		case v, ok := <-chExit:
			PShow("try break from select case with value:", v, ok)
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	//never go here
	PShow("exit from select of select inner for")
}

func TestSelectForCanNotBreakFromFor(t *testing.T) {

	c := make(chan bool)

	go selectCannotBreakFor(c)

	c <- true

	c <- false

	close(c)

	time.Sleep(time.Millisecond * 500)
}

func selectForCanBreakWithBreakLabel(chExit chan bool) {
EXIT:
	for {
		select {
		case v, ok := <-chExit:
			PShow("try break from select case with value:", v, ok)
			break EXIT
			time.Sleep(time.Millisecond * 200)
		}
	}
	PShow("exit from select of select inner for")
}

func TestSelectForCanBreak(t *testing.T) {

	c := make(chan bool)

	go selectForCanBreakWithBreakLabel(c)

	c <- true

	c <- false

	close(c)

	time.Sleep(time.Millisecond * 500)
}

func TestSelectForCanBreakGoto(t *testing.T) {

	c := make(chan bool)

	go selectForCanBreakWithGoto(c)

	c <- true

	c <- false

	close(c)

	time.Sleep(time.Millisecond * 500)
}

func selectForCanBreakWithGoto(chExit chan bool) {

	for {
		select {
		case v, ok := <-chExit:
			PShow("try break(goto) from select case with value:", v, ok)
			goto EXIT
			time.Sleep(time.Millisecond * 200)
		}
	}
EXIT:
	PShow("exit from select of select inner for")
}
