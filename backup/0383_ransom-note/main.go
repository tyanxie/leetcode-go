package main

import "fmt"

/*
https://leetcode.cn/problems/ransom-note
*/

func canConstruct(ransomNote string, magazine string) bool {
	var chars [26]int
	for _, b := range []byte(magazine) {
		b = b - 'a'
		chars[b] += 1
	}
	for _, b := range []byte(ransomNote) {
		b = b - 'a'
		if chars[b] == 0 {
			return false
		}
		chars[b] -= 1
	}
	return true
}

func main() {
	fmt.Println(canConstruct("a", "b"))
}
