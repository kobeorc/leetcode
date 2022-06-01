package main

func singleNumber(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		if _, ok := m[v]; ok {
			delete(m, v)
		} else {
			m[v] = v
		}
	}

	for _, v := range m {
		return v
	}

	return 0
}
