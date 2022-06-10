package main

import (
	"strings"
)

func convertToTitle(c int) string {
	var r []string

	for {
		if c%26 == 0 {
			r = append(r, string(rune('A'-1+26)))
			c = c / 26
			c--
			if c == 0 {
				break
			}
		} else {
			r = append(r, string(rune('A'-1+(c%26))))
			c = c / 26
		}

		if c <= 0 {
			break
		}
	}

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return strings.Join(r, "")
}
