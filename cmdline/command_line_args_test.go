package cmdline

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestCommandLineArgs(t *testing.T) {

	commandLineArgsTest()

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	for _, item := range os.Args {
		pShow(item)
	}

	assertEq(6, len(argsWithProg))

	assertEq(true, strings.Contains(argsWithProg[0], "test"))
	pShow("------------------------")
	for _, item := range argsWithoutProg {
		pShow(item)
	}

	pShow("------------------------")
	assertEq(true, strings.Contains(argsWithoutProg[0], "log"))

	st := strings.Index(argsWithoutProg[3], "(")
	ed := strings.Index(argsWithoutProg[3], ")")

	fucName := argsWithoutProg[3][st+1 : ed]

	assertEq("TestCommandLineArgs", fucName)

	assertEq("-test.timeout=30s", arg)

}

func commandLineArgsTest() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
