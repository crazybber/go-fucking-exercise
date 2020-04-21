package iostream

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestLinesFilters(t *testing.T) {

	//simulate read content from console
	rdr := strings.NewReader("hello\nhello\r\nhello\ne")

	//bs := bufio.NewScanner(r)
	scanner := bufio.NewScanner(rdr)
	//	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())

		if ucl == "E" {
			pShow("Will exit")
			break
		}
		assertEq("HELLO", ucl)
		pShow("Receive Hello")

	}

	if err := scanner.Err(); err != nil {
		pShow("error:")
		os.Exit(1)
	}

}

func linesFilters() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
