package binaryboarding

import (
	"reflect"
	"testing"
)

var testCasesPart2 = []struct {
	input  []int
	output int
}{
	{
		[]int{
			12,
			11,
			20,
			19,
			16,
			18,
			10,
			13,
			14,
			15,
		},
		17,
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
