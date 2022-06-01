package main

func runningSum(nums []int) []int {
	var r []int

	for i, v := range nums {
		if i == 0 {
			r = append(r, v)
			continue
		}
		r = append(r, v+r[i-1])
	}

	return r
}
