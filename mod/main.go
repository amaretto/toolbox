package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initialBufSize = 10000
	maxBufSize     = 100000000
	mod            = 1000000007
)

var sc *bufio.Scanner = func() *bufio.Scanner {
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, initialBufSize)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords)
	return sc
}()

func scanString() string { sc.Scan(); return sc.Text() }
func scanInt() int       { n, _ := strconv.Atoi(scanString()); return n }

// This is just example codes mod 100000007
func main() {
	N := scanInt()
	n, e, t := 1, 1, 1
	for i := 0; i < N; i++ {
		// [+/*]you can divide by mod each add/time ops
		e = (e * 8) % mod
		n = (n * 9) % mod
		t = (t * 10) % mod
	}
	ans := t - n - n + e
	// [+]consider if above time calc result over mod
	ans %= mod
	// [-]consider if above sub calc result goes negative number
	ans = (ans + mod) % mod
	fmt.Println(ans)
}
