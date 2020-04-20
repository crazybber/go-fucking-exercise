package process

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestEnvironmentVars(t *testing.T) {

}

func environmentVars() {

	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
