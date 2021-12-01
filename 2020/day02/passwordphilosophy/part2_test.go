package passwordphilosophy

import (
	"reflect"
	"testing"
)

var testCasesPart2 = []struct {
	input  string
	output bool
}{
	{
		"1-3 a: abcde",
		true,
	},
	{
		"1-3 b: cdefg",
		false,
	},
	{
		"2-9 c: ccccccccc",
		false,
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
