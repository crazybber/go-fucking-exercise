package skill

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync/atomic"
	"testing"
	"time"
)

func TestStatefulGoroutines(t *testing.T) {

	type readOperation struct {
		key  int
		resp chan string
	}
	type writeOperation struct {
		key  int
		val  string
		resp chan bool
	}

	var readOpsCounts uint64
	var writeOpsCounts uint64

	readsCh := make(chan readOperation)
	writesCh := make(chan writeOperation)

	//single goroutine to handle read /write
	go func() {
		var state = make(map[int]string)
		for {
			select {
			case read := <-readsCh:
				read.resp <- state[read.key] //read a value to read chan
				pShow(read.key)
			case write := <-writesCh:
				state[write.key] = write.val
				write.resp <- true
				pShow(write.key)
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOperation{
					key:  rand.Intn(5),
					resp: make(chan string)}

				readsCh <- read
				<-read.resp
				atomic.AddUint64(&readOpsCounts, 1) //add read time
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOperation{
					key:  rand.Intn(5),
					val:  "index:" + strconv.Itoa(rand.Intn(100)),
					resp: make(chan bool)}
				writesCh <- write
				<-write.resp
				//assertEq(true, res)
				atomic.AddUint64(&writeOpsCounts, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOpsCounts)
	fmt.Println("read Counts:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOpsCounts)
	fmt.Println("write Counts:", writeOpsFinal)
}

func statefulGoroutines() {

	type readOp struct {
		key  int
		resp chan int
	}
	type writeOp struct {
		key  int
		val  int
		resp chan bool
	}

	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
