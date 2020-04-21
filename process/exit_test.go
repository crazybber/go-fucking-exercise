package process

import (
	"fmt"
	"os"
	"testing"
)

func TestExitFunction(t *testing.T) {

	defer deferExitProc()

	os.Exit(0)
}

func deferExitProc() {
	pShow("1-------------------------never show")
	go func() {
		pShow("2-------------------------never show")
	}()

}

func eExit() {

	defer fmt.Println("!")

	os.Exit(3)
}
