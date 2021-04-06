package main

import (
	"fmt"
	"strconv"
)

// popcount return how many bit is 1 and bit string of the number
func popcount(n int) (int, string) {
	var count int
	var str string
	for i := n; i > 0; i = i >> 1 {
		if i&1 == 1 {
			count++
		}
		str = strconv.Itoa(i&1) + str
	}
	return count, str
}

func main() {
	var count int
	var str string
	// e.g.) until 2^10
	for i := 0; i < 1<<10; i++ {
		count, str = popcount(i)
		fmt.Println(i, str, count)
	}
}
