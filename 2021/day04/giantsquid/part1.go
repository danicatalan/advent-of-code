package giantsquid

import (
	"strconv"
	"strings"
)

//Part1 resolves a bingo game and picks the winner.
func Part1(input string) int {
	draw, boards := importGame(input)
	winner, last := pickWinner(draw, boards)
	return sum(winner[:len(winner)/2]) * last
}

func sum(board [][]string) int {
	output := 0
	for _, row := range board {
		for _, number := range row {
			n, _ := strconv.Atoi(number)
			output += n
		}
	}
	return output
}

func pickWinner(draw []string, boards [][][]string) ([][]string, int) {
	for _, number := range draw {
		for i, board := range boards {
			for j, series := range board {
				res, win := tick(number, series)
				boards[i][j] = res
				if win {
					n, _ := strconv.Atoi(number)
					return boards[i], n
				}
			}
		}
	}
	return [][]string{}, 0
}

func tick(number string, series []string) ([]string, bool) {
	output := []string{}
	for _, item := range series {
		if item != number {
			output = append(output, item)
		}
	}
	return output, len(output) == 0
}

func importGame(input string) ([]string, [][][]string) {
	data := strings.Split(input, "\n")
	draw := strings.Split(data[0], ",")
	boards := strings.Split(strings.Join(data[2:], "\n"), "\n\n")
	normalizedBoards := [][][]string{}
	for _, b := range boards {
		normalizedBoards = append(normalizedBoards, normalize(b))
	}
	return draw, normalizedBoards
}

func normalize(input string) [][]string {
	rows, columns := [][]string{}, [][]string{}
	lines := strings.Split(strings.ReplaceAll(input, "  ", " "), "\n")
	for _, line := range lines {
		row := strings.Split(strings.TrimSpace(line), " ")
		rows = append(rows, row)
	}
	for i := range rows[0] {
		column := []string{}
		for _, r := range rows {
			column = append(column, r[i])
		}
		columns = append(columns, column)
	}
	output := [][]string{}
	output = append(output, rows...)
	output = append(output, columns...)
	return output
}
