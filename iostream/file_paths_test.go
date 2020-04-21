package iostream

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
)

func TestFilePaths(t *testing.T) {

	p := filepath.Join("dir1", "dir2", "filename")

	assertEq(`dir1\dir2\filename`, p)
	t.Log("p:", p)

	p = filepath.Join("dir1//", "filename")

	assertEq("dir1\\filename", p)

	pShow(p)

	p = filepath.Join("dir1/../dir1", "filename")

	pShow("--------------")

	assertEq("dir1\\filename", p)

	assertEq("dir1", filepath.Dir(p))
	assertEq("filename", filepath.Base(p))

	assertEq(false, filepath.IsAbs("dir/file"))
	assertEq(false, filepath.IsAbs("/dir/file")) // ??

	filename := "config.json"

	ext := filepath.Ext(filename)
	assertEq(".json", ext)

	assertEq("config", strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	assertEq("t\\file", rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	assertEq("..\\c\\t\\file", rel)

}

func filePaths() {

	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	ext := filepath.Ext(filename)
	fmt.Println(ext)

	fmt.Println(strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
