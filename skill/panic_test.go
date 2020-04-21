package skill

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func pPanic() {

	f, err := os.Create("./testfile")
	if err != nil {
		panic(err)
	}
	f.WriteString("this is a test file")

	f.Close()

	os.Remove("./testfile")

	time.Sleep(time.Second)

	panic("mock a problem")

}

func TestPPanic(t *testing.T) {

	assert.Panics(t, func() { pPanic() }, "manmade panic")

}
