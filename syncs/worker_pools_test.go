package syncs

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestWorkersQueue(t *testing.T) {

	jobsQueue := make(chan func(id int) string, 3)
	results := make(chan string, 5)

	var workerFunc = func(id int) func(order int) string {

		return func(order int) string {
			fmt.Println("worker:", id, "order:", order, "do  job")
			result := "Done: worker" + strconv.Itoa(id) + "order:" + strconv.Itoa(order) + " do  job"
			results <- result
			return result
		}

	}

	//Queue Worker
	for w := 1; w <= 3; w++ {
		jobsQueue <- workerFunc(w)
	}
	close(jobsQueue)

	index := 0
	for work := range jobsQueue {
		index++
		go work(index)
	}

	for a := 1; a <= 3; a++ {
		<-results
	}

	workerPools()

}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func workerPools() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
