package skill

import (
	"fmt"
	"testing"
	"time"
)

/*
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
*/
func TestTimeFormationFunctions(t *testing.T) {

	tm := time.Now()
	tmf := tm.Format(time.RFC3339)
	pShow(tmf)

	assertEq("2020", tmf[:4])

	t1, e := time.Parse(time.RFC3339, "2020-10-01T10:08:41+08:00")
	pShow(t1)

	assertEq("10:08AM", t1.Format("3:04PM"))
	assertEq("Thu Oct  1 10:08:41 2020", t1.Format("Mon Jan _2 15:04:05 2006"))

	assertEq("2020-10-01T10:08:41+08:00", t1.Format("2006-01-02T15:04:05.999999-07:00"))

	t2, e := time.Parse("3 04 PM", "8:41PM") //non regular formation
	assertNotEq("8 41 PM", t2)
	pShow("----------------------")

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		tm.Year(), tm.Month(), tm.Day(),
		tm.Hour(), tm.Minute(), tm.Second())

	assertEq(2020, tm.Year())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41") //will error

	assertEq(2020, tm.Year())

	assertNotEq(nil, e)

	pShow("----------------------")

	timeFormationFunctions()
}

func timeFormationFunctions() {

	t := time.Now()
	pShow(t.Format(time.RFC3339))

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	pShow(t1)

	pShow(t.Format("3:04PM"))
	pShow(t.Format("Mon Jan _2 15:04:05 2006"))

	pShow(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	pShow(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41")
	pShow(e)
}
