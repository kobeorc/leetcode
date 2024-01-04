package main

import "strconv"

func addBinary(a string, b string) string {
	var plus int
	var result string
	la := len(a)
	lb := len(b)

	for {
		if la <= 0 && lb <= 0 {
			if plus == 1 {
				result = "1" + result
			}
			break
		}
		f := 0
		s := 0
		if la > 0 {
			f, _ = strconv.Atoi(string(a[la-1]))
			la--
		}
		if lb > 0 {
			s, _ = strconv.Atoi(string(b[lb-1]))
			lb--
		}
		switch f + s + plus {
		case 0:
			result = "0" + result
			plus = 0
		case 1:
			result = "1" + result
			plus = 0
		case 2:
			result = "0" + result
			plus = 1
		case 3:
			result = "1" + result
			plus = 1
		}
	}
	return result
}
