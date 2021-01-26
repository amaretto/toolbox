package main

import "sort"

func SortInt(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	return nums
}

func SortIntStable(nums []int) []int {
	sort.SliceStable(nums, func(i, j int) bool { return nums[i] < nums[j] })
	return nums
}
