package tickettranslation

import (
	"reflect"
	"testing"
)

// Note: field names modified to test the "departure" values.
// Also added an invalid ticket (last one) to cover exclusion of invalid tickets.
var testCasesPart2 = []struct {
	input  string
	output int
}{
	{
		`departure class: 0-1 or 4-19
departure row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9
6,10,20`,
		132,
	},
}

func TestPart2(t *testing.T) {
	for i, tc := range testCasesPart2 {
		expected := tc.output
		actual := Part2(tc.input)
		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("TestCase %d failed.\n\tExpected: %v\n\tGot: %v", i, expected, actual)
		}
	}
}
