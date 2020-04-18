package basic

import (
	"fmt"
	"testing"
	"time"
)

type typeType int

const (
	typebool typeType = iota
	typeint
	typestring
	typeNone
)

var whatTypeAmI = func(i interface{}) typeType {
	switch t := i.(type) {
	case bool:
		return typebool
	case int:
		return typeint
	case string:
		return typestring
	default:
		fmt.Printf("Don't know type %T\n", t)
		return typeNone
	}
}

var getweek = func() string {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		return "weekend"
	default:
		return "weekday"
	}
}

func sSwitch() {

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatTypeAmI(true)
	whatTypeAmI(1)
	whatTypeAmI("hey")
}

func TestSSwitch(t *testing.T) {

	sSwitch()

	expect(getweek() == "weekend")

	assertEq(typebool, whatTypeAmI(true))

	assertEq(typestring, whatTypeAmI("abcde"))

	assertEq(typeint, whatTypeAmI(123))
}
