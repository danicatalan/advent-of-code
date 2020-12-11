package encodingerror

import (
	"reflect"
	"testing"
)

var testCasesPart1 = []struct {
	input    string
	preamble int
	output   int
}{
	{
		`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`,
		5,
		127,
	},
}

func TestPart1(t *testing.T) {
	for i, tc := range testCasesPart1 {
		expected := tc.output
		actual := Part1(tc.input, tc.preamble)
		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("TestCase %d failed.\n\tExpected: %v\n\tGot: %v", i, expected, actual)
		}
	}
}
