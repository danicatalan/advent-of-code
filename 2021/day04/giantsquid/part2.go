package giantsquid

import (
	"strconv"
)

//Part2 resolves a bingo and picks the last winner.
func Part2(input string) int {
	draw, boards := importGame(input)
	looser, last := pickLastWinner(draw, boards)
	return sum(looser[:len(looser)/2]) * last
}

func pickLastWinner(draw []string, boards [][][]string) ([][]string, int) {
	for _, number := range draw {
		for i := 0; i < len(boards); i++ {
			board := boards[i]
			for j := 0; j < len(board); j++ {
				series := board[j]
				res, win := tick(number, series)
				boards[i][j] = res
				if win && len(boards) == 1 {
					n, _ := strconv.Atoi(number)
					return boards[i], n
				}
				if win {
					boards = append(boards[:i], boards[i+1:]...)
					i--
					break
				}
			}
		}
	}
	return [][]string{}, 0
}
