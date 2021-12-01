package encodingerror

import (
	"reflect"
	"testing"
)

var testCasesPart2 = []struct {
	input  string
	target int
	output int
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
		127,
		62,
	},
}

func TestPart2(t *testing.T) {
	for i, tc := range testCasesPart2 {
		expected := tc.output
		actual := Part2(tc.input, tc.target)
		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("TestCase %d failed.\n\tExpected: %v\n\tGot: %v", i, expected, actual)
		}
	}
}
