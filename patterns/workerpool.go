package patterns

import (
	"fmt"
	"sync"
	"time"
)

// Data to be proccessed.
var taskCount = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

const (
	// Number of concurrent workers.
	numberOfWorkers = 3
)

func RunWorkerPool() {
	jobs := make(chan struct{}, numberOfWorkers)
	wg := sync.WaitGroup{}

	//add workers.
	for id := range taskCount {
		wg.Add(1)
		jobs <- struct{}{}

		go func(id int) {
			workerp(id)
			<-jobs
			defer wg.Done()
		}(id)
	}

	//wait for all workers to complete
	wg.Wait()
}

func workerp(id int) {
	fmt.Println(id, "working on")
	time.Sleep(2 * time.Second)
}

/*
When to Use This Pattern
Worker pool pattern is useful when it is required to process a large number of tasks, but we want to limit the number of concurrently executing goroutines, which will have a positive impact on code performance and avoid overloading the system. This pattern also allows easy system scaling by increasing the number of allowable concurrent workers.

A worker pool can be used for processing client requests on the server side or for performing background tasks, such as report generation or data processing.

I often use this pattern when there is a need to process and save a large amount of data rows in a database. Additionally, this pattern is very convenient for processing events from a distributed queue, such as Kafka.
*/
