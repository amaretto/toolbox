package main

import (
	"reflect"
	"testing"
)

func TestSortInt(t *testing.T) {
	nums := []int{10, 5, 8, 3, 2, 9, 4, 6, 1, 7}

	result := SortInt(nums)
	answer := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(result, answer) {
		t.Fatal("failed test")
	}
}

func TestSortIntStable(t *testing.T) {
	nums := []int{10, 5, 8, 3, 2, 9, 4, 6, 1, 7}

	result := SortIntStable(nums)
	answer := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(result, answer) {
		t.Fatal("failed test")
	}
}
