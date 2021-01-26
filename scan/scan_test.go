package main

import "testing"

var stringTestCases = []struct {
	Input  string
	Output string
}{
	{"a", "a"},
	{"abc abc abc", "abc abc abc"},
}

var stringsTestCases = []struct {
	Input  string
	Output []string
}{
	{"a", []string{"a"}},
	{"abc abc abc", []string{"abc abc abc"}},
}

func TestSumSuccess(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	result := sum(nums)

	if result != 15 {
		t.Fatal("failed test")
	}
}
