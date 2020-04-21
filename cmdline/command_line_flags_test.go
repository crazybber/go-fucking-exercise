package cmdline

import (
	"flag"
	"fmt"
	"testing"
)

func TestCommandLineFlags(t *testing.T) {

	wordPtr := flag.String("hello", "world", "a string")

	numbPtr := flag.Int("numb", 42, "an int")

	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string

	flag.StringVar(&svar, "strVarName", "bar", "a string var")

	var ivar int

	flag.IntVar(&ivar, "intVarName", 200, "a int var")

	flag.Parse()

	assertEq("world", *wordPtr)
	assertEq(42, *numbPtr)
	assertEq(false, *boolPtr)
	assertEq("bar", svar)
	assertEq(200, ivar)

	assertEq([]string{}, flag.Args())
}

func commandLineFlagsTest() {

	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
