package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	initialBufSize = 10000
	maxBufSize     = 100000000
)

var sc *bufio.Scanner = func() *bufio.Scanner {
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, initialBufSize)
	sc.Buffer(buf, maxBufSize)
	return sc
}()

/////////////////////////////////////////
///////////////// Input /////////////////
/////////////////////////////////////////
// Input number
func scanIntNum() (num int) {
	sc.Scan()

	num, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return num
}

func scanIntNums() (nums []int) {
	sc.Scan()
	numString := strings.Split(sc.Text(), " ")

	nums = make([]int, len(numString))
	var err error

	for i := 0; i < len(nums); i++ {
		nums[i], err = strconv.Atoi(numString[i])
		if err != nil {
			panic(err)
		}
	}
	return nums
}

// ToDo : Input matrix

// Input string
func scanString() (s string) {
	sc.Scan()
	return sc.Text()
}

func scanStrings() []string {
	sc.Scan()
	return strings.Split(sc.Text(), " ")
}

func main() {
	input := scanIntNum()
	fmt.Println(input)
	inputs := scanIntNums()
	fmt.Println(inputs)
}
