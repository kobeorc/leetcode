package main

func plusOne(digits []int) []int {
	plus := 1

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+plus > 9 {
			plus = 1
			digits[i] = 0
			if i == 0 {
				return append([]int{1}, digits...)
			}
		} else {
			digits[i] = digits[i] + plus
			plus = 0
		}
	}

	return digits
}
