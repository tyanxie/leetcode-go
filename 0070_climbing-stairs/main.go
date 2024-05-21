package main

/*
https://leetcode.cn/problems/climbing-stairs
*/

// climbStairs F(n) = F(n-1) + F(n-2)
func climbStairs(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	}

	first, second, third := 1, 2, 3
	for i := 3; i <= n; i++ {
		third = first + second
		first = second
		second = third
	}
	return third
}

func main() {

}
