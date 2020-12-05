package binaryboarding

import (
	"reflect"
	"testing"
)

var testCasesPart1 = []struct {
	input  string
	output int
}{
	{
		`FBFBBFFRLR`,
		357,
	},
	{
		`BFFFBBFRRR`,
		567,
	},
	{
		`FFFBBBFRRR`,
		119,
	},
	{
		`BBFFBBFRLL`,
		820,
	},
}

func TestPart1(t *testing.T) {
	for i, tc := range testCasesPart1 {
		expected := tc.output
		actual := Part1(tc.input)
		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("TestCase %d failed.\n\tExpected: %v\n\tGot: %v", i, expected, actual)
		}
	}
}
