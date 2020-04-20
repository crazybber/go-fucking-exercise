package iostream

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFilesAndDirectories(t *testing.T) {

	f, err := ioutil.TempFile("", "sample")
	assertEq(nil, err)

	pShow("Temp file name:", f.Name())

	assertEq(true, strings.Contains(f.Name(), "sample"))

	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	assertEq(nil, err)

	dname, err := ioutil.TempDir("", "sampledir")

	assertEq(nil, err)
	assertEq(true, strings.Contains(dname, "sampledir"))
	pShow("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	assertEq(nil, err)

	contents, err := ioutil.ReadFile(fname)
	assertEq(nil, err)
	assertEq([]byte{1, 2}, contents)

}

func tempFilesAndDirectories() {

	f, err := ioutil.TempFile("", "sample")
	check(err)

	fmt.Println("Temp file name:", f.Name())

	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	dname, err := ioutil.TempDir("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
