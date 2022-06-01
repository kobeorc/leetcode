package main

func checkValid(matrix [][]int) bool {
	l := len(matrix)
	for i := 0; i < l; i++ {
		b := make(map[int]bool)
		for j := 0; j < l; j++ {
			v := matrix[i][j]
			if _, ok := b[v]; ok {
				return false
			}
			b[v] = true
		}
	}
	for i := 0; i < l; i++ {
		b := make(map[int]bool)
		for j := 0; j < l; j++ {
			v := matrix[j][i]
			if _, ok := b[v]; ok {
				return false
			}
			b[v] = true
		}
	}
	return true
}
