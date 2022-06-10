package main

func lengthOfLongestSubstring(s string) int {
	var lc, c int
	m := make(map[string]bool)

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if _, ok := m[string(s[j])]; !ok {
				m[string(s[j])] = true
				lc++
			} else {
				break
			}
		}
		if c < lc {
			c = lc
		}
		lc = 0
		m = make(map[string]bool)
	}

	return c
}
