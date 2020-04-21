package syncs

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func wgWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func wgWorkerMain() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go wgWorker(i, &wg)
	}

	wg.Wait()
}

func working(jobs <-chan struct {
	Name string
	Age  int
}, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Logging age %d for citizen %s done\n", job.Age, job.Name)
	}
}

func TestWaitGroup(t *testing.T) {

	wgWorkerMain()

	jobsQueue := make(chan struct {
		Name string
		Age  int
	}, 3)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go working(jobsQueue, &wg)
	}

	for _, content := range []struct {
		Name string
		Age  int
	}{
		{"crazybber", 10000},
		{"alice", 10},
		{"zhangsan", 88},
	} {
		jobsQueue <- content

	}
	//remember to close channel
	close(jobsQueue)

	wg.Wait()

}
