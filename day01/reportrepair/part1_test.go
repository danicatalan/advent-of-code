package reportrepair

import (
	"reflect"
	"testing"
)

var testCasesPart1 = []struct {
	input  []int
	output int
}{
	{
		[]int{
			1721,
			979,
			366,
			299,
			675,
			1456,
		},
		514579,
	},
}

func TestPart1(t *testing.T) {
	for _, tc := range testCasesPart1 {
		expected := tc.output
		actual := Part1(tc.input)
		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("\n\tExpected: %v\n\tGot: %v", expected, actual)
		}
	}
}
