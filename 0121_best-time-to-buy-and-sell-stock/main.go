package main

/*
https://leetcode.cn/problems/best-time-to-buy-and-sell-stock
*/

func maxProfit(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	res := 0
	minp := prices[0]
	for _, price := range prices {
		if minp >= price {
			minp = price
		} else {
			curr := price - minp
			if curr > res {
				res = curr
			}
		}
	}
	return res
}

func main() {

}
