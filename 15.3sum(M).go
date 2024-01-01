package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	buf := make(map[string][]int)
	storage := make(map[int]int)
	sort.Ints(nums)
	d := 0
	tmp := abs(nums[0])
	for i, v := range nums {
		storage[v] = i
		if abs(v) < tmp {
			tmp = abs(v)
			d = i
		}
	}

	for j := 0; j <= d; j++ {
		for k := d; k < len(nums); k++ {
			key := -(nums[j] + nums[k])
			if i, ok := storage[key]; ok && !(i == j || i == k || j == k) {
				buf = calc([]int{key, nums[j], nums[k]}, buf)
			}
		}
	}
	var r [][]int

	for _, v := range buf {
		r = append(r, v)
	}

	return r
}

func calc(a []int, buf map[string][]int) map[string][]int {
	sort.Ints(a)
	key := fmt.Sprintf("%d", a[0]) + "|" + fmt.Sprintf("%d", a[1]) + "|" + fmt.Sprintf("%d", a[2])
	buf[key] = a
	return buf
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
