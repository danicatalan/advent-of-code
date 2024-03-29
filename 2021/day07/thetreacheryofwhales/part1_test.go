package thetreacheryofwhales

import (
	"reflect"
	"testing"
)

var testCasesPart1 = []struct {
	input  string
	output int
}{
	{
		`16,1,2,0,4,2,7,1,2,14`,
		37,
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
