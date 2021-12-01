package customcustoms

import (
	"reflect"
	"testing"
)

var testCasesPart2 = []struct {
	input  string
	output int
}{
	{
		`abc`,
		3,
	},
	{
		`a
		b
		c`,
		0,
	},
	{
		`ab
		ac`,
		1,
	},
	{
		`a
		a
		a
		a`,
		1,
	},
	{
		`b`,
		1,
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
