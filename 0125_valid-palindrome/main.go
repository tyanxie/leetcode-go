package main

import "fmt"

/*
https://leetcode.cn/problems/valid-palindrome
*/

const diff = 'a' - 'A'

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	bs := []byte(s)
	l := 0
	r := len(bs) - 1
	for l <= r {
		var lb byte
		for l <= r {
			lb = toLower(bs[l])
			if lb > 0 {
				break
			}
			l++
		}

		var rb byte
		for l <= r {
			rb = toLower(bs[r])
			if rb > 0 {
				break
			}
			r--
		}

		if lb != rb {
			return false
		}
		l++
		r--
	}
	return true
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		b += diff
	}
	if (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
		return b
	}
	return 0
}

func main() {
	// fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	// fmt.Println(isPalindrome("a."))
	fmt.Println(isPalindrome("0P"))
}
