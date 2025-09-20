package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var MaxInt int64 = 100000000
var totalPrimeNum int64 = 0
var Concurrency = 10

func checkPrime(x int) {
	if x < 2 {
		return
	}
	if x == 2 {
		atomic.AddInt64(&totalPrimeNum, 1)
		return
	}
	if x&1 == 0 {
		return
	}
	sqrtX := int(math.Sqrt(float64(x)))
	for i := 3; i <= sqrtX; i += 2 {
		if x%i == 0 {
			return
		}
	}
	atomic.AddInt64(&totalPrimeNum, 1)
}

func doBatch(id string, wg *sync.WaitGroup, start, end int) {
	defer wg.Done()

	startTime := time.Now()

	for i := start; i <= end; i++ {
		checkPrime(i)
	}

	duration := time.Since(startTime)

	fmt.Printf("Thread %s: [%d - %d] complete in %s\n", id, start, end, duration)
}


func main() {
	startTime := time.Now()

	var wg sync.WaitGroup

	batchSize := MaxInt / int64(Concurrency) //100000000/10=> 10000000

	for i := 0; i < Concurrency; i++ {
		wg.Add(1)

		start := int(i*int(batchSize)) + 1. // 1;
		end := int((i + 1) * int(batchSize))  // 10000000

		if i == Concurrency-1 {
			end = int(MaxInt)
		}

		go doBatch( strconv.Itoa(i),  &wg, start, end)
	}

	wg.Wait()


	fmt.Printf("Total prime numbers up to %d: %d\n", MaxInt, totalPrimeNum)
	fmt.Printf("Time taken: %s\n", time.Since(startTime))
	/*
	Total prime numbers up to 100000000: 5761455
  Time taken: 4.048066s
	*/
}
