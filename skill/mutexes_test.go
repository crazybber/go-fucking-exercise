package skill

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestMMutexes(t *testing.T) {

	// map{"crazybber":00860000}
	var state = make(map[string]int)
	var mutex = &sync.Mutex{}

	var readOps uint64
	var writeOps uint64

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {

				key := "index" + strconv.Itoa(rand.Intn(10))
				mutex.Lock()
				// read value
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				// simlate latency
				time.Sleep(time.Millisecond * 10)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := "index" + strconv.Itoa(rand.Intn(5))
				val := rand.Intn(100)
				mutex.Lock()
				// set value
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond * 6)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}

func mMutexes() {

	var state = make(map[int]int)

	var mutex = &sync.Mutex{}

	var readOps uint64
	var writeOps uint64

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {

				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
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

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
