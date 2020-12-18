package rainrisk

import (
	"reflect"
	"testing"
)

var testCasesPart2 = []struct {
	input  string
	output int
}{
	{
		`F10
N3
F7
R90
F11`,
		286,
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
