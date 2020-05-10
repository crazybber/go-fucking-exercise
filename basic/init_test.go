package basic

import (
	"fmt"
	"testing"
)

func init() {
	fmt.Print("initialized")
}

func TestInitFunc(t *testing.T) {
	fmt.Print("Test")
}
