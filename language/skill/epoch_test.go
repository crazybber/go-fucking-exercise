package skill

import (
	"fmt"
	"testing"
	"time"
)

//epoch usually means unix timestamp
func TestEpochFunctions(t *testing.T) {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()

	pShow(secs)

	pShow(nanos)

	assertNotEq(0, time.Unix(secs, 0))
	assertNotEq(0, time.Unix(0, nanos))

}

func epochFunctions() {

	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
