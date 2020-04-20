package iostream

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestWritingFiles(t *testing.T) {

	err := ioutil.WriteFile("./dat1", []byte("i am contents in file"), 0644)
	assertEq(nil, err)

	f, err := os.Create("./dat1")
	assertEq(nil, err)

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	assertEq(nil, err)
	t.Logf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	assertEq(nil, err)
	t.Logf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	assertEq(nil, err)
	t.Logf("wrote %d bytes\n", n4)

	w.Flush()

	f.Close()

	os.Remove("./dat1")
}

func writingFiles() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("./dat1", d1, 0644)
	check(err)

	f, err := os.Create("./dat1")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

}
