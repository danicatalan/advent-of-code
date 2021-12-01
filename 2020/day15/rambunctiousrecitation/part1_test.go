package rambunctiousrecitation

import (
	"reflect"
	"testing"
)

var testCasesPart1 = []struct {
	input  string
	output int
}{
	{`0,3,6`, 436},
	{`1,3,2`, 1},
	{`2,1,3`, 10},
	{`1,2,3`, 27},
	{`2,3,1`, 78},
	{`3,2,1`, 438},
	{`3,1,2`, 1836},
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
