package main

import (
	"testing"
)

type testCases struct {
	assert int
	input  []int
}

func getTestCases() []testCases {
	return []testCases{
		{assert: 5, input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
		{assert: 2, input: []int{1, 1, 2}},
	}
}

func TestRemoveDuplicates(t *testing.T) {
	for _, testCase := range getTestCases() {
		res := removeDuplicates(testCase.input)
		if testCase.assert != res {
			t.Error("assert: ", testCase.assert, " != ", "result: ", res)
		}
	}
}
