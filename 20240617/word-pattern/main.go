package main

import "strings"

/*
https://leetcode.cn/problems/word-pattern/
*/

func wordPattern(pattern string, s string) bool {
	words := strings.Fields(s)
	if len(pattern) != len(words) {
		return false
	}

	length := len(pattern)
	m := make(map[byte]string)
	exists := make(map[string]byte)

	for i := 0; i < length; i++ {
		word := words[i]
		b := pattern[i]
		if known, ok := m[b]; ok {
			if known != word {
				return false
			}
		} else {
			m[b] = word
		}

		if known, ok := exists[word]; ok {
			if known != b {
				return false
			}
		} else {
			exists[word] = b
		}
	}
	return true
}

func main() {

}
