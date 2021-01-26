package main

import (
	"fmt"
	"math"
)

func main() {
	num := 23
	fmt.Println(isPrimeSmart(num))
}

// isPrime checks if the target is prime number
func isPrime(target int) bool {
	limit := int(math.Sqrt(float64(target)))

	if target < 2 {
		return false
	}

	for i := 2; i <= limit; i++ {
		if target%i == 0 {
			return false
		}
	}

	return true
}

// isPrimeSmart checks if the target is prime number by Sieve of Eratosthenes
func isPrimeSmart(target int) bool {
	limit := int(math.Sqrt(float64(target)))
	nums := make([]int, limit)

	if target < 2 {
		return false
	}
	//init
	for i := 0; i < limit; i++ {
		if i == 0 {
			continue
		}
		nums[i] = i + 1
	}

	for i := 1; i < limit; i++ {
		// skip sieved numbers
		if nums[i] == 0 {
			continue
		}
		// check prime
		if target%nums[i] == 0 {
			return false
		}
		//sieves numbers
		for j := nums[i]; j < limit; j += j {
			nums[j] = 0
		}
	}

	return true
}
