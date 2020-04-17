package basic

import (
	"testing"
)

func Helloworld() string {
	return "hello world"
}

func TestHelloworld(t *testing.T) {

	expected := "hello world"

	if Helloworld() != expected {

		t.Errorf("wrong result for %v", Helloworld())
	}

}
