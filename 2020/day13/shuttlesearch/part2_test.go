package shuttlesearch

import (
	"reflect"
	"testing"
)

var testCasesPart2 = []struct {
	input  string
	output int
}{
	{
		`7,13`,
		77,
	},
	{
		`17,x,13,19`,
		3417,
	},
	{
		`67,7,59,61`,
		754018,
	},
	{
		`67,x,7,59,61`,
		779210,
	},
	{
		`7,13,x,x,59,x,31,19`,
		1068781,
	},
	{
		`67,7,x,59,61`,
		1261476,
	},
	{
		`1789,37,47,1889`,
		1202161486,
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
