package network

import (
	"bufio"
	"fmt"
	"net/http"
	"testing"
)

func TestHttpClient(t *testing.T) {

	resp, err := http.Get("https://bing.com")
	assertEq(nil, err)
	defer resp.Body.Close()
	assertEq("200 OK", resp.Status)

	scanner := bufio.NewScanner(resp.Body)

	for i := 0; scanner.Scan() && i < 10; i++ {
		pShow(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

func httpClientTest() {

	resp, err := http.Get("http://bing.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
