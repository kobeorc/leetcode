package main

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	m := make(map[string]int)
	m["abcabcbb"] = 3
	m["bbbbb"] = 1
	m["pwwkew"] = 3

	for k, v := range m {
		got := lengthOfLongestSubstring(k)
		if got != v {
			t.Errorf("got:%d, need:%d", got, v)
		}
	}
}
