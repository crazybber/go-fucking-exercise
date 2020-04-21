package  routine

import (
	"fmt"
	"testing"
)


func rRangeChannel() {

    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    for elem := range queue {
        fmt.Println(elem)
    }
}


func TestRRangeChannel(t *testing.T) {
	rRangeChannel()

    queue := make(chan string, 3)
    queue <- "str"
	queue <- "str"
	queue <- "str"

    close(queue)

	for elem := range queue {
		assertEq("str",elem)
    }
}