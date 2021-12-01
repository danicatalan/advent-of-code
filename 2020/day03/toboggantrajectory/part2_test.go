package toboggantrajectory

import (
	"reflect"
	"testing"
)

var testCasesPart2 = []struct {
	input  string
	slopes [][]int
	output int
}{
	{
		`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`,
		[][]int{
			{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2},
		},
		336,
	},
}

func TestPart2(t *testing.T) {
	for i, tc := range testCasesPart2 {
		expected := tc.output
		actual := Part2(tc.input, tc.slopes)
		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("TestCase %d failed.\n\tExpected: %v\n\tGot: %v", i, expected, actual)
		}
	}
}

var testCasesPart2_1 = []struct {
	input  string
	sx     int
	sy     int
	output int
}{
	{
		`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`,
		1,
		1,
		2,
	},
	{
		`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`,
		3,
		1,
		7,
	},
	{
		`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`,
		5,
		1,
		3,
	},
	{
		`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`,
		7,
		1,
		4,
	},
	{
		`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`,
		1,
		2,
		2,
	},
}

func TestPart2_1(t *testing.T) {
	for i, tc := range testCasesPart2_1 {
		expected := tc.output
		actual := Part2_1(tc.input, tc.sx, tc.sy)
		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("TestCase %d failed.\n\tExpected: %v\n\tGot: %v", i, expected, actual)
		}
	}
}
