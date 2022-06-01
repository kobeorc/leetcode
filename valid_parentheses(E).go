package main

import "strings"

func isValid(s string) bool {
	var stack []string
	var m, r map[string]string
	m = make(map[string]string)
	r = make(map[string]string)
	m["("] = ")"
	m["["] = "]"
	m["{"] = "}"
	r[")"] = "("
	r["]"] = "["
	r["}"] = "{"

	a := strings.Split(s, "")

	for _, v := range a {
		if _, ok := m[v]; ok {
			stack = append(stack, v)
		} else {
			if len(stack) <= 0 {
				return false
			}
			if r[v] == stack[len(stack)-1] {
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}

	return false
}
