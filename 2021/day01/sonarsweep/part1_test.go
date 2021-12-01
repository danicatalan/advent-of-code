package sonarsweep

import (
	"reflect"
	"testing"
)

var testCasesPart1 = []struct {
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
		7,
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
