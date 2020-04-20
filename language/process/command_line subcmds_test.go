package process

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"
)

//we will parsing test cmd lines
func TestCommandLineSubCmds(t *testing.T) {

	argsWithoutProg := os.Args[1:]

	assertEq(5, len(argsWithoutProg))

	assertEq(true, strings.Contains(argsWithoutProg[0], "test"))

	pShow("------------------------")
	for _, item := range os.Args {
		pShow(item)
	}
	pShow("------------------------")

	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("test.v", false, "enable log")
	fooRunName := fooCmd.String("test.run", "", "method name")
	timout := fooCmd.String("test.timeout", "", "time out")

	expect(len(os.Args) > 2)

	fooCmd.Parse(argsWithoutProg[1:])
	pShow("------------------------")
	pShow("parsing 'go test'")
	assertEq(true, *fooEnable)
	assertEq("^(TestCommandLineSubCmds)$", *fooRunName)
	assertEq("30s", *timout)
	pShow("fool  tail:", fooCmd.Args())

	//mock a cmdLines
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level of pressure")
	onOff := barCmd.Bool("on", true, "on or off")

	mockedArgs := []string{"-level=10", "-on=false", "abc"}

	barCmd.Parse(mockedArgs)
	pShow("------------------------")
	pShow("mocked subcommand 'bar'")
	assertEq(10, *barLevel)
	assertEq(false, *onOff)
	fmt.Println("abc", barCmd.Args())

}

func commandLineSubCmds() {

	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
