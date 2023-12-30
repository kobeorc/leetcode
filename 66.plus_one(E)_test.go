package main

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input  []int
	output []int
}

func TestPlusOne(t *testing.T) {
	var testCase []TestCase
	testCase = append(testCase, TestCase{
		input:  []int{1, 2, 3},
		output: []int{1, 2, 4},
	})
	testCase = append(testCase, TestCase{
		input:  []int{4, 3, 2, 1},
		output: []int{4, 3, 2, 2},
	})
	testCase = append(testCase, TestCase{
		input:  []int{9},
		output: []int{1, 0},
	})
	testCase = append(testCase, TestCase{
		input:  []int{9, 9},
		output: []int{1, 0, 0},
	})

	for _, v := range testCase {
		r := plusOne(v.input)
		fmt.Println("res: ", fmt.Sprintf("%+v", r), "testcase: ", fmt.Sprintf("%+v", v.output))
	}
}
