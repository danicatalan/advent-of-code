package sevensegmentsearch

import (
	"reflect"
	"testing"
)

var testCasesPart2 = []struct {
	input  string
	output int
}{
	{
		`be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`,
		61229,
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

var testCasesDecode = []struct {
	input  string
	output int
}{
	{
		`acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf`,
		5353,
	},
}

func TestDecode(t *testing.T) {
	for i, tc := range testCasesDecode {
		expected := tc.output
		actual := decode(tc.input)
		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("TestCase %d failed.\n\tExpected: %v\n\tGot: %v", i, expected, actual)
		}
	}
}

var testCasesContainsAll = []struct {
	input  string
	input2 string
	output bool
}{
	{
		"abc",
		"ab",
		true,
	},
	{
		"acb",
		"ab",
		true,
	},
	{
		"abcdef",
		"abef",
		true,
	},
	{
		"bcdefg",
		"abef",
		false,
	},
}

func TestContainsAll(t *testing.T) {
	for i, tc := range testCasesContainsAll {
		expected := tc.output
		actual := containsAll(tc.input, tc.input2)
		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("TestCase %d failed.\n\tExpected: %v\n\tGot: %v", i, expected, actual)
		}
	}
}
