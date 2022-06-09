package main

func twoSum2(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i == j {
				continue
			}
			if numbers[i]+numbers[j] == target {
				return []int{i+1, j+1}
			}
		}
	}

	return nil
}
