package customcustoms

import (
	"reflect"
	"testing"
)

var testCasesPart1 = []struct {
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
		3,
	},
	{
		`ab
		ac`,
		3,
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

func TestPart1(t *testing.T) {
	for i, tc := range testCasesPart1 {
		expected := tc.output
		actual := Part1(tc.input)
		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("TestCase %d failed.\n\tExpected: %v\n\tGot: %v", i, expected, actual)
		}
	}
}
