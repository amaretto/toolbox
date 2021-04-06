package main

import "fmt"

// calcIntersection calculate coordinate of intersect two line.
// One of these line contains (x1,y1) and (x2,y2)
// The other side contains (x3,y3) and (x4,y4)
func calcIntersection(x1, y1, x2, y2, x3, y3, x4, y4 float64) (x, y float64) {
	a1 := (y2 - y1) / (x2 - x1)
	a3 := (y4 - y3) / (x4 - x3)
	x = (a1*x1 - y1 - a3*x3 + y3) / (a1 - a3)
	y = (y2-y1)/(x2-x1)*(x-x1) + y1
	return x, y
}

func main() {
	x, y := calcIntersection(0, 0, 1, 1, 0, 1, 1, 0)
	// The result is (0.5,0.5)
	fmt.Println(x, y)
}
