package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

/**
 * job can be anything that client want to execute us, so we have to take job and make it function
 */
type jobs func() error

var wg sync.WaitGroup

func worker(allQueue chan jobs) {
	for job := range allQueue {
		// worker have to execute this jobs now
		job()
	}
}
func main() {
	// job implementation
	var link string
	fmt.Print("enter the link that you want to send request: ")
	fmt.Scan(&link)

	// channel for storing jobs which can be upto 30
	jobsQueue := make(chan jobs, 30)

	// spawn worker pool
	workerTheadLimit := 5
	for i := 1; i <= workerTheadLimit; i++ {
		go worker(jobsQueue)
	}
	myJobs := func() error {
		resp, err := http.Get(link)
		defer wg.Done() // mark job done when finished
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		// converting all this into string
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Worker fetched %d bytes from %s\n", len(body), link)
		return nil
	}
	wg.Add(1) // increment WaitGroup before submitting job
	jobsQueue <- myJobs

	wg.Wait()

}
