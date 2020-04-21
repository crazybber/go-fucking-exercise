package iostream

import (
	"fmt"
	"net"
	"net/url"
	"testing"
)

func TestUrlParsing(t *testing.T) {
	s := "mongodb://user:pass@host.com:27127/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	assertEq(nil, err)

	assertEq("mongodb", u.Scheme)

	assertEq("user:pass", u.User.String())
	assertEq("user", u.User.Username())
	p, _ := u.User.Password()
	assertEq("pass", p)

	assertEq("host.com:27127", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	assertEq("host.com", host)
	assertEq("27127", port)

	assertEq("/path", u.Path)
	assertEq("f", u.Fragment)

	assertEq("k=v", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)

	target := map[string][]string{"k": {"v"}}

	assertEq(url.Values(target), m)

	pShow(m)
	assertEq("v", m["k"][0])
}
func urlParsing() {

	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)

	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
