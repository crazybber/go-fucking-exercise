package skill

import (
	"fmt"
	"os"
	"testing"
)

type point struct {
	x, y float32
}

func TestStringFormat(t *testing.T) {

	stringFormat()

	type place struct {
		x, y float32
	}

	p := place{2.1, 3.2}

	assertEq("{2.1 3.2}", fmt.Sprintf("%v", p))

	assertEq("{x:2.1 y:3.2}", fmt.Sprintf("%+v", p))

	assertEq("skill.place{x:2.1, y:3.2}", fmt.Sprintf("%#v", p))

	assertEq("skill.place", fmt.Sprintf("%T", p))

	//reflect.TypeOf(true)
	assertEq("true", fmt.Sprintf("%t", true))

	assertEq("123", fmt.Sprintf("%d", 123))

	assertEq("1110", fmt.Sprintf("%b", 14))

	//single ASCII
	assertEq("!", fmt.Sprintf("%c", 33))

	//0x1234567abcde，十六进制
	assertEq("1c8", fmt.Sprintf("%x", 456))

	//32位，小数点后6位，有效位7位
	assertEq("78.900000", fmt.Sprintf("%f", 78.9))

	assertEq("1.234000e+09", fmt.Sprintf("%e", 1234000000.0))
	assertEq("1.234000E+09", fmt.Sprintf("%E", 1234000000.0))

	assertEq("\"string\"", fmt.Sprintf("%s", "\"string\""))

	assertEq("\"\\\"string\\\"\"", fmt.Sprintf("%q", "\"string\""))
	assertEq(`"\"string\""`, fmt.Sprintf("%q", "\"string\""))

	assertEq("6865782074686973", fmt.Sprintf("%x", "hex this"))

	//%p :pointer
	assertEq("0xc000016810", fmt.Sprintf("%p", &p))

	//right alignment
	assertEq("|    56|   789|", fmt.Sprintf("|%6d|%6d|", 56, 789))
	//right alignment,2 bit after dot
	assertEq("|  5.60|  7.89|", fmt.Sprintf("|%6.2f|%6.2f|", 5.6, 7.89))
	//left alignment
	assertEq("|1.20  |3.45  |", fmt.Sprintf("|%-6.2f|%-6.2f|", 1.2, 3.45))

	assertEq("| crazy|  bber|", fmt.Sprintf("|%6s|%6s|", "crazy", "bber"))

	assertEq("|crazy |bber  |", fmt.Sprintf("|%-6s|%-6s|", "crazy", "bber"))

	s := fmt.Sprintf("mock a %s in Log for crazy bber \n", "string")
	t.Log(s)

	fmt.Fprintf(os.Stderr, "mock an %s in console \n", "error")

}

func stringFormat() {

	p := point{1, 2}
	fmt.Printf("%v\n", p)

	fmt.Printf("%+v\n", p)

	fmt.Printf("%#v\n", p)

	fmt.Printf("%T\n", p)

	fmt.Printf("%t\n", true)

	fmt.Printf("%d\n", 123)

	fmt.Printf("%b\n", 14)

	fmt.Printf("%c\n", 33)

	fmt.Printf("%x\n", 456)

	fmt.Printf("%f\n", 78.9)

	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	fmt.Printf("%s\n", "\"string\"")

	fmt.Printf("%q\n", "\"string\"")

	fmt.Printf("%x\n", "hex this")

	fmt.Printf("%p\n", &p)

	fmt.Printf("|%6d|%6d|\n", 12, 345)

	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
