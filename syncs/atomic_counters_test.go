package syncs

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func atomicCounter() {

	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {

				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops)
}

func TestAtomicCounter(t *testing.T) {

	var target uint32

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 10; c++ {
				atomic.AddUint32(&target, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	assertEq(uint32(100), target)

}
