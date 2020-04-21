package network

import (
	"bufio"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestHttpServer(t *testing.T) {

	go func() {
		http.HandleFunc("/hello", hello)
		http.HandleFunc("/headers", headers)
		http.ListenAndServe(":8099", nil)
	}()

	time.Sleep(time.Second)

	resp, err := http.Get("http://localhost:8099/hello")
	assertEq(nil, err)
	assertEq("200 OK", resp.Status)
	scanner := bufio.NewScanner(resp.Body)
	assertEq("", scanner.Text())
	scanner.Scan()
	assertEq("hello", scanner.Text())
	resp.Body.Close()

	resp, err = http.Get("http://localhost:8099/headers")
	assertEq(nil, err)
	defer resp.Body.Close()
	assertEq("200 OK", resp.Status)
	scanner = bufio.NewScanner(resp.Body)
	assertEq("", scanner.Text())
	scanner.Scan()
	assertEq("User-Agent: Go-http-client/1.1", scanner.Text())

	for scanner.Scan() {
		pShow(scanner.Text())
	}
	assertEq(nil, scanner.Err())

}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func httpServerTest() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
