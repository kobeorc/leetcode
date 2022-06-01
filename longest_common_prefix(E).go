package main

import "strings"

func longestCommonPrefix(s []string) string {
	p := s[0]

	for _, v := range s {
		p = c(v, p)
	}

	return p
}

func c(s string, p string) string {
	if !strings.HasPrefix(s, p) {
		tm := strings.Split(p, "")
		tm = tm[:len(tm)-1]

		return c(s, strings.Join(tm, ""))
	}
	return p
}
