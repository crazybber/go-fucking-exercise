package iostream

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestConsoleFiltersByEcho1(t *testing.T) {

	pShow("------------------------")
	filterCmd := exec.Command("bash", "-c", "xargs echo")
	filterIn, _ := filterCmd.StdinPipe()
	filterOut, _ := filterCmd.StdoutPipe()
	filterCmd.Start()
	filterIn.Write([]byte("hello"))
	filterIn.Close()
	scanner := bufio.NewScanner(filterOut)
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		if ucl == "E" {
			pShow("Receive E(xit) ,Will exit")
			break
		}
		assertEq("HELLO", ucl)
		pShow("Receive Hello")
	}

	if err := scanner.Err(); err != nil {
		pShow("error:")
		os.Exit(1)
	}
	filterCmd.Wait()
}

func TestConsoleFilters(t *testing.T) {

	pShow("------------------------")
	filterCmd := exec.Command("bash", "-c", `grep -E "hello|E"`)
	filterIn, _ := filterCmd.StdinPipe()
	filterOut, _ := filterCmd.StdoutPipe()
	filterCmd.Start()

	//mock input data
	filterIn.Write([]byte("hello\nhello\nhello\nE\n"))
	filterIn.Close()

	scanner := bufio.NewScanner(filterOut)
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())

		if ucl == "E" {
			pShow("Receive E(xit) ,Will exit")
			break
		}
		assertEq("HELLO", ucl)
		pShow("Receive Hello")
	}

	if err := scanner.Err(); err != nil {
		pShow("error:")
		os.Exit(1)
	}
	filterCmd.Wait()
}
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
