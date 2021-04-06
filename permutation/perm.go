package main

import "fmt"

func main() {
	//pattern := order(8)
	//fmt.Println(pattern)
	rest := []int{1, 1, 2, 3}
	head := []int{}
	result := perm(head, rest)
	fmt.Println(result)
}

func order(size int) [][]int {
	rest := make([]int, size)
	var head []int
	for i := 0; i < size; i++ {
		rest[i] = i
	}
	return perm(head, rest)
}

func perm(head, rest []int) [][]int {
	var res [][]int
	if len(rest) == 0 {
		res = append(res, head)
		return res
	}
	for i := 0; i < len(rest); i++ {
		restx := make([]int, len(rest))
		headx := make([]int, len(head)) //
		copy(restx, rest)
		copy(headx, head) //
		headx = append(headx, restx[i])
		restx = append(restx[:i], restx[i+1:]...)
		res = append(res, perm(headx, restx)...)
	}
	return res
}
