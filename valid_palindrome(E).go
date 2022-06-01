package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	r := reg.ReplaceAllString(s, "")
	r = strings.ToLower(r)
	ss := strings.Split(r, "")

	for i := 0; i < (len(ss) / 2); i++ {
		if ss[i] != ss[len(ss)-1-i] {
			return false
		}
	}

	return true
}
