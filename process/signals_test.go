package process

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestSignalFunction(t *testing.T) {
	//	signalFunction()

	sigsCh := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigsCh, syscall.SIGINT, syscall.SIGTERM)

	pShow("------------------------")
	go func() {
		pShow("wating for signal in background worker")
		sig := <-sigsCh
		pShow("get a signal from system")
		assertEq(sig, syscall.SIGTERM)
		done <- true
	}()

	go func() {
		pShow("signal it")
		sigsCh <- syscall.SIGTERM
	}()

	<-done
	pShow("signal received Process exit")

}

func signalFunction() {
	sigsCh := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigsCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigsCh
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
