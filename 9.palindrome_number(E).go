package main

import "strconv"

// func isPalindrome(x int) bool {
func isPalindromeNumber(x int) bool {
	s := strconv.Itoa(x)
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
