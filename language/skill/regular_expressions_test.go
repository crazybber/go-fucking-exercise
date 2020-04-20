package skill

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"
)

func TestRegularExpressionsString(t *testing.T) {

	r, _ := regexp.Compile("p([a-z]+)ch")

	assertEq(true, r.MatchString("peach"))

	assertEq("peach", r.FindString("peach punch"))

	assertEq([]int{0, 5}, r.FindStringIndex("peach punch"))

	assertEq([]string{"peach", "ea"}, r.FindStringSubmatch("peach punch"))

	assertEq([]int{0, 5, 1, 3}, r.FindStringSubmatchIndex("peach punch"))

	assertEq([]string{"peach", "punch", "pinch"}, r.FindAllString("peach punch pinch", -1))

	assertEq([][]int{{0, 5, 1, 3}, {6, 11, 7, 9}, {12, 17, 13, 15}}, r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	assertEq([]string{"peach", "punch"}, r.FindAllString("peach punch pinch", 2))

	assertEq(true, r.Match([]byte("peach")))

	t.Log(regexp.MustCompile("p([a-z]+)ch"))

	t.Log(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	t.Log(string(out))

}

//https://www.cnblogs.com/williamjie/p/9686311.html
//https://blog.csdn.net/wangchaoqi1985/article/details/82810471/
//https://www.jianshu.com/p/82886d77440c
func TestRegularExpressions(t *testing.T) {

	regularExpressions()

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")

	assertEq(true, match)

	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	match, _ = regexp.MatchString(pattern, "123@456.cc")

	assertEq(true, match)

	pattern = `[a-zA-z]+://[^\s]*`

	match, _ = regexp.MatchString(pattern, "https://gitee.com/")

	assertEq(true, match)

	pattern = `(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`

	match, _ = regexp.MatchString(pattern, "http://www.gitee.com")

	assertEq(true, match)

	pattern = `^(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\d{8}$`

	match, _ = regexp.MatchString(pattern, "13679801234")

	assertEq(true, match)

	pattern = `(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){3}`

	match, _ = regexp.MatchString(pattern, "192.168.1.1")

	assertEq(true, match)

}

func regularExpressions() {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))

	fmt.Println(r.FindStringIndex("peach punch"))

	fmt.Println(r.FindStringSubmatch("peach punch"))

	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	fmt.Println(r.FindAllString("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
