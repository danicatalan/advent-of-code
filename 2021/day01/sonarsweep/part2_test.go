package sonarsweep

import (
	"reflect"
	"testing"
)

var testCasesPart2 = []struct {
	input  string
	output int
}{
	{
		`199
200
208
210
200
207
240
269
260
263`,
		5,
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
