package main

import (
	"fmt"
	"math"
	"time"
)

var MaxInt int64 = 100000000
var totalPrimeNum int64 = 0

func checkPrime(x int) {
	if x < 2 {
		return
	}
	if x == 2 {
		totalPrimeNum++
		return
	}
	if x&1 == 0 { // x is even, so not prime
		return
	}
	// Correct square root calculation
	sqrtX := int(math.Sqrt(float64(x)))

	for i := 3; i <= sqrtX; i++ {
		if x%i == 0 {
			return
		}
	}
	// If no divisor found, it's prime
	totalPrimeNum++
}


func main() {
 start := time.Now()

 for i:=3;i<int(MaxInt); i++{
	checkPrime(i);
 }

 fmt.Println("checking till", MaxInt,"found", totalPrimeNum, "prime numbers.took",time.Since(start)); 

}
