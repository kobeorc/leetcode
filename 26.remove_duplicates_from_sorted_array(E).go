package main

func removeDuplicates(nums []int) int {
	for i := len(nums); i >= 2; i-- {
		if nums[i-1] == nums[i-2] {
			nums = remove(nums, i-1)
		}
	}
	return len(nums)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
