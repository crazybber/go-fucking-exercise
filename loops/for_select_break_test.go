package loops

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
			if !ok {
				PShow("try break(Label) from select case with value:", v, ok)
				break EXIT //goto EXIT2
			}
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

func selectForCanBreakWithGoto(chExit chan bool) {

	for {
		select {
		case v, ok := <-chExit:
			if !ok {
				PShow("try break(goto) from select case with value:", v, ok)
				goto EXIT
			}
		}
	}
EXIT:
	PShow("exit from select of select inner for")
}

func TestSelectForCanBreakGoto(t *testing.T) {

	c := make(chan bool)

	go selectForCanBreakWithGoto(c)

	c <- true

	c <- false

	close(c)

	time.Sleep(time.Millisecond * 500)
}

func selectForCanBreakWithGotoDirectGoto(chExit chan bool) {
	for {
		select {
		case <-chExit:
			PShow("try break(goto) from select case with value:")
			goto EXIT
		}
	}
EXIT:
	PShow("exit from select of select inner for")
}

func TestSelectForCanBreakDirectGoto(t *testing.T) {

	c := make(chan bool)

	go selectForCanBreakWithGoto(c)

	c <- true

	c <- false

	close(c)

	time.Sleep(time.Millisecond * 500)
}

func selectForCanBreakWithGotoDirectBreak(chExit chan bool) {
EXIT:
	for {
		select {
		case <-chExit:
			PShow("try break(goto) from select case with value:")
			break EXIT
		}
	}
	PShow("exit from select of select inner for")
}

func TestSelectForCanBreakDirectBreak(t *testing.T) {

	c := make(chan bool)

	go selectForCanBreakWithGoto(c)

	c <- true

	c <- false

	close(c)

	time.Sleep(time.Millisecond * 500)
}
