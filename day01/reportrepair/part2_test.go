package reportrepair

import (
	"reflect"
	"testing"
)

var testCasesPart2 = []struct {
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
		241861950,
	},
}

func TestPart2(t *testing.T) {
	for _, tc := range testCasesPart2 {
		expected := tc.output
		actual := Part2(tc.input)
		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("\n\tExpected: %v\n\tGot: %v", expected, actual)
		}
	}
}
