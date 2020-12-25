package rambunctiousrecitation

import (
	"reflect"
	"testing"
)

var testCasesPart2 = []struct {
	input  string
	output int
}{
	{`0,3,6`, 175594},
	{`1,3,2`, 2578},
	{`2,1,3`, 3544142},
	{`1,2,3`, 261214},
	{`2,3,1`, 6895259},
	{`3,2,1`, 18},
	{`3,1,2`, 362},
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
