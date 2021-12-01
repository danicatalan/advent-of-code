package handheldhalting

import (
	"reflect"
	"testing"
)

var testCasesPart1 = []struct {
	input  string
	output int
}{
	{
		`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`,
		5,
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
