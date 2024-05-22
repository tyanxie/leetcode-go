package main

import "fmt"

/*
https://leetcode.cn/problems/isomorphic-strings
*/

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	length := len(s)
	sm := make(map[byte]byte)
	tm := make(map[byte]byte)
	for i := 0; i < length; i++ {
		sb := s[i]
		tb := t[i]

		if stb, ok := sm[sb]; ok {
			if tb != stb {
				return false
			}
		} else {
			sm[sb] = tb
		}

		if ssb, ok := tm[tb]; ok {
			if sb != ssb {
				return false
			}
		} else {
			tm[tb] = sb
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("badc", "bada"))
}
