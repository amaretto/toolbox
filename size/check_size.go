package main

import (
	"fmt"
	"unsafe"
)

// CheckTypes have some type variables for checking these sizes
type CheckTypes struct {
	flag  bool
	num   int64
	ptr   *int64
	mini  int32
	str   string
	nums  []int64
	nums5 [5]int64
	strs  []string
	strs5 [5]string
	b     byte
	bs    []byte
	bs5   [5]byte
	r     rune
	rs    []rune
	rs5   [5]rune
}

func main() {
	a := CheckTypes{}
	// byte
	fmt.Println("struct\t\t: ", unsafe.Sizeof(a))
	fmt.Println("int64\t\t: ", unsafe.Sizeof(a.num))
	fmt.Println("*int64\t\t: ", unsafe.Sizeof(a.ptr))
	fmt.Println("int32\t\t: ", unsafe.Sizeof(a.mini))
	fmt.Println("string\t\t: ", unsafe.Sizeof(a.str))
	fmt.Println("[]int64\t\t: ", unsafe.Sizeof(a.num))
	fmt.Println("[5]int64\t: ", unsafe.Sizeof(a.nums5))
	fmt.Println("[]string\t: ", unsafe.Sizeof(a.strs))
	fmt.Println("[5]string\t: ", unsafe.Sizeof(a.strs5))
	fmt.Println("byte\t\t: ", unsafe.Sizeof(a.b))
	fmt.Println("[]byte\t\t: ", unsafe.Sizeof(a.bs))
	fmt.Println("[5]byte\t\t: ", unsafe.Sizeof(a.bs5))
	fmt.Println("rune\t\t: ", unsafe.Sizeof(a.r))
	fmt.Println("[]rune\t\t: ", unsafe.Sizeof(a.rs))
	fmt.Println("[5]rune\t\t: ", unsafe.Sizeof(a.rs5))
}
