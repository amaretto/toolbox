package main

import "testing"

var testCases = []struct {
	Num    int
	Output bool
}{
	{1, false},
	{11, true},
	{20, false},
	{23, true},
}

func TestIsPrime(t *testing.T) {
	for _, testCase := range testCases {
		result := isPrime(testCase.Num)
		if result != testCase.Output {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}

func TestIsPrimeSmart(t *testing.T) {
	for _, testCase := range testCases {
		result := isPrimeSmart(testCase.Num)
		if result != testCase.Output {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}
