package skill

import (
	"fmt"
	s "strings"
	"testing"
)

var p = fmt.Println

func TestStringFunctions(t *testing.T) {

	//a space at end
	ss := "abcd_e-eFgK s134%& "

	assertEq(true, s.Contains(ss, "%"))
	assertEq(1, s.Index(ss, "b"))

	assertEq("abcd_e-eFgK s134%&", s.TrimSpace(ss))

	assertEq("abcd_e-efgk s134%& ", s.ToLower(ss))

	assertEq("abcd_e-eFgKs134%&", s.Replace(ss, " ", "", -1))

}

func stringFunctions() {

	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}
