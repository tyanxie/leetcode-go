package main

/*
https://leetcode.cn/problems/determine-the-winner-of-a-bowling-game
*/

func isWinner(player1 []int, player2 []int) int {
	score1 := score(player1)
	score2 := score(player2)

	if score1 > score2 {
		return 1
	} else if score2 > score1 {
		return 2
	} else {
		return 0
	}
}

func score(player []int) int {
	sum := 0
	s10Index := -3
	for i, s := range player {
		sum += s
		if i-2 <= s10Index {
			sum += s
		}
		if s == 10 {
			s10Index = i
		}
	}
	return sum
}

func main() {

}
