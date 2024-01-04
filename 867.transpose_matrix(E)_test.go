package main

import (
	"fmt"
	"testing"
)

func Test_transpose(t *testing.T) {
	i1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	a1 := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	i2 := [][]int{{1, 2, 3}, {4, 5, 6}}
	a2 := [][]int{{1, 4}, {2, 5}, {3, 6}}
	fmt.Println("r:", transpose(i1), " expect: ", a1)
	fmt.Println("r:", transpose(i2), " expect: ", a2)
}
