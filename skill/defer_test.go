package skill

import (
	"fmt"
	"os"
	"testing"
)

func TestDDefer(t *testing.T) {

	result := make(chan bool)

	go dDefer(result)

	res := <-result

	assertEq(true, res)
}

func dDefer(result chan bool) {

	f := createFile("./defer.txt")
	defer func() {
		go func() {
			closeFile(f)
			result <- true
		}()

	}()
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("removing file")
	os.Remove("./defer.txt")

}
