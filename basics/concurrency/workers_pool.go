package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output // put the output into the results channel
	}
}

// createWorkerPool creates workers, each call the worker() function in a thread
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results) // close the results channel
}

// createJobs creates a number of jobs and put them into the jobs channel
func createJobs(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		job := Job{
			id:       i,
			randomno: rand.Intn(999),
		}
		// note that the jobs channel has capacity 10, so this writing blocks when the channel is full.
		// But as soon as the worker reads from the channel, this gets unblocked.
		jobs <- job // put a job into the channel
	}
	close(jobs) // close the channel to inform the other party the end of work
}

func printResults() {
	// keep reading until the channel is closed
	count := 0
	for res := range results {
		count++
		fmt.Println(res.job.id, ":", res.sumofdigits, "job count: ", count)
	}
}

func main() {
	// let's create 50 jobs and 10 workers
	// Note that the jobs and results channel have capacity 10 each.
	go createJobs(50)
	go createWorkerPool(10)

	// We do not need to use waitGroup for the two go(s) above
	// The reason is, the printResults func blocks until all the jobs are done,
	// meaning until the results channel is closed, which also means the two go(s) above have finished.
	// But if we ever need to wait for the go(s) above to finish, you can either:
	//    - use a waitGroup
	//    - or a buffered channel of capacity 1:
	//           a -> try to read from it here, so it blocks main thread
	//           b -> only write to it when the child threads are done
	//              when there is an item in the channel, the reading (a) gets unblocked.
	printResults()
}

// digits sums the digits in a number and return the sum
func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(1 * time.Second)
	return sum
}
