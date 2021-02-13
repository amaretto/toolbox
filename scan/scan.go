package scan

import (
	"bufio"
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
	// Release comment out if you need scan each words
	// sc.Split(bufio.ScanWords)
	return sc
}()

// Input string
func scanString() string    { sc.Scan(); return sc.Text() }
func scanStrings() []string { sc.Scan(); return strings.Split(sc.Text(), " ") }

// Input number
func scanInt() int { n, _ := strconv.Atoi(scanString()); return n }
func scanInts(n int) []int {
	ns := make([]int, n)
	for i := 0; i < n; i++ {
		ns[i] = scanInt()
	}
	return ns
}

func scanFloat() float64 { n, _ := strconv.ParseFloat(scanString(), 64); return n }
func scanFloats(n int) []float64 {
	ns := make([]float64, n)
	for i := 0; i < n; i++ {
		ns[i] = scanFloat()
	}
	return ns
}
