package main

func transpose(matrix [][]int) [][]int {
	w := len(matrix)
	h := len(matrix[0])

	nm := make([][]int, h)

	for i, _ := range nm {
		nm[i] = make([]int, w)
	}

	for y := 0; y < len(matrix[0]); y++ {
		for x := 0; x < len(matrix); x++ {
			nm[y][x] = matrix[x][y]
		}
	}
	return nm
}
