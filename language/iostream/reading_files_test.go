package iostream

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestReadingFiles(t *testing.T) {

	f, err := os.Create("./dat")
	assertEq(nil, err)

	f.WriteString("this is a test string title\n")

	f.Write([]byte(`Base64 Decode online. 
	Base64Decoder is a simple and easy to use 
	online tool to decode any base64 encoded data to text. 
	It also contains several`))

	f.Close()

	dat, err := ioutil.ReadFile("./dat")
	assertEq(nil, err)
	pShow(string(dat))

	f, _ = os.Open("./dat")

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	assertEq(nil, err)
	assertEq("this ", string(b1[0:n1]))

	o2, err := f.Seek(6, 0)
	assertEq(nil, err)

	b2 := make([]byte, 2)

	n2, err := f.Read(b2)
	assertEq(nil, err)
	assertEq("s ", string(b2[0:n2]))

	fmt.Printf("%d bytes @ %d: ", n2, o2)

	o3, err := f.Seek(6, 0)
	assertEq(nil, err)
	b3 := make([]byte, 3)

	n3, err := io.ReadAtLeast(f, b3, 2)
	assertEq(nil, err)

	assertEq("s a", string(b3[0:n3]))

	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	assertEq(nil, err)
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	assertEq(nil, err)

	assertEq("this ", string(b4[0:]))

	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()

	os.Remove("./dat")

}

func readingFiles() {

	dat, err := ioutil.ReadFile("./dat")
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open("./dat")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()
}
