package skill

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeDate(t *testing.T) {

	now := time.Now()
	pShow(now)

	location, _ := time.LoadLocation("Asia/Shanghai")
	then := time.Date(3020, 10, 17, 20, 34, 58, 651387237, location)
	pShow(then)

	assertEq(3020, then.Year())
	assertEq(time.October, then.Month())
	assertEq(17, then.Day())
	assertEq(20, then.Hour())
	assertEq(34, then.Minute())
	assertEq(58, then.Second())
	assertEq(651387237, then.Nanosecond())
	assertEq(location, then.Location())

	assertEq(time.Tuesday, then.Weekday())

	assertNotEq(true, then.Before(now))

	assertEq(true, then.After(now))

	assertNotEq(true, then.Equal(now))

	diff := time.Duration(time.Hour * 24)

	pShow(diff)

	assertEq(float64(24), diff.Hours())
	assertEq(float64(24*60), diff.Minutes())
	assertEq(float64(24*3600), diff.Seconds())
	pShow(diff.Nanoseconds())

	pShow(then.Add(diff))
	pShow(then.Add(-diff))

}

func timeDate() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}
