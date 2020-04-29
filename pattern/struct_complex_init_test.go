package pattern

import (
	"testing"
	"time"
)

//User for user
type User struct {
	ID   uint64
	Name string
}

//Session inherit user
type Session struct {
	User
	Timestamp time.Time
}

//Subscription for user
//Subscription is a session
type Subscription struct {
	Session
	ch chan string //发送队列
}

func TestInitialize(t *testing.T) {

	sub := Subscription{
		Session: Session{User{ID: uint64(123)}, time.Now()},
		ch:      make(chan string, 2),
	}

	show(t, sub)
	sub = Subscription{
		Session: Session{User: User{ID: uint64(123)}, Timestamp: time.Now()},
		ch:      make(chan string, 2),
	}

	show(t, sub)
}
