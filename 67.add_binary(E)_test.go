package main

import "testing"

type TestCaseAddBinary struct {
	input  InputAddBinary
	output string
}
type InputAddBinary struct {
	v1 string
	v2 string
}

func TestAddBinary(t *testing.T) {
	var testCase []TestCaseAddBinary

	testCase = append(testCase, TestCaseAddBinary{
		input: InputAddBinary{
			v1: "11",
			v2: "1",
		},
		output: "100",
	})
	testCase = append(testCase, TestCaseAddBinary{
		input: InputAddBinary{
			v1: "1010",
			v2: "1011",
		},
		output: "10101",
	})

	for _, v := range testCase {
		r := addBinary(v.input.v1, v.input.v2)

		if r != v.output {
			t.Error(v.input.v1, " + ", v.input.v2, " != ", v.output, " result: ", r)
		}
	}
}
