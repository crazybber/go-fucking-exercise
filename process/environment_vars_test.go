package process

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestEnvironmentVars(t *testing.T) {

	os.Setenv("testenv", "test")

	assertEq("test", os.Getenv("testenv"))

	assertEq("", os.Getenv("EMPTY"))

	//only work in windows
	assertEq("C:", os.Getenv("SystemDrive"))

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		pShow(pair[0], "-->", pair[1])
	}

}

func environmentVars() {

	os.Setenv("FOO", "1")
	pShow("FOO:", os.Getenv("FOO"))
	pShow("BAR:", os.Getenv("BAR"))

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
