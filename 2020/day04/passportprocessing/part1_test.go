package passportprocessing

import (
	"reflect"
	"testing"
)

var testCasesPart1 = []struct {
	input  string
	output bool
}{
	{
		`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm`,
		true,
	},
	{
		`iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929`,
		false,
	},
	{
		`hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm`,
		true,
	},
	{
		`hcl:#cfa07d eyr:2025 pid:166559648
	iyr:2011 ecl:brn hgt:59in`,
		false,
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
