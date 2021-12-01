package tickettranslation

import (
	"log"
	"strconv"
	"strings"
)

//Part1 ...
func Part1(input string) int {
	notes := strings.Split(input, "\n\n")
	fields := strings.Split(notes[0], "\n")
	rules := make(map[string]func(int) bool)
	for _, f := range fields {
		name := strings.Split(f, ": ")[0]
		ranges := strings.Split(strings.Split(f, ": ")[1], " or ")
		range1, range2 := strings.Split(ranges[0], "-"), strings.Split(ranges[1], "-")

		var numbers1, numbers2 []int
		for _, n := range range1 {
			ni, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			numbers1 = append(numbers1, ni)
		}
		for _, n := range range2 {
			ni, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			numbers2 = append(numbers2, ni)
		}

		rules[name] = func(i int) bool {
			if numbers1[0] <= i && i <= numbers1[1] {
				return true
			}
			if numbers2[0] <= i && i <= numbers2[1] {
				return true
			}
			return false
		}
	}

	var invalidValues int
	tickets := strings.Split(notes[2], "\n")[1:]
	for _, t := range tickets {
		for _, v := range strings.Split(t, ",") {
			vi, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			var valid bool
			for _, f := range rules {
				if f(vi) {
					valid = true
					break
				}
			}
			if !valid {
				invalidValues += vi
			}
		}
	}
	return invalidValues
}
